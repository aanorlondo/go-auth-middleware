#!/bin/bash

# prepare init.sql template
sh prepare_mysql_init.sh

# clean first
docker rm -f MYSQL-AUTH-LOCAL
docker rmi negan/mysql-auth:local

# dynamically configure the mysql to be in the same network as the host
host_ip=$(ipconfig getifaddr en0)

# build the proxy
docker build -t negan/mysql-auth:local ../db/mysql

# push the image to remote
docker push negan/mysql-auth:local

# read en vars
source prepare_env.sh

# run the mysql db
docker run \
    -d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
    --add-host="host:${host_ip}" \
    --name MYSQL-AUTH-LOCAL \
    negan/mysql-auth:local