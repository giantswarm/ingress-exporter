package service

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microendpoint/service/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/client/k8srestconfig"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/giantswarm/ingress-exporter/flag"
	"github.com/giantswarm/ingress-exporter/service/collector"
)

type Config struct {
	Logger micrologger.Logger

	Description string
	Flag        *flag.Flag
	GitCommit   string
	ProjectName string
	Source      string
	Viper       *viper.Viper
}

type Service struct {
	Version *version.Service

	bootOnce          sync.Once
	exporterCollector *collector.Set
}

func New(config Config) (*Service, error) {
	// Dependencies.
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	// Settings.
	if config.Flag == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Flag must not be empty", config)
	}
	if config.Viper == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Viper must not be empty", config)
	}

	var err error

	var kubernetesClient kubernetes.Interface
	{
		var config *rest.Config
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, microerror.Maskf(err, "failed to create in-cluster config for kubernetes client")
		}

		kubernetesClient, err = kubernetes.NewForConfig(config)
		if err != nil {
			return nil, microerror.Maskf(err, "failed to create kubernetes client")
		}
	}

	var g8sClient *versioned.Clientset
	{
		var restConfig *rest.Config
		{
			c := k8srestconfig.Config{
				Logger: config.Logger,

				Address:    config.Viper.GetString(config.Flag.Service.Kubernetes.Address),
				InCluster:  config.Viper.GetBool(config.Flag.Service.Kubernetes.InCluster),
				KubeConfig: config.Viper.GetString(config.Flag.Service.Kubernetes.KubeConfig),
				TLS: k8srestconfig.ConfigTLS{
					CAFile:  config.Viper.GetString(config.Flag.Service.Kubernetes.TLS.CAFile),
					CrtFile: config.Viper.GetString(config.Flag.Service.Kubernetes.TLS.CrtFile),
					KeyFile: config.Viper.GetString(config.Flag.Service.Kubernetes.TLS.KeyFile),
				},
			}
			restConfig, err = k8srestconfig.New(c)
			if err != nil {
				return nil, microerror.Mask(err)
			}
		}

		g8sClient, err = versioned.NewForConfig(restConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var exporterCollector *collector.Set
	{
		c := collector.SetConfig{
			G8sClient:        g8sClient,
			KubernetesClient: kubernetesClient,
			Logger:           config.Logger,
		}

		exporterCollector, err = collector.NewSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var versionService *version.Service
	{
		c := version.Config{
			Description: config.Description,
			GitCommit:   config.GitCommit,
			Name:        config.ProjectName,
			Source:      config.Source,
		}

		versionService, err = version.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	s := &Service{
		Version: versionService,

		bootOnce:          sync.Once{},
		exporterCollector: exporterCollector,
	}

	return s, nil
}

func (s *Service) Boot(ctx context.Context) {
	s.bootOnce.Do(func() {
		go s.exporterCollector.Boot(ctx)
	})
}

func mustParseJSONList(s string) []string {
	var l []string
	err := json.Unmarshal([]byte(s), &l)
	if err != nil {
		panic(err)
	}

	return l
}
