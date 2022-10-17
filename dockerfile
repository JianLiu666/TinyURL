FROM golang:1.18.6-alpine3.16 AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on 
RUN mkdir -p /app
WORKDIR /app
COPY go.mod . 
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags "-s -w" -o tinyurl

FROM centos:centos7
RUN mkdir -p /app
COPY --from=builder /app/tinyurl /app/tinyurl
WORKDIR /app
COPY /conf.d/env.yaml .
CMD ["./tinyurl", "-f", "./env.yaml", "server"]