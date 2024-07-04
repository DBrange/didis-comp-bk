# Build stage
FROM golang:1.22.3-alpine3.20 AS builder

# Installing git, openssh and other necessary dependencies
RUN apk add --no-cache git openssh-client upx

# Configure SSH
RUN mkdir -p /root/.ssh && chmod 700 /root/.ssh
COPY id_rsa /root/.ssh/id_rsa
COPY id_rsa.pub /root/.ssh/id_rsa.pub
RUN chmod 600 /root/.ssh/id_rsa
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

# Configuring Git for use SSH
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

WORKDIR /app

# Copying and downloading dependencies
COPY go.mod go.sum ./
RUN go mod download -x

# Copy the rest of the source code
COPY . .

# Building the application
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

ENTRYPOINT ["./app"]