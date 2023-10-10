#!/bin/bash

export VAULT_TOKEN=myroot
export VAULT_ADDR=http://0.0.0.0:8200

vault kv put -mount=secret core/database/postgres/registry url=postgres://admin:admin@localhost:5432/dev?sslmode=disable
vault kv put -mount=secret core/database/mongo/registry url=mongodb://admin:admin@localhost:27017
vault kv put -mount=secret core/messaging/rabbitmq/registry url=amqp://guest:guest@localhost:5672