FROM alpine:3.13.5

RUN apk add --update ca-certificates \
    && rm -rf /var/cache/apk/*

ADD ./ingress-exporter /ingress-exporter

ENTRYPOINT ["/ingress-exporter"]
