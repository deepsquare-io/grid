# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM registry-1.docker.io/library/golang:1.22.1-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS TARGETARCH
ARG BUILDVERSION=dev
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg \
  CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -ldflags "-s -w -X main.version=${BUILDVERSION}" -o /out/supervisor ./

# ---
FROM --platform=$BUILDPLATFORM registry-1.docker.io/library/alpine as certs
RUN apk update && apk add ca-certificates

# ---
FROM registry-1.docker.io/library/busybox:1.36.1

ARG TARGETOS TARGETARCH
ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static-$TARGETARCH /tini
RUN chmod +x /tini

RUN mkdir /app
RUN addgroup -S app && adduser -S -G app app
WORKDIR /app

COPY --from=builder /out/supervisor .
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

ADD https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.13/grpc_health_probe-$TARGETOS-$TARGETARCH /bin/grpc_health_probe

RUN chmod +x /bin/grpc_health_probe

RUN chown -R app:app .
USER app

EXPOSE 3000

ENTRYPOINT [ "/tini", "--", "./supervisor" ]
