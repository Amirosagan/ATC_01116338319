#!/bin/bash

# Install swag if not installed
if ! command -v swag &> /dev/null; then
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate swagger documentation
cd ..
swag init -g cmd/server/main.go 