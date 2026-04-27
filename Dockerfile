# syntax=docker/dockerfile:1.7

# ---------- builder ----------
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git build-base

WORKDIR /src

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

# Install templ (pinned to the version in go.mod)
RUN go install github.com/a-h/templ/cmd/templ@v0.3.1001

# Copy source and generate templ files
COPY . .
RUN templ generate

# CGO=1 is required for the sqlite driver
ENV CGO_ENABLED=1 GOOS=linux
RUN go build -trimpath -ldflags="-s -w" -o /out/lagarigo ./cmd/server

# ---------- runtime ----------
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata sqlite-libs \
    && adduser -D -u 10001 app

WORKDIR /app

COPY --from=builder /out/lagarigo /app/lagarigo
COPY --from=builder /src/public /app/public

USER app

ENV APP_PORT=3000
EXPOSE 3000

CMD ["/app/lagarigo"]
