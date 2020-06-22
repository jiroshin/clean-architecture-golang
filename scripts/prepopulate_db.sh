#!/bin/bash

source "./scripts/$ENV_FILE"

if ! [ -x "$(command -v psql)" ]
then
    echo "psql likely not installed, follow instruction in README.md"
else
    cat ./fixtures/01-test-fixture.sql | psql "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"
fi
