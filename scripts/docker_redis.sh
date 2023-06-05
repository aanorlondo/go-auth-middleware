#!/bin/bash

# prepare init.sh & redis.conf templates
sh prepare_redis_init.sh

# clean first
docker rm -f REDIS-RBAC-LOCAL
docker rmi redis/redis-stack:latest

# dynamically configure redis to be in the same network as the host
host_ip=$(ipconfig getifaddr en0)

# build the proxy
docker build -t negan/redis-rbac:local ../db/redis

# push the image to remote
docker push redis/redis-stack:latest

# run redis
VOLUME_ABSPATH=$(readlink -f "../db/redis/local-persistence")
CONFIG_ABSPATH=$(readlink -f "../db/redis/redis.conf")
docker run \
    -d \
    -v ${VOLUME_ABSPATH}:/data \
    -v ${CONFIG_ABSPATH}:/redis-stack.conf \
    -p 6379:6379 \
    -p 8001:8001 \
    --add-host="host:${host_ip}" \
    --name REDIS-RBAC-LOCAL \
    redis/redis-stack:latest