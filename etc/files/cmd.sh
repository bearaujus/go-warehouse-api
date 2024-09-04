#!/bin/bash

# Function to display help message
help_message() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -S, --start, start        Start the application (default action if no option is provided)"
    echo "  -s, --stop, stop          Stop the application"
    echo "  -h, --help, help          Display this help message"
    echo ""
}

# Function to start the application
start() {
    # Check if the .env file exists in the same directory as the script
    if [ -f ".env" ]; then
        echo ".env file found."
        # Load the environment variables from .env
        export $(grep -v '^#' ".env" | xargs)
    else
        echo ".env file not found. Exiting."
        exit 1
    fi

    mkdir -p "${LOCAL_LOG_PATH}"

    # Iterate over all environment variables that end with _CONTAINER_NAME
    for var in $(compgen -A variable | grep '_CONTAINER_NAME$'); do
        # Get the container name from the variable
        container_name="${!var}"

        # Check if the log file already exists
        log_file="${LOCAL_LOG_PATH}${container_name}${LOG_EXTENSION}"
        if [ ! -f "$log_file" ]; then
            touch "$log_file"
            echo "Created log file: $log_file"
        else
            echo "Log file already exists: $log_file"
        fi
    done
    docker compose down -v || true
    echo "Starting application..."
    docker compose up -d
}

# Function to stop the application
stop() {
    echo "Stopping application..."
    docker compose down -v
}

# Default action is to start the application
action="start"

# Check if any arguments were passed
if [ $# -gt 0 ]; then
    case "$1" in
        -s|--stop|stop)
            action="stop"
            ;;
        -S|--start|start)
            action="start"
            ;;
        -h|--help|help)
            help_message
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            help_message
            exit 1
            ;;
    esac
fi

# Execute the selected action
$action
