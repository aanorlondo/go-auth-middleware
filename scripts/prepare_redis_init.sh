#!/bin/bash

source prepare_env.sh

# redis.conf file
input_conf_file_path="../db/redis/redis_template.conf"
output_conf_file_path="../db/redis/redis.conf"
envsubst < $input_conf_file_path > $output_conf_file_path