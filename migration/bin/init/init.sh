#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER breadcrumbs WITH PASSWORD 'breadcrumbs';
    CREATE DATABASE breadcrumbs;
    GRANT ALL PRIVILEGES ON DATABASE breadcrumbs TO breadcrumbs;
EOSQL

export PGPASSWORD=breadcrumbs

psql -v ON_ERROR_STOP=1 --username "toreglia" --dbname "breadcrumbs"  <<-EOSQL
    CREATE EXTENSION IF NOT EXISTS postgis;
EOSQL