# Stage 1: Build stage
FROM golang:1.23-alpine AS build

ARG PORT=8001

EXPOSE $PORT

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o prometheus-manager .

# Stage 2: Execute
FROM alpine

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/prometheus-manager /app/prometheus-manager

# Set the entrypoint command
ENTRYPOINT ["/app/prometheus-manager"]