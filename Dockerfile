# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.25.4 AS builder
WORKDIR /app

# Copy go.mod first (better caching)
COPY go.mod ./
RUN go mod download

# Copy source
COPY . .

# Build the binary
RUN go build -o deploy-tool .

# Runtime stage
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/deploy-tool .

# Run the binary
ENTRYPOINT ["./deploy-tool"]