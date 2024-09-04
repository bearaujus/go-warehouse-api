#!/bin/bash

# Check if the .env file is provided as an argument
if [ -z "$1" ]; then
    echo "Usage: $0 path to .env file"
    exit 1
fi

# Load environment variables from the provided .env file
export $(grep -v '^#' "$1" | xargs)

# Remove the existing log directory and create a new one
rm -rf "${LOCAL_LOG_PATH}"
mkdir -p "${LOCAL_LOG_PATH}"

# Iterate over all environment variables that end with _CONTAINER_NAME
for var in $(compgen -A variable | grep '_CONTAINER_NAME$'); do
    # Get the container name from the variable
    container_name="${!var}"

    # Create the corresponding log file
    touch "${LOCAL_LOG_PATH}${container_name}${LOG_EXTENSION}"

    echo "Created log file: ${LOCAL_LOG_PATH}${container_name}${LOG_EXTENSION}"
done
