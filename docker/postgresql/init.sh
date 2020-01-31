#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    ALTER USER postgres PASSWORD 'password';
    GRANT ALL PRIVILEGES ON DATABASE gymshark TO gymshark;

EOSQL

psql gymshark -c 'create extension hstore;'