# syntax=docker/dockerfile:1

FROM golang:1.23.4 AS test

# Install git (needed for Go modules)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Run tests
RUN go test ./... -v
