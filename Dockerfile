FROM node:18-alpine AS typescript

WORKDIR /app

COPY . .

RUN npm install -g typescript

WORKDIR /app

RUN ["tsc"]

FROM golang:latest AS base

FROM base AS dev

ENTRYPOINT ["tail", "-f", "/dev/null"]

FROM base AS prod

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN mkdir -p "/app/static/js"

COPY --from=typescript /app/static/js /app/static/js

RUN go build . 

EXPOSE 8080

CMD ["/app/easyblog"]