# About

This directory is a thin wrapper around [buf](https://docs.buf.build/generate/usage) and is used to generate several artifacts the `.proto` files within the `api/protos/` directory. The generated code is meant to be consumed by `galactus` services and clients:

- [GORM binding types](https://gorm.io/docs/models.html) for easy SQL database integration.
- validators using [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)
- [Swagger](https://swagger.io/) documentation based on comments in the protos.
- grpc bindings to Go and grpc-web

## System Requirements

- `docker`: install instructions [here](https://docs.docker.com/get-docker/)
- `gctl`: install instructions [here](../tools/gctl/README.md)

## How to use

Everything is automated using the `gctl` tool.

For first setup: `gctl api init`

When needing to rebuild protos: `gctl api build`

## Structure

Generated code is generated and checked in under the `gen/` directory. All generated artifacts are placed under there into directories bucketed by language. This is dictated by `buf` and configured in the `buf.gen.yaml` file.
