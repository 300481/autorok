# Building stage
FROM golang:1.13.1-buster AS builder

WORKDIR /go/src/github.com/300481/autorok
COPY . .
WORKDIR /go/src/github.com/300481/autorok/cmd/autorok
RUN go get -d -v && \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /autorok

# Run stage
FROM alpine:3.10.2

COPY --from=builder /autorok /autorok
COPY dnsmasq.sh /

RUN apk --no-cache add \
    dnsmasq \
    sipcalc && \
    mkdir /tftp && cd /tftp && \
    wget http://boot.ipxe.org/undionly.kpxe && \
    cp undionly.kpxe undionly.kpxe.0 && \
    wget https://github.com/just-containers/s6-overlay/releases/download/v1.21.8.0/s6-overlay-amd64.tar.gz -O /tmp/s6-overlay.tar.gz && \
    tar xvzf /tmp/s6-overlay.tar.gz -C /

COPY services.d/ /etc/services.d/

WORKDIR /
EXPOSE 67 69 8080

ENTRYPOINT ["/init"]
