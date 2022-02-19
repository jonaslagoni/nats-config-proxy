FROM golang:1.15-alpine3.13 AS builder

WORKDIR $GOPATH/src/github.com/nats-io/nats-config-proxy/

MAINTAINER Waldemar Quevedo <wally@synadia.com>

RUN apk add --update git

COPY . .

RUN CGO_ENABLED=0 GO111MODULE=on go build -v -a -tags netgo -installsuffix netgo -ldflags "-s -w" -o /nats-config-proxy

FROM alpine:3.13

RUN apk add --update ca-certificates && mkdir -p /nats/bin && mkdir /nats/conf

COPY --from=builder /nats-config-proxy /nats/bin/nats-config-proxy

RUN ln -ns /nats/bin/nats-config-proxy /bin/nats-config-proxy

EXPOSE 4567

ENTRYPOINT ["/bin/nats-config-proxy"]
