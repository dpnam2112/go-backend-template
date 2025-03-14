# ───────────────────────────────────────────────
# Stage 1: Build the Go binary
# ───────────────────────────────────────────────
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build the Go binary with debugging enabled
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

# Ensure the binary is executable BEFORE copying it to the next stage
RUN chmod +x main

# ───────────────────────────────────────────────
# Stage 2: Create a Runtime Image with Debugging Tools
# ───────────────────────────────────────────────
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    ca-certificates \
  # Install debugging tools (vim, psql, bash)
    bash \
    vim \
    postgresql-client \
    curl \
    net-tools \
    && rm -rf /var/lib/apt/lists/*

# Set a non-root user for security
RUN addgroup --system appgroup && adduser --system --ingroup appgroup appuser
USER appuser

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Environment configuration
COPY config.env /app/config.env

# Expose the application port
EXPOSE 8080
ENV GIN_MODE=release
CMD ["./main"]

