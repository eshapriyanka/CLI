# --- STAGE 1: The Builder (Heavy) ---
# We use the full Go image (~800MB) to compile the code
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
# Compile the binary named "netcheck"
RUN go build -o netcheck main.go

# --- STAGE 2: The Runner (Lightweight) ---
# We switch to Alpine Linux (~5MB) for the final image
FROM alpine:latest
# Install security certificates for HTTPS
RUN apk add --no-cache ca-certificates
WORKDIR /root/
# Copy ONLY the compiled binary from Stage 1
COPY --from=builder /app/netcheck .

# Set the binary as the entrypoint
ENTRYPOINT ["./netcheck"]