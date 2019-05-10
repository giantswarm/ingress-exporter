package service

import (
	"github.com/giantswarm/operatorkit/flag/service/kubernetes"

	"github.com/giantswarm/ingress-exporter/flag/service/ingress"
)

type Service struct {
	Ingress ingress.Ingress
	Kubernetes     kubernetes.Kubernetes
}
