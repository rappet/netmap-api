# Container that builds the app
FROM golang:alpine AS builder

MAINTAINER Raphael Peters <raphael.r.peters@gmail.com>

WORKDIR $GOPATH/src/git.rappet.de/rappet/netmap-api

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /netmap-api -v

# Container to be deployed
FROM alpine:latest
COPY --from=builder /netmap-api ./

ENV DB_CONN_STR="postgres://netmap:netmap@db/netmap?sslmode=disable"
EXPOSE 8080/tcp

ENTRYPOINT ["./netmap-api"]