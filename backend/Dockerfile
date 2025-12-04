# syntax=docker/dockerfile:1.7

FROM golang:1.24-alpine AS builder
WORKDIR /workspace
RUN apk add --no-cache build-base git
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go test ./...
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go build -trimpath -ldflags="-s -w" -o /workspace/bin/market-api ./cmd/api

FROM gcr.io/distroless/base-debian12 AS runtime
WORKDIR /app
COPY --from=builder /workspace/bin/market-api ./market-api
EXPOSE 8080
ENV GIN_MODE=release
ENTRYPOINT ["./market-api"]
