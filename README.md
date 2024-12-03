# galactus

**NOTE**: `galactus` has been superceded by my work on the [Draft framework](https://github.com/steady-bytes/draft). I suggest following the work there as it has all the same value as `galactus` and more!

This project is still in a pre-alpha state. Basic functionality when running locally will mostly work, but MANY features are not yet implemented and some aspects may be completely non-functional. Some future features will include:

- Better documentation/"getting started" pages
- "Easy" remote setup using Terraform
- Example consumers
- Example client
- Domain support
- Hasura as a generic query handler
- CI/CD with Docker images hosted in Docker Hub

## Overview

The `galactus` repository is a place to hold the "galaxy" of microservices, modules, tools, and clients that make up a modern application. It is based off of [this](https://github.com/golang-standards/project-layout) Github project which defines a generic project layout for Go development and has been expanded to support clients, Kubernetes deployments, and much, much more.

It is based off of the internal platform used by Circadence but heavily altered to be as generic as possible. This project can be used as a template or starting point for a new project or you can take pieces of it as you find them useful. Please note that the project is still in very early days as an open-source framework and is not yet ready for use except by the most daring. Stay tuned for an official public release.

TL;DR: `galactus` is a generic development framework for modern applications.

## How To

### Requirements

The absolute basic tooling you need need to develop with `galactus` is:

- `docker`: install instructions [here](https://docs.docker.com/get-docker/)
- `go` (*1.19+*): it is suggested to use [gvm](https://github.com/moovweb/gvm)

### Getting started

First install the `gctl` tool:

```shell
cd tools/gctl
go install
```

Now move to the root of the repository and initialize the `gctl` configuration:

```shell
cd ../..
gctl config
```

You're now ready to initialize and start your local infra (databases, brokers, etc.):

```shell
gctl infra init
gctl infra start
```

Once all the infra is running, you can start your core services:

```shell
# tip: this will take a bit to run the first time, but will be quick after that
gctl run core
```

Everything is now running! To learn more, try out this demo: **TODO**

### Structure

A quick overview of the directories is below:

- `api/`: the Protobuf definitions and generated API code
- `client/`: the client code for the application
- `data/`: data files like configs, binaries, etc.
- `deployments/`: Helm charts, configs, values, etc.
- `docs/`: wikis, diagrams, etc.
- `examples/`: templates, example configs, etc.
- `functions/`: cloud function apps like Azure Functions, AWS Lambda, etc.
- `infrastructure/`: infrastructure code like Terraform, Cloud Formation, etc.
- `pipelines/`: pipelines configuration for CI/CD
- `pkg/`: common Go modules used across different services and tools
- `scripts/`: helpful scripts to use during local dev/debugging
- `services/`: microservice directories each in their own folder
- `test/`: functional and integration tests
- `third_party/`: third party libraries like git submodules
- `tools/`: CLI tools like pipeline binaries, operational tools, etc.

### Docs

Here is a table of contents of the docs:

- Agreements
  - [Go Logger](docs/agreements/go-logger.md)
- General
  - [How to create a new service](docs/general/how-to-create-a-microservice.md)
  - [How to create a PR](docs/general/how-to-create-a-pr.md)
  - [How to use the Makefile](docs/general/how-to-use-the-makefile.md)
- Go
  - [How to create a new Go module](docs/go/how-to-create-a-go-module.md)
  - [How to develop Go modules](docs/go/how-to-develop-go-modules.md)
  - [How to profile a Go service](docs/go/how-to-profile-a-go-service.md)
- Testing
  - [How to write unit tests](docs/testing/how-to-write-unit-tests.md)
