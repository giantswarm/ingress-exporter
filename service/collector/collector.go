package collector

const (
	namespace = "ingress_exporter"
	subsystem = "endpoint"
)

const (
	labelClusterID     = "cluster_id"
	labelIP            = "ip"
	labelProtocol      = "protocol"
	labelProxyProtocol = "proxy_protocol"
)

const (
	crdNamespace   = "default"
	workerEndpoint = "worker"
	masterEndpoint = "master"
)

const (
	maxIdleConnection = 50
	maxTimeoutSec     = 5

	ingresSchemeHttp   = "http"
	ingressSchemeHttps = "https"

	ingressPortHttp  = 30010
	ingressPortHttps = 30011
)
