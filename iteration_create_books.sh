#!/bin/bash

for i in {1..10}; do
    title="Test Book$i"
    author="Test Author"
    curl -X POST -H "Content-Type: application/json" -d "{\"title\": \"$title\", \"author\": \"$author\"}" http://localhost:8080/books
done
