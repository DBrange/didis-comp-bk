# Build stage
FROM golang:1.22.3-alpine3.20 AS builder

# Instalar git y otras dependencias necesarias
RUN apk add --no-cache git openssh-client upx

# Configurar SSH para usar la clave privada
RUN mkdir -p /root/.ssh && chmod 700 /root/.ssh
COPY keys/ssh.pub /root/.ssh/id_rsa.pub
COPY keys/ssh /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

# Configurar Git para usar SSH
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

WORKDIR /app

# Copiar los archivos necesarios
COPY go.mod go.sum ./
RUN go mod download -x

COPY . .

# Construir la aplicación
RUN go build \
    -ldflags="-s -w" \
    -o app -v ./cmd

RUN upx app

# Final stage
FROM alpine:3.20

LABEL Name=didis-comp-bk

RUN apk update && apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT [ "./app" ]