# functional tests

The purpose of these tests are to check functionality of the galactus services from the client perspective. These are scripts wrapped in dockerfiles that utilize cli tools to test functionality.

## How to deploy

Deploy a service with the test you want with the correct values set using the `Makefile` from the root of `galactus`. For example:

```bash
HAS_FUNCTIONAL_TESTS=true make eventstore
```

This will build and deploy the service and tests into your personal dev namespace.

**NOTE**: You'll need the secret `argo-testing` k8s secret in your namespace.

## How to dev

### Setup

Copy the `config/example.*.env` files to `config/*.env` and edit the `CHANGEME` values as needed.

### Running the tests

Simply run the test script from the `test/functional` directory with the `RUN_LOCAL` variable:

```bash
RUN_LOCAL=true ./TEST_FOLDER/TEST_SCRIPT.sh
```
