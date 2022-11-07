FROM golang:1.19.3-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

# ---------------------------------------------------------
FROM ghcr.io/squarefactory/slurm:22.05.3-1-2-login-rocky9.0
# ---------------------------------------------------------

COPY --from=builder /build/app /usr/bin/provider-ssh-authorized-keys

RUN chown root:root /usr/bin/provider-ssh-authorized-keys
