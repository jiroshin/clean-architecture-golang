#!/bin/bash

source "./scripts/$ENV_FILE"

docker-compose down --volumes
docker-compose up --detach

echo "Waiting for db to be ready"
#To avoid double restart issue https://github.com/docker-library/postgres/issues/146#issuecomment-349290569
NETWORK=$(docker network ls | grep "${COMPOSE_PROJECT_NAME}_default" | awk '{print $1}')
PID=$(docker ps | grep "${COMPOSE_PROJECT_NAME}_db_1" | awk '{print $1}')

until docker run --rm --link "$DOCKER_IMAGE" --net "$NETWORK" "$DOCKER_IMAGE" pg_isready -U "$DB_USER" -h "$PID"; do sleep 1; done