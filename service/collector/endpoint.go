package collector

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	g8sClient     *versioned.Clientset
	httpClient    *http.Client
	httpTransport *http.Transport
	k8sClient     kubernetes.Interface
	localIP       string
	logger        micrologger.Logger

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

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:    maxIdleConnection,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   maxTimeoutSec * time.Second,
	}

	i := &Endpoint{
		g8sClient:     config.G8sClient,
		httpClient:    client,
		httpTransport: tr,
		k8sClient:     config.K8sClient,
		localIP:       getLocalIP(),
		logger:        config.Logger,
	}

	return i, nil
}

func (e *Endpoint) Collect(ch chan<- prometheus.Metric) error {
	// TODO
	// Here we will put implementation of collecting data for each endpoint IP
	listOpts := metav1.ListOptions{}
	getOpts := metav1.GetOptions{}

	kvmConfigs, err := e.g8sClient.ProviderV1alpha1().KVMConfigs(crdNamespace).List(listOpts)
	if err != nil {
		return microerror.Mask(err)
	}

	for _, kvmConfig := range kvmConfigs.Items {
		if kvmConfig.DeletionTimestamp != nil {
			// ignore clusters that are marked for deletion
			continue
		}

		endpoint, err := e.k8sClient.CoreV1().Endpoints(kvmConfig.Name).Get(workerEndpoint, getOpts)
		if err != nil {
			return microerror.Mask(err)
		}

		ipList := getEndpointIps(endpoint)
		for _, ip := range ipList {
			err := e.ingressEndpointUpClassicHttp(ip, ingresSchemeHttp, ingressPortHttp)

			// io.EOF error can be indicator of enabled proxy-protocol
			if err != nil {
				err := e.ingressEndpointUpProxyProtocol(ip, ingresSchemeHttp, ingressPortHttp)
				if err != nil {
					fmt.Printf("proxy protocol http error: %s\n", err)
					return microerror.Mask(err)
				}
				// ingress check failed
			}
		}
	}

	return nil
}

func (e *Endpoint) Describe(ch chan<- *prometheus.Desc) error {
	ch <- endpointLabelsDesc
	return nil
}

func getEndpointIps(endpoint *corev1.Endpoints) []string {
	var ipList []string
	for _, subset := range endpoint.Subsets {
		for _, address := range subset.Addresses {
			ipList = append(ipList, address.IP)
		}
	}

	return ipList
}

func (e *Endpoint) buildHttpRequest(ipAddress string, scheme string, port int) (*http.Request, error) {
	u := url.URL{
		Host:   fmt.Sprintf("%s:%d", ipAddress, port),
		Path:   "healthz",
		Scheme: scheme,
	}

	// be sure to close idle connection after health check is finished
	defer e.httpTransport.CloseIdleConnections()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, microerror.Maskf(err, "unable to construct health check request")
	}

	// close connection after health check request (the TCP connection gets
	// closed by deferred s.tr.CloseIdleConnections()).
	req.Header.Add("Connection", "close")

	return req, nil
}

// ingressEndpointUpClassicHttp send http packet to the endpoint to ensure target ingress endpoint ip is up
func (e *Endpoint) ingressEndpointUpClassicHttp(ipAddress string, scheme string, port int) error {
	req, err := e.buildHttpRequest(ipAddress, scheme, port)
	if err != nil {
		// failed to build http req
		return microerror.Mask(err)
	}

	// send request to http endpoint
	response, err := e.httpClient.Do(req)
	if err != nil {
		// ingress endpoint failed to respond properly
		return microerror.Mask(err)
	}
	fmt.Printf("classic http response status %s\n", response.Status)

	return nil
}

// ingressEndpointUpProxyProtocol send encapsulated http packet into proxy-protocol to the endpoint to ensure target ingress endpoint ip is up
func (e *Endpoint) ingressEndpointUpProxyProtocol(ipAddress string, scheme string, port int) error {
	var buffer bytes.Buffer
	// build http request
	req, err := e.buildHttpRequest(ipAddress, scheme, port)
	if err != nil {
		return microerror.Mask(err)
	}
	// write proxy-protocol header to buffer
	_, err = fmt.Fprintf(&buffer, "PROXY TCP4 %s %s 80 80\r\n", e.localIP, e.localIP)
	if err != nil {
		return microerror.Mask(err)
	}
	// write http request to buffer
	err = req.Write(&buffer)
	if err != nil {
		return microerror.Mask(err)
	}

	// open tcp connection to ingress endpoint
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ipAddress, port))
	if err != nil {
		// handle error
		return microerror.Mask(err)
	}
	defer conn.Close()

	// write buffer to connection
	_, err = buffer.WriteTo(conn)
	if err != nil {
		// handle error
		return microerror.Mask(err)
	}

	// read http response from connection
	response, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		// handle error
		return microerror.Mask(err)
	}
	fmt.Printf("proxy protocol response status %s\n", response.Status)

	return nil
}

// getLocalIP returns the non loopback local IP of the host.
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
