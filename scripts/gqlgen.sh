#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f pkg/gql/generated.go \
    pkg/gql/models/generated.go \
    pkg/gql/resolvers/generated.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"