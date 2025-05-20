# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o simulator ./cmd/simulator

# Final minimal runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/simulator /usr/local/bin/simulator
ENTRYPOINT ["simulator"]
CMD ["--help"]
