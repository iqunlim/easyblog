FROM golang:latest AS base

FROM base AS dev

ENTRYPOINT ["tail", "-f", "/dev/null"]

FROM base as prod

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download 

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /ebdatapull

EXPOSE 8080

CMD ["/ebdatapull"]