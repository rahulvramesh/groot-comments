#!/bin/bash
#@author : rahul@rahulvramesh.me

##### Commands Decumentation
# sh run.sh r
## starts the main function with go run 
# sh run.sh d
## deploy to development
# sh run.sh s
## deploy to staging
# sh run.sh p
## deploy to production
# sh run.sh h
## run with hot reload


if [ "$1" = "r" ]; then
   HOSTNAME='aka' DATABASE_CONNECTION_STRING='postgres://postgres:@localhost:32768/comments_groot?sslmode=disable' go run main.go
fi
