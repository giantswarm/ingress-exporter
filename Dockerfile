FROM quay.io/giantswarm/alpine:3.16.3

RUN apk add --update ca-certificates \
    && rm -rf /var/cache/apk/*

ADD ./ingress-exporter /ingress-exporter

ENTRYPOINT ["/ingress-exporter"]
