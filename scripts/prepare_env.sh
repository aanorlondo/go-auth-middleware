#!/bin/bash

# for mysql init
export MYSQL_ROOT_PASSWORD=root

# for auth app
export DATABASE_HOSTNAME=$(hostname)
export DATABASE_PORT=3306
export APP_SECRET_KEY=azertyuiop12345

# for both
export DATABASE_NAME=authdb
export DATABASE_TABLENAME=users
export DATABASE_USERNAME=auth_admin
export DATABASE_PASSWORD=admin