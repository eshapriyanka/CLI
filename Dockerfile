# Stage 1: Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o netcheck main.go

# Stage 2: Run (Alpine Linux)
FROM alpine:latest
# Install certificates so we can make HTTPS requests
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/netcheck .
ENTRYPOINT ["./netcheck"]