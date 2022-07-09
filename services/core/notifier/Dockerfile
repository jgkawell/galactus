# Set base image
FROM golang:1.17 AS builder

# Install xz
RUN apt-get update && apt-get install -y xz-utils \
    && rm -rf /var/lib/apt/lists/*

# Install UPX
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

# Project setup
WORKDIR /bin


# We want to populate the module cache based on the go.{mod,sum} files.
COPY ./service/core/notifier/go.mod ./service/core/notifier/
COPY ./service/core/notifier/go.sum ./service/core/notifier/

# Change to the service directory
WORKDIR /bin/internal/notifier

# Fetch go modules
RUN go mod download
RUN go mod verify

# Grab that code
COPY ./service/core/notifier .

# Run test codes
# Adding `-race` flag to the test so any potential race conditions will be identified and
# published.
RUN go test -race -gcflags=-l -test.v ./...

# Build the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go

# Strip and compress the binary
RUN strip --strip-unneeded main
RUN upx main

FROM alpine:3.14 as alpine
RUN apk add -U --no-cache ca-certificates

# Copy the binary from builder
COPY --from=builder /bin/internal/notifier/main /etc/main

# setting context to run main
WORKDIR /etc

# Run
ENTRYPOINT ["/etc/main"]

EXPOSE 8080
EXPOSE 8090
