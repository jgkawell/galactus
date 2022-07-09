#!/bin/sh

# move into the generated directory
cd ./gen/go/ || exit 1

# loop over parent directories
echo "Generating go.mod files"
for parentPath in */ ; do
    # loop over package directories
    for packagePath in "$parentPath"*/ ; do
        # remove final slash from package path
        packageName=${packagePath%?}
        echo "module github.com/circadence-official/galactus/api/gen/go/$packageName" > "$packageName"/go.mod
    done
done
echo "Finished"
