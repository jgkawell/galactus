#!/bin/bash

# check if mockery is installed, install if not
if ! mockery --version > /dev/null 2>&1 ; then
    echo "Mockery is not installed. Installing..."
    ./scripts/install-mockery.sh
fi

AGGREGATE=$1
VERSION=$2
if [ -z "${VERSION}" ]; then
    VERSION=1
fi
DOMAIN=$3
if [ -z "${DOMAIN}" ]; then
    DOMAIN="atlas"
fi
# Linux has different sed flags than Darwin (OSX)
OS=$(uname)

echo "Aggregate: ${AGGREGATE}"
echo "Version: ${VERSION}"
echo "Domain: ${DOMAIN}"

# copy over proto files
cp api/gen/go/${DOMAIN}/${AGGREGATE}/v${VERSION}/*.pb.go pkg/chassis/clientfactory/mockclient

# move into the mockclient directory
cd pkg/chassis/clientfactory/mockclient || exit 1

# update the package name in the proto filies
if [ "${OS}" = "Linux" ]; then
    sed -i "s/package v${VERSION}/package mockclient/g" ./*.pb.go
else
    sed -i '' "s/package v${VERSION}/package mockclient/g" ./*.pb.go
fi

# remove previously generated mock files
rm mock_${AGGREGATE}_client.go
rm mock_${AGGREGATE}_server.go

# run mockery to generate new files (THIS REQUIRES mockery v2.X.X)
mockery --all --inpackage --case underscore

# remove proto files
rm ./*.pb.go

# update the references to the proto files in the generated mock files
if [ "${OS}" = "Linux" ]; then
    # add the import to the proto
    sed -i "s/import (/import ( \n\t${DOMAIN}_${AGGREGATE}_v${VERSION} \"github.com\/circadence-official\/galactus\/api\/gen\/go\/${DOMAIN}\/${AGGREGATE}\/v${VERSION}\"/g" mock_${AGGREGATE}_client.go
    sed -i "s/import (/import ( \n\t${DOMAIN}_${AGGREGATE}_v${VERSION} \"github.com\/circadence-official\/galactus\/api\/gen\/go\/${DOMAIN}\/${AGGREGATE}\/v${VERSION}\"/g" mock_${AGGREGATE}_server.go
    # update all references to * to instead target the proto
    sed -i "s/\*/\*${DOMAIN}_${AGGREGATE}_v${VERSION}\./g" mock_${AGGREGATE}_client.go
    sed -i "s/\*/\*${DOMAIN}_${AGGREGATE}_v${VERSION}\./g" mock_${AGGREGATE}_server.go
    # fix the reference to * on the struct
    sed -i "s/_m \*${DOMAIN}_${AGGREGATE}_v${VERSION}\./_m \*/g" mock_${AGGREGATE}_client.go
    sed -i "s/_m \*${DOMAIN}_${AGGREGATE}_v${VERSION}\./_m \*/g" mock_${AGGREGATE}_server.go
else
    # add the import to the proto
    sed -i '' "s/import (/import ( \n\t ${DOMAIN}_${AGGREGATE}_v${VERSION} \"github.com\/circadence-official\/galactus\/api\/gen\/go\/${DOMAIN}\/${AGGREGATE}\/v${VERSION}\"/g" mock_${AGGREGATE}_client.go
    sed -i '' "s/import (/import ( \n\t ${DOMAIN}_${AGGREGATE}_v${VERSION} \"github.com\/circadence-official\/galactus\/api\/gen\/go\/${DOMAIN}\/${AGGREGATE}\/v${VERSION}\"/g" mock_${AGGREGATE}_server.go
    # update all references to * to instead target the proto
    sed -i '' "s/\*/\*${DOMAIN}_${AGGREGATE}_v${VERSION}\./g" mock_${AGGREGATE}_client.go
    sed -i '' "s/\*/\*${DOMAIN}_${AGGREGATE}_v${VERSION}\./g" mock_${AGGREGATE}_server.go
    # fix the reference to * on the struct
    sed -i '' "s/_m \*${DOMAIN}_${AGGREGATE}_v${VERSION}\./_m \*/g" mock_${AGGREGATE}_client.go
    sed -i '' "s/_m \*${DOMAIN}_${AGGREGATE}_v${VERSION}\./_m \*/g" mock_${AGGREGATE}_server.go
fi
