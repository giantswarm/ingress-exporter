module github.com/giantswarm/ingress-exporter

go 1.14

require (
	github.com/giantswarm/apiextensions/v2 v2.6.2
	github.com/giantswarm/exporterkit v1.0.0
	github.com/giantswarm/k8sclient/v4 v4.1.0
	github.com/giantswarm/microendpoint v1.0.0
	github.com/giantswarm/microerror v0.4.0
	github.com/giantswarm/microkit v1.0.0
	github.com/giantswarm/micrologger v0.6.0
	github.com/giantswarm/operatorkit/v2 v2.0.2
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/prometheus/client_golang v1.15.1
	github.com/spf13/viper v1.15.0
	k8s.io/api v0.18.19
	k8s.io/apimachinery v0.18.19
	k8s.io/client-go v0.18.19
)

replace (
	github.com/caddyserver/caddy v1.0.3 => github.com/caddyserver/caddy v1.0.5
	github.com/coreos/etcd v3.3.10+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/gin-gonic/gin v1.4.0 => github.com/gin-gonic/gin v1.8.1
	github.com/gogo/protobuf v1.3.1 => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket v1.4.0 => github.com/gorilla/websocket v1.4.2
	github.com/kataras/iris/v12 v12.0.1 => github.com/kataras/iris/v12 v12.1.8
	github.com/labstack/echo/v4 v4.1.11 => github.com/labstack/echo/v4 v4.7.2
	github.com/microcosm-cc/bluemonday v1.0.2 => github.com/microcosm-cc/bluemonday v1.0.18
	github.com/valyala/fasthttp v1.6.0 => github.com/valyala/fasthttp v1.38.0
	go.mongodb.org/mongo-driver v1.1.2 => go.mongodb.org/mongo-driver v1.9.1
)
