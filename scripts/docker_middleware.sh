#!/bin/bash

# read en vars
source prepare_env.sh

# clean first
docker rm -f GO-AUTH-LOCAL
docker rmi negan/go-auth:local

# dynamically configure the nginx proxy to be in the same network than the host
host_ip=$(ipconfig getifaddr en0)

# build the proxy
docker build -t negan/go-auth:local ../app

# push the image to remote
docker push negan/go-auth:local

# run the proxy
docker run \
    -d \
    -p 3456:3456 \
    -e DATABASE_HOSTNAME=${DATABASE_HOSTNAME}\
    -e DATABASE_PORT=${DATABASE_PORT}\
    -e DATABASE_NAME=${DATABASE_NAME}\
    -e DATABASE_TABLENAME=${DATABASE_TABLENAME}\
    -e DATABASE_USERNAME=${DATABASE_USERNAME}\
    -e DATABASE_PASSWORD=${DATABASE_PASSWORD}\
    -e REDIS_HOSTNAME=${REDIS_HOSTNAME}\
    -e REDIS_PORT=${REDIS_PORT}\
    -e REDIS_PASSWORD=${REDIS_PASSWORD}\
    -e APP_SECRET_KEY=${APP_SECRET_KEY}\
    --add-host="host:${host_ip}" \
    --name GO-AUTH-LOCAL \
    negan/go-auth:local