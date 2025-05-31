# syntax=docker/dockerfile:1

ARG MODULE=github.com/iamgaru/human-presence-sim

FROM golang:1.21-alpine AS builder
ENV GO111MODULE=on
WORKDIR /src

# Copy and initialize module
COPY go.mod go.sum ./
RUN go mod init $MODULE || true
RUN go mod tidy

COPY . .

RUN go build -o simulator ./cmd/simulator

FROM alpine:latest
COPY --from=builder /src/simulator /usr/local/bin/simulator
ENTRYPOINT ["simulator"]
CMD ["--help"]
