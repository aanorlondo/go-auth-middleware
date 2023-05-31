#!/bin/bash

source prepare_env.sh

# Read the setup_with_vars.sql file
input_sql_file_path="../db/mysql/init_template.sql"
output_sql_file_path="../db/mysql/init.sql"

# Replace the environment variable placeholders with their values
envsubst < $input_sql_file_path > $output_sql_file_path