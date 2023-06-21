#!/bin/bash

# MYSQL
export MYSQL_ROOT_PASSWORD=root

# AUTH SERVER
export DATABASE_HOSTNAME=host.docker.internal #$(hostname)
export DATABASE_PORT=3306
export REDIS_HOSTNAME=host.docker.internal #$(hostname)
export REDIS_PORT=6379
export APP_SECRET_KEY=azertyuiop12345
export API_ADMIN_TOKEN=admin

# SWAGGER
export HOSTNAME=$(hostname)
input_swagger_file="../app/server/api/api_template.yaml"
output_swagger_file="../app/server/api/api.yaml"
envsubst < $input_swagger_file > $output_swagger_file

# INTEGRATION
# // app and mysql
export DATABASE_NAME=authdb
export DATABASE_TABLENAME=users
export DATABASE_USERNAME=auth_admin
export DATABASE_PASSWORD=admin
# // app and redis
export REDIS_PASSWORD=admin