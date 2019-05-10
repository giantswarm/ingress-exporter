package collector

import (
	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/kubernetes"
)

var (
	endpointLabelsDesc *prometheus.Desc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, "avaiable"),
		"Availability of ingress controller on ingress endpoint ip.",
		[]string{
			labelClusterID,
			labelIP,
			labelProtocol,
			labelProxyProtocol,
		},
		nil,
	)
)

type EndpointConfig struct {
	G8sClient *versioned.Clientset
	K8sClient kubernetes.Interface
	Logger    micrologger.Logger
}

type Endpoint struct {
	g8sClient *versioned.Clientset
	k8sClient kubernetes.Interface
	logger    micrologger.Logger

	customLabels []string
}

func NewEndpoint(config EndpointConfig) (*Endpoint, error) {
	if config.G8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.G8sClient must not be empty", config)
	}
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	i := &Endpoint{
		g8sClient: config.G8sClient,
		k8sClient: config.K8sClient,
		logger:    config.Logger,
	}

	return i, nil
}

func (i *Endpoint) Collect(ch chan<- prometheus.Metric) error {
	// TODO
	// Here we will put implementation of collecting data for each endpoint IP

	return nil
}

func (i *Endpoint) Describe(ch chan<- *prometheus.Desc) error {
	ch <- endpointLabelsDesc
	return nil
}
