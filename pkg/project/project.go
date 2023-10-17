package project

var (
	description = "The ingress-exporter exports Prometheus metrics for tenant cluster ingress."
	gitSHA      = "n/a"
	name        = "ingress-exporter"
	source      = "https://github.com/giantswarm/ingress-exporter"
	version     = "1.2.1-dev"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
