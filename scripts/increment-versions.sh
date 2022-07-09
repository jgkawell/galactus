#!/bin/bash

# This script helps automate managing VERSION files. Using it is simple:
# ./scripts/increment-versions.sh -d PATH [-p PLACE -a ACCEPT_ALL]
#
# Simply provide a path (e.g. ./internal or ./internal/users or ./functions, etc.)
# and the script will search all subdirectories for VERSION files to update. It will
# then ask you what place you want to increment (MAJOR=0, MINOR=1, PATCH=2) and ask
# for confirmation of the change, finally updating the VERSION file if requested.
#
# This enables easy incrementing of many services/functions all at once instead of
# manually updating each file when a common directory (e.g. ./pkg) is modified.

while getopts d:p:a: option; do
    case "${option}" in
      d) directory=${OPTARG};; # directory to search for version files under (e.g. ./internal)
      p) default_place=${OPTARG};; # [OPTIONAL] default place (major.minor.patch) version to update
      a) all=${OPTARG};; # [OPTIONAL] accept all prompts
      *) echo "Unsupported option: ${OPTARG}";;
    esac
done

### Increments the part of the semantic version string
### Modified from: https://stackoverflow.com/a/64390598/8431936
## $1: version itself
## $2: number of part: 0 – major, 1 – minor, 2 – patch
increment_version() {
    local delimiter=.
    # shellcheck disable=SC2207
    local array=($(echo "$1" | tr $delimiter '\n'))
    array[$2]=$((array[$2]+1))
    if [ $2 -lt 2 ]; then array[2]=0; fi
    if [ $2 -lt 1 ]; then array[1]=0; fi
    echo "$(local IFS=$delimiter ; echo "${array[*]}")"
}

version_files=( $(IFS=$'\n' find "${directory}" -name "VERSION" | sort) )

# update versions
for version_file in "${version_files[@]}"; do
    echo "---------------------------"
    echo "Processing: ${version_file}"

    # get current position from file
    cur_version=$(cat "${version_file}")

    # if default place not set, get from user
    if [ -z ${default_place+x} ]; then
        read -p "ENTER version place to change (0 – major, 1 – minor, 2 – patch): " -n 1 -r
        echo # move to new line
        if [[ ! $REPLY =~ ^[0-2]+$ ]] ; then
            echo "Invalid place input. Must be 0, 1, or 2."
            exit 1
        else
            place=$REPLY
        fi
    else
        place=${default_place}
    fi

    # calculate new version
    new_version=$(increment_version "${cur_version}" "${place}")
    echo "About to update ${version_file}: ${cur_version} -> ${new_version}"

    # if accept all is not yes, ask for confirmation
    if [[ ${all} =~ ^[Yy]$ ]]; then
        echo "Updating..."
        echo -n "${new_version}" > "${version_file}"
    else
        read -p "Would you like to update (y/N)? " -n 1 -r
        echo # move to new line
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            echo "Updating..."
            echo -n "${new_version}" > "${version_file}"
        else
            echo "Skipping..."
        fi
    fi

    # reset for next loop
    place=

done
