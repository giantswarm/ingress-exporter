module github.com/giantswarm/ingress-exporter

go 1.14

require (
	github.com/giantswarm/apiextensions v0.3.0
	github.com/giantswarm/exporterkit v0.2.0
	github.com/giantswarm/k8sclient v0.2.0
	github.com/giantswarm/microendpoint v0.2.0
	github.com/giantswarm/microerror v0.2.0
	github.com/giantswarm/microkit v0.2.1
	github.com/giantswarm/micrologger v0.3.1
	github.com/giantswarm/operatorkit v0.2.0
	github.com/prometheus/client_golang v1.3.0
	github.com/spf13/viper v1.6.3
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	k8s.io/api v0.16.6
	k8s.io/apimachinery v0.16.6
	k8s.io/client-go v0.16.6
)

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.5-beta.1
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191114101535-6c5935290e33
)
