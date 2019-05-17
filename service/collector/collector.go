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
	maxIdleConnection = 5
	maxTimeoutSec     = 5

	ingressCheckSuccessful = 1
	ingressCheckFailure   = 0

	ingressSchemeHttp   = "http"
	ingressSchemeHttps = "https"

	ingressPortHttp  = 30010
	ingressPortHttps = 30011

	proxyProtocolTrue    = "true"
	proxyProtocolFalse   = "false"
	proxyProtocolUnknown = "unknown"
)
