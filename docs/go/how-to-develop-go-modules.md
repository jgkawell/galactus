# Developing Go modules

## Overview

We used to have a system where we would [replace](https://golang.org/ref/mod#go-mod-file-replace) within our `go.mod` files with the target being the local go module under the `/pkg` folder. This meant we could develop quickly since all services/tools used the same version of all modules and could all be updated concurrently.

However, as our services grew we quickly ran into [dependency hell](https://en.wikipedia.org/wiki/Dependency_hell) trying to manage tens of services all depending on the same core files which were in constant flux. To fix this, we moved to **remote dependencies**.

Doing this was as simple as versioning, building, and `git` tagging them since `go mod` uses `git` under the covers anyway.. Our services now reference all the modules under `/pkg` as if they were 3rd party dependencies and download them from Azure DevOps using `git`.

This does mean that developing our modules is slightly more involved than before, **but** it improves our overall development velocity by making all our modules versioned and keeping developers from tripping over each other when making module changes.

## Creating new module

If you want to create a new Go module, _go_ [here](how-to-create-a-go-module.md)

## Modifying an existing module

Let's assume you want to update the `logging` module. In order to do this you can simply open that directory `/pkg/logging` and start editing it as you would any go module.

To test out your changes you _could_ create a playground application to test run it, but even better you can modify an existing service to inherit it with a local replace and test out the updated module in a full remote namespace. Let's use an example of testing out `logging` using the `asset` service:

### Update the go.mod

Update the file `/services/core/eventstore/v1/go.mod` to include the replace call on the `logging` module (update versions accordingly):

```go.mod
module asset

go 1.19

require (
    ...
    github.com/jgkawell/galactus/pkg/logging/v2 v2.0.4
    ...
)

require (
    # indirect dependencies
    ...
)

replace github.com/jgkawell/galactus/pkg/logging/v2 v2.0.4 => ../../pkg/logging

```

You can stop here if you're only going to build/run the service locally. For remote testing...

### Update the Dockerfile

Update the file `/services/core/eventstore/v1/Dockerfile` so that you are copying over the `/pkg/logging` directory:

```Dockerfile
# We want to populate the module cache based on the go.{mod,sum} files.
COPY ./pkg/logging ./pkg/logging
COPY ./services/core/eventstore/v1/go.mod ./services/core/eventstore/v1/go.mod
COPY ./services/core/eventstore/v1/go.sum ./services/core/eventstore/v1/go.sum
```

Now you can run the service remotely using: TODO

## Open a PR

Open your PR to `main` following the PR guidelines [here](./creating-a-pr.md).
