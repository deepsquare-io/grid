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
  CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -ldflags "-s -w -X main.version=${BUILDVERSION}" -o /out/provider-ssh-authorized-keys ./

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

COPY --from=builder /out/provider-ssh-authorized-keys .
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

RUN chown -R app:app .
USER app

EXPOSE 3000

ENTRYPOINT [ "/tini", "--", "./provider-ssh-authorized-keys" ]
