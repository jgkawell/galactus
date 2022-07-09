# How to create a *new* Go module

NOTE: In the below examples we will be creating a module called `basketball`.

## Set up the new module

Create the new directory for your module:

```shell
mkdir pkg/basketball
```

Move into your new directory:

```shell
cd pkg/basketball
```

Initialize your new module:

```shell
go mod init github.com/circadence-official/galactus/pkg/basketball
```

## Write the module code

Write out all your code for your module according to business needs.

## Tests

Write your unit tests along side your module code: [How to write unit tests](../testing/how-to-write-unit-tests.md)

## Open a PR

Open your PR to `main` following the PR guidelines [here](./creating-a-pr.md).