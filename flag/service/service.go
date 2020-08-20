package service

import (
	"github.com/giantswarm/operatorkit/v2/pkg/flag/service/kubernetes"

	"github.com/giantswarm/ingress-exporter/flag/service/ingress"
)

type Service struct {
	Ingress    ingress.Ingress
	Kubernetes kubernetes.Kubernetes
}
