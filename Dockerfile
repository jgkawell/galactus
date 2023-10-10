ARG GO_VERSION=1.21
ARG ALPINE_VERSION=3.18

# This layer builds the binary
FROM golang:${GO_VERSION} AS builder

# Install build dependencies
RUN apt update && apt install -y upx-ucl && rm -rf /var/lib/apt/lists/*

# Project setup
WORKDIR /app

# Build arguments
ARG DOMAIN
ARG SERVICE
ARG VERSION

# We want to populate the module cache based on the go.{mod,sum} files.
COPY ./services/${DOMAIN}/${SERVICE}/${VERSION}/go.mod ./services/${DOMAIN}/${SERVICE}/${VERSION}/go.mod
COPY ./services/${DOMAIN}/${SERVICE}/${VERSION}/go.sum ./services/${DOMAIN}/${SERVICE}/${VERSION}/go.sum
COPY ./api/gen/go ./api/gen/go
COPY ./pkg ./pkg

# Change to the service directory
WORKDIR /app/services/${DOMAIN}/${SERVICE}/${VERSION}

# Fetch go modules
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod verify

# Grab that code
COPY ./services/${DOMAIN}/${SERVICE}/${VERSION} .

# Run test codes
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build \
    go test -race -test.v ./...

# Build the source
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build \
    GOOS=linux go build -o main ./main.go

# Compress the binary
RUN upx main

# This layer runs the binary
FROM alpine:${ALPINE_VERSION} AS runner

# Install dependencies
RUN apk add -U --no-cache ca-certificates

# Build arguments
ARG DOMAIN
ARG SERVICE
ARG VERSION

# Copy the binary from builder
COPY --from=builder /app/services/${DOMAIN}/${SERVICE}/${VERSION}/main /etc/main

# Set context to run main
WORKDIR /etc

# Run
ENTRYPOINT ["/etc/main"]

# Runtime arguments
ARG HTTP_PORT=8080
ARG GRPC_PORT=8090

# Expose needed ports
EXPOSE ${HTTP_PORT}
EXPOSE ${GRPC_PORT}
