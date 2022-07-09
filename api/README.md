# About

This directory is a thin wrapper around [buf](https://docs.buf.build/generate/usage) and is used to generate several artifacts the `.proto` files within the `api/protos/` directory. The generated code is meant to be consumed by galactus services and clients:

- [GORM binding types](https://gorm.io/docs/models.html) for easy SQL database integration.
- validators using [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)
- [Swagger](https://swagger.io/) documentation based on comments in the protos.
- grpc bindings to Go and grpc-web

## Quick start

*Only if you need to change a .proto* `make clean && make build` and then you should check in your interface protos and the generated code with your PR. If you are just consuming the generated goodness just link your service project (described below) and enjoy.

## System Requirements

- MacOS or Linux (verified using Ubuntu 20.04)
- Docker
- GNU Make

## Structure

Generated code is generated and checked in the `gen` directory. All generated artifacts are placed under there into directories bucketed by language. This is dictated by buf and configured in the `buf.gen.yaml` file.

### Makefiles

The top level Makefile is set up to call more specialized Makefiles in the generated directories. By default the Makefile should contain a `clean` which removes generated sources for that language and the default target should be the target that does any building and verifying of the generated sources. The name doesn't matter, but `build` is a reasonable choice.
