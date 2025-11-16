# stage 1
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./docs /app/docs
COPY . .

RUN apk add --no-cache make bash \
    && make build


# stage 2
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server ./server
COPY --from=builder /app/migrator ./migrator
COPY --from=builder /app/config ./config
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/docs ./docs

RUN apk add --no-cache bash

CMD ["./server", "--config_path", "/app/config/docker.yaml"]
