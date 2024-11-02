FROM golang:alpine3.19

ENV CGO_ENABLED 1
WORKDIR /go/src/job/

RUN apk update && apk add --no-cache gcc musl-dev curl bash
RUN go install github.com/air-verse/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest