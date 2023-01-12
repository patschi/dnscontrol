FROM golang:1.19.5-alpine3.16@sha256:65060885c8882f0119d77edcb42b414a0b96f6a8d466e9bbd782311ae51bf9f1 AS build

WORKDIR /go/src/github.com/StackExchange/dnscontrol

ARG BUILD_VERSION

ENV GO111MODULE on

COPY . .

# build dnscontrol
RUN apk update \
    && apk add --no-cache ca-certificates curl gcc build-base git \
    && update-ca-certificates \
    && go build -v -trimpath -buildmode=pie -ldflags="-s -w -X main.SHA=${BUILD_VERSION}"

# Validation check
RUN cp dnscontrol /go/bin/dnscontrol
RUN dnscontrol version

# -----

FROM alpine:3.16.2@sha256:bc41182d7ef5ffc53a40b044e725193bc10142a1243f395ee852a8d9730fc2ad

COPY --from=build /etc/ssl/certs /etc/ssl/certs
COPY --from=build /go/bin/dnscontrol /usr/local/bin

WORKDIR /dns

ENTRYPOINT ["/usr/local/bin/dnscontrol"]
