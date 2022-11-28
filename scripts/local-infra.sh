#!/bin/bash

command=$1

# if command == start
if [ "$command" == "start" ]; then
    echo "Starting local infra"

    # mongodb
    docker run --name galactus-nosql-db \
        -d -it \
        -p 27017:27017 \
        mongo:4.2.21

    # postgresql
    docker run --name galactus-sql-db \
        -d -it \
        -e POSTGRES_PASSWORD=admin \
        -e POSTGRES_USER=admin \
        -e POSTGRES_DB=dev \
        -p 5434:5432 \
        postgres:14.4

    # rabbitmq
    docker build --rm . -f ./third_party/rabbitmq/Dockerfile -t broker
    docker run --name galactus-broker \
        -d -it \
        -p 15672:15672 \
        -p 5672:5672 \
        broker

    # hasura
    docker run --name galactus-queryhandler \
        -d -it \
        -p 8082:8080 \
        -e HASURA_GRAPHQL_METADATA_DATABASE_URL='postgres://admin:admin@host.docker.internal:5434/dev' \
        -e PG_DATABASE_URL='postgres://admin:admin@host.docker.internal:5434/dev' \
        -e HASURA_GRAPHQL_ENABLE_CONSOLE='true' \
        -e HASURA_GRAPHQL_ENABLED_LOG_TYPES='startup, http-log, webhook-log, websocket-log, query-log' \
        -e HASURA_GRAPHQL_ENABLE_TELEMETRY='false' \
        hasura/graphql-engine:v2.15.2

    # display results
    docker ps

    echo "Local infra started"
    exit 0
fi

if [ "$command" == "stop" ]; then
    echo "Stopping local infra"

    # mongodb
    docker stop galactus-nosql-db

    # postgresql
    docker stop galactus-sql-db

    # rabbitmq
    docker stop galactus-broker

    # envoy proxy
    docker stop galactus-queryhandler

    echo "Local infra stopped"
    exit 0
fi
