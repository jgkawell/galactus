# hammer

> "I suppose it is tempting, if the only tool you have is a hammer, to treat everything as if it were a nail." - Abraham Maslow

`hammer` is a simple CLI tool to generate microservices in the `galactus` framework. The idea for it is based off the above quote from Maslow which epitomizes the [Law of the instrument](https://en.wikipedia.org/wiki/Law_of_the_instrument). Basically, in our case, we always use microservices (nails) to solve all our problems and thus `hammer` is our tool we always use to generate these microservices.

## Overview

There are three different types of services you can create with `hammer`:
- `aggregate`: A service that represents an Aggregate in our data model. This service implements a CRUL (CRUDL without Delete) gRPC interface as well as listens for Create and Update commands on the same Aggregate from the `messagebus`. Each time you create a new Aggregate under under `/api` you'll want to create a matching service using this command.
- `consumer`: A service that consumes certain **asynchronous** commands from the `messagebus` and executes them on a certain Aggregate. Create this when you need a command handled asynchronously for a certain Aggregate.
- `rpc`: A service that implements a gRPC interface to run **synchronous** commands on a certain Aggregate. Create this when you need a command handled synchronously for a certain Aggregate.

## Usage

**NOTE 1**: Run all the below commands from `hammer` directory.

**NOTE 2**: Note the capitalization of the aggregate, service, and command names. These are important and should be Pascal Case.

### Creating services

To create a new `aggregate` service for an aggregate named "Football", run the following command:

```sh
go run main.go aggregate --aggregate_name="Football" --output_path="../../internal"
```

To create a new `consumer` service for an aggregate named "Football" that will implement the command "PauseGame", run the following command:

```sh
# NOTE: the capitalization of the service_name is important
go run main.go consumer --service_name="GameOfficiating" --aggregate_name="Football" --command_name="PauseGame" --output_path="../../internal"
```

To create a new `rpc` service for an aggregate named "Football" that will implement the command "PauseGame", run the following command:

```sh
# NOTE: the capitalization of the service_name is important
go run main.go rpc --service_name="GameOfficiating" --aggregate_name="Football" --command_name="PauseGame" --output_path="../../internal"
```
