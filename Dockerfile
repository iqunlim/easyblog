FROM golang:latest AS base

FROM base AS dev

ENTRYPOINT ["tail", "-f", "/dev/null"]

FROM base AS prod

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build . 

EXPOSE 8080

CMD ["/app/easyblog"]