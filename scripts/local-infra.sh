#!/bin/bash

command=$1

# if command == start
if [ "$command" == "start" ]; then
    echo "Starting local infra"

    # mongodb
    docker run --name galactus-nosql-db \
        -d --rm -it \
        -p 27017:27017 \
        mongo:4.2.21

    # postgresql
    docker run --name galactus-sql-db \
        -d --rm -it \
        -e POSTGRES_PASSWORD=admin \
        -e POSTGRES_USER=admin \
        -e POSTGRES_DB=dev \
        -p 5434:5432 \
        postgres:14.4

    # rabbitmq
    docker build --rm . -f ./third_party/rabbitmq/Dockerfile -t broker
    docker run --name galactus-broker \
        -d --rm -it \
        -p 15672:15672 \
        -p 5672:5672 \
        broker

    # envoy proxy
    docker build --rm . -f ./third_party/proxy/Dockerfile -t proxy
    docker run --name galactus-proxy \
        -d --rm -it \
        --net=host \
        -p 10000:10000 \
        proxy

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
    docker stop galactus-proxy

    echo "Local infra stopped"
    exit 0
fi
