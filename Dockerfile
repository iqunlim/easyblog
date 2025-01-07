FROM node:18-alpine AS dependencies

RUN npm install -g typescript

FROM dependencies AS typescript

WORKDIR /app

COPY ./static/ts/ ./
COPY tsconfig.json ./tsconfig.json

WORKDIR /app

RUN ["tsc"]

FROM golang:1.23-alpine AS builder 

RUN apk add --no-cache \
    build-base \
    gcc \
    libc-dev \
    musl-dev \
    pkgconf \
    zlib-dev \
    openssl-dev

ENV CGO_ENABLED=1
ENV GOOS linux

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download 

COPY main.go main.go
COPY crypt crypt
COPY config config
COPY model model
COPY service service
COPY repository repository
COPY controller controller

RUN go build -o easyblog

FROM alpine:latest AS deploy

WORKDIR /app
ENV GOPATH /app

COPY --from=builder /app/easyblog /app/easyblog
COPY static /app/static/
COPY --from=typescript /app/static/js/ /app/static/js/

EXPOSE 8080

ENTRYPOINT ["./easyblog"]