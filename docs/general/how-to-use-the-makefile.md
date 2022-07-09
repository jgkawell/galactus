
# How to use the Makefile

Most everything you need to do in `galactus` is automated through the main `Makefile` in the root of the repository. See below for examples on how to use it.

## Local Development

Working on services locally is sometimes needed, we have a few dependencies that are required by the framework.

(RabbitMQ, Postgres, Mongo). Running all the se dependencies locally is automated with a few `make` targets.

Below is an example workflow:

```sh
# deploy core infrastructure with docker
make local

# deploy the services you are working on (user service in the example)
# NOTE: This needs to be run from a differnt terminal window
make service NAME=user

# once your done, clean up your services
make clean-local
```

## Remote Development

Below are examples of different ways to use the `Makefile` to automate your workflow:

```sh
# to update all services to latest by pulling remote images from Docker Hub
# that were built by piplines off of main. this means you don't have to build
# all services locally, thus accelerating deployment time
git checkout main
git pull main
make remote

# launch a service into your own namespace
make asset

# launch all services
make all

# launch a service WITHOUT argo rollouts (not recommended)
ROLLOUTS_ENABLED=false make asset

# launch a service and have argo run INTEGRATION tests
HAS_INTEGRATION_TESTS=true make asset

# launch a service and have argo run FUNCTIONAL tests
HAS_FUNCTIONAL_TESTS=true make asset

# launch a service by name on your local machine
make service NAME=eventstore

# launch local development system (ie. `rabbitMQ` and `mongo`)
make local

# clean up local development system.
make clean-local
```

## Helm charts

```sh
# render the template for a service to check the output
make template SVC=asset NAMESPACE=dev OVERRIDE_VALUES=dev BUILD_ID=25776
```
