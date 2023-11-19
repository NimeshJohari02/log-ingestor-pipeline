#!/bin/bash

# List of possible error messages
error_messages=("Application Log : NullPointerException" "Application Log : Details Not Found" "Application Log : Can't find customer" "Application Log: Nimesh's Log Injestor")

while true; do
    # Select a random error message
    random_error=${error_messages[$((RANDOM % ${#error_messages[@]}))]}

    # Generate random values for timestamp, commit, and span
    timestamp=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
    commit=$(uuidgen)
    span=$(uuidgen)

    # JSON body with random values and error message
    json_body='{
        "level": "error",
        "message": "'"$random_error"'",
        "resourceId": "server-1234",
        "timestamp": "'"$timestamp"'",
        "traceId": "abc-xyz-123",
        "spanId": "'"$span"'",
        "commit": "'"$commit"'",
        "metadata": {
            "parentResourceId": "server-0987"
        }
    }'

    # Send POST request to localhost:3000
    curl -X POST -H "Content-Type: application/json" -d "$json_body" http://localhost:3000 &

    # Introduce a delay to achieve 100 requests per second
    sleep 0.1
done
