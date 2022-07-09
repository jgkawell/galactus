#!/bin/bash

OS=$(uname)
VERSION="2.9.4"
TEMP_PATH="./mockery_install_temp_folder"

# make staging folder
mkdir "${TEMP_PATH}"

# download and extract mockery
curl -s -L "https://github.com/vektra/mockery/releases/download/v${VERSION}/mockery_${VERSION}_${OS}_x86_64.tar.gz" | tar xvz -C "${TEMP_PATH}"

# move mockery into path
mv "${TEMP_PATH}/mockery" "${GOPATH}/bin/"

# clean up
rm -r "${TEMP_PATH}"
