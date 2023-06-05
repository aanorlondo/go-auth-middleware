FROM mysql:8.0.33

ENV MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD
COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 3306