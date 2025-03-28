FROM golang:1.21-alpine
RUN apk add --no-cache git

ENV GO111MODULE="on"

WORKDIR /go/src/github.com/mparvin/netinfo-exporter
COPY . .

RUN go get -d -v ./...
RUN go build -o /go/bin/netinfo-exporter

FROM alpine:3.16
COPY --from=0 /go/bin/netinfo-exporter /usr/local/bin/netinfo-exporter
ENTRYPOINT ["/usr/local/bin/netinfo-exporter"]