#!/bin/bash

command=$1

if [ "$command" == "start" ]; then
    echo "Starting core services"

    cd ./services/core/registry || exit 1
    echo "Starting registry"
    go run main.go &
    REGISTRY_PID=$!
    cd ../../..

    # wait until registry returns a 200 OK on it's health endpoint before starting other services
    while true; do
        status=$(curl -s -o /dev/null -w '%{http_code}' http://localhost:35000/health)
        if [ "$status" == "200" ]; then
            echo "Registry is up and running"
            break
        fi
        sleep 1
    done

    cd ./services/core/commandhandler || exit 1
    echo "Starting commandhandler"
    go run main.go &
    COMMANDHANDLER_PID=$!
    cd ../../..

    cd ./services/core/eventstore || exit 1
    echo "Starting eventstore"
    go run main.go &
    EVENTSTORE_PID=$!
    cd ../../..

    cd ./services/core/notifier || exit 1
    echo "Starting notifier"
    go run main.go &
    NOTIFIER_PID=$!
    cd ../../..

    # TODO: run Hasura (queryhandler) locally

    # keep running until user cancels (ctrl-c), then kill all processes
    wait $REGISTRY_PID $COMMANDHANDLER_PID $EVENTSTORE_PID $NOTIFIER_PID
fi

if [ "$command" == "stop" ]; then
    echo "Stopping core services"

    # NOTE: this will kill all process called `main`. probably want to find a better way to do this.
    pkill main

    echo "Core services stopped"
    exit 0
fi
