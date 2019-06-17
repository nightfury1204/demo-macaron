#!/usr/bin/env bash

curl -X PUT -d '{"name": "lord of rings", "author": "me", "price": 5000, "description":"this is a book"}' http://localhost:4000/books/2147483647

curl -X POST -d '{"name": "lord of rings", "author": "me", "price": 1000, "description":"this is a book"}' http://localhost:4000/books

curl -X DELETE http://localhost:4000/books/5577006791947779410