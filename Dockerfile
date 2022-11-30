FROM golang:1.19.3-alpine3.16 as builder
LABEL maintainer="dru90i"


COPY . /go/src/rancher-exporter/
RUN apk --update add ca-certificates \
 && apk --update add --virtual build-deps go git \ 
 && cd /go/src/rancher-exporter \
 && GOPATH=/go go get \
 && GOPATH=/go go build -o /bin/rancher_exporter \
 && apk del --purge build-deps \
 && rm -rf /go/bin /go/pkg /var/cache/apk/*

FROM alpine:latest

EXPOSE 9191
RUN addgroup exporter \
 && adduser -S -G exporter exporter \
 && apk --no-cache add ca-certificates

COPY --from=builder /bin/rancher_exporter /bin/rancher_exporter

USER exporter

ENTRYPOINT [ "/bin/rancher_exporter" ]
