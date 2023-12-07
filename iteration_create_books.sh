#!/bin/bash

for i in {1..10}; do
    title="Test Book$i"
    author="Test Author"
    curl -X POST -H "Content-Type: application/json" -d "{\"title\": \"$title\", \"author\": \"$author\"}" http://localhost:8080/books
    if [ $? -ne 0 ]; then
        echo "error: need to check the server status"
        exit 1
    fi
done
