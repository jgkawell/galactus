FROM --platform=linux/amd64 bufbuild/buf:1.4.0 AS buffer

# Install dependencies
RUN apk add --update --no-cache ca-certificates curl npm protoc
RUN apk add --update --no-cache go=1.19.3-r0 --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community

# Set go env
ENV GOPATH=$HOME/go
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
RUN mkdir -p $GOBIN

# Download ts-protoc-gen
RUN npm config set unsafe-perm true && \
    npm install -g ts-protoc-gen@0.15.0 grpc-tools@1.11.1 @grpc/grpc-js@1.4.4 && \
    npm cache clean --force

RUN chmod +x /usr/local/lib/node_modules/ts-protoc-gen/bin/protoc-gen-ts

# Download and set protoc-gen-web
RUN curl -L -o protoc-gen-grpc-web https://github.com/grpc/grpc-web/releases/download/1.3.1/protoc-gen-grpc-web-1.3.1-linux-x86_64
RUN chmod +x ./protoc-gen-grpc-web
RUN mv protoc-gen-grpc-web $GOBIN/

# install protoc plugins
RUN go install -mod=readonly github.com/favadi/protoc-go-inject-tag@v1.3.0
RUN go install -mod=readonly github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.10.0
RUN go install -mod=readonly google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
RUN go install -mod=readonly google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
RUN go install -mod=readonly github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1
RUN go install -mod=readonly github.com/circadence-official/protoc-gen-validate@v0.6.8
RUN go install -mod=readonly github.com/circadence-official/protoc-gen-gorm@v1.1.3
