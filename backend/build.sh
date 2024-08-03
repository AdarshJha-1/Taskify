#!/bin/sh

# Create bin directory if it doesn't exist
mkdir -p bin

# Build the project
go build -o bin/taskify-backend ./cmd/

echo "Build completed. Binary is located in bin/taskify-backend"

# Running build file
echo 
echo "<-------------------------Output--------------------------------->"
echo 
./bin/taskify-backend