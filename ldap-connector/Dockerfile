# syntax=docker/dockerfile:1
FROM registry-1.docker.io/library/golang:1.19.4-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app ./cmd

# ---
FROM registry-1.docker.io/library/alpine as certs
RUN apk update && apk add ca-certificates

# ---
FROM registry-1.docker.io/library/busybox:1.35.0-musl

ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static-muslc-amd64 /tini
RUN chmod +x /tini

RUN mkdir /app
RUN addgroup -S app && adduser -S -G app app
WORKDIR /app

COPY --from=builder /build/app .
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

RUN chown -R app:app .
USER app

ENTRYPOINT [ "/tini", "--", "./app" ]
