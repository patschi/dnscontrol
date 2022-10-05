FROM golang:1.19.2-alpine3.16@sha256:2baa528036c1916b23de8b304083c68fb298c5661203055f2b1063390e3cdddb AS build

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
