# syntax=docker/dockerfile:1

FROM golang:alpine AS base
WORKDIR /app
RUN apk add --no-cache git ca-certificates

FROM base AS dev
RUN go install github.com/air-verse/air@latest && \
    go install github.com/a-h/templ/cmd/templ@latest

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]

FROM base AS builder
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server ./cmd/web

FROM alpine:latest AS prod
RUN apk --no-cache add ca-certificates tzdata && \
    adduser -D -g '' appuser

WORKDIR /app
COPY --from=builder /app/server /app/server

USER appuser
EXPOSE 8080

CMD ["/app/server"]
