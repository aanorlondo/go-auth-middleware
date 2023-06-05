# Dockerfile
FROM redis/redis-stack:latest

COPY redis.conf /usr/local/etc/redis/redis.conf
COPY init.sh /usr/local/bin/init.sh

WORKDIR /usr/local/bin
RUN chmod +x init.sh

EXPOSE 6379
CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]
