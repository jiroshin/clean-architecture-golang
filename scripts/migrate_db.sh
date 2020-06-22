#!/bin/bash

source "./scripts/$ENV_FILE"

if ! [ -x "$(command -v migrate)" ]
then
    echo "Migrate tool likely not installed, follow instruction in README.md"
else
  migrate -verbose -path ./migrations/ -database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" up
fi
