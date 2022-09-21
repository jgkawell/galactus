#!/bin/bash

# This script helps automate running the same commands against many subdirectories.
# Using it is simple:
# ./scripts/increment-versions.sh -d PATH [-a ACCEPT_ALL]

while getopts d:a: option; do
    case "${option}" in
    d) directory=${OPTARG} ;; # directory to iterate subdirectories within (e.g. services/)
    a) all=${OPTARG} ;;       # [OPTIONAL] accept all prompts
    *) echo "Unsupported option: ${OPTARG}" ;;
    esac
done

# NOTE: Update this function with anything you need to run within all directories
function something() {
    echo "Doing something..."

    # EXAMPLE: update go mods
    # go mod download
    # go mod verify
    # go mod tidy

    # EXAMPLE: test go compiliation to binary
    # go build main.go
    # rm main

    # EXAMPLE: try to build go module
    # go build

    # EXAMPLE: run gofmt against all files
    # gofmt -l -s -w .

    # EXAMPLE: run golangci-lint
    # golangci-lint run -E gosec -E dupl -E goconst -E gocyclo -E exportloopref -E gofmt -E prealloc --fast
}

# update versions
for d in "${directory}"*/; do
    echo "---------------------------"
    echo "Processing: ${d}"

    # if accept all is not yes, ask for confirmation
    if [[ ${all} =~ ^[Yy]$ ]]; then
        cd "$d" || exit 1
        something
        cd ../..
    else
        read -p "Would you like to update (y/N)? " -n 1 -r
        echo # move to new line
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            cd "$d" || exit 1
            something
            cd ../..
        else
            echo "Skipping..."
        fi
    fi
done
