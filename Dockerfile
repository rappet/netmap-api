# Container that builds the app
FROM instrumentisto/dep:alpine AS builder

MAINTAINER Raphael Peters <raphael.r.peters@gmail.com>

WORKDIR $GOPATH/src/git.rappet.de/rappet/netmap-api

COPY Gopkg.lock Gopkg.toml vendor ./
RUN dep ensure --vendor-only

COPY . ./
RUN go build -o /netmap-api -v

# Container to be deployed
FROM alpine:latest
COPY --from=builder /netmap-api ./

EXPOSE 8080/tcp

ENTRYPOINT ["./netmap-api"]