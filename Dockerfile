FROM golang:1.15-alpine3.12

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go get github.com/cespare/reflex
ENV CGO_ENABLED=0

WORKDIR /go/src/user-auth
COPY ./ ./

RUN go mod download