#!/bin/bash

# MYSQL
export MYSQL_ROOT_PASSWORD=root

# AUTH SERVER
export DATABASE_HOSTNAME=$(hostname)
export DATABASE_PORT=3306
export REDIS_HOSTNAME=$(hostname)
export REDIS_PORT=6379
export APP_SECRET_KEY=azertyuiop12345

# INTEGRATION
# // app and mysql
export DATABASE_NAME=authdb
export DATABASE_TABLENAME=users
export DATABASE_USERNAME=auth_admin
export DATABASE_PASSWORD=admin
# // app and redis
export REDIS_PASSWORD=admin