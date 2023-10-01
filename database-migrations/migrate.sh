#!/bin/sh

psql -h $PGHOST -d $PGDATABASE -U $POSTGRES_USER -a -f ./migration.sql
psql -h $PGHOST -d $PGDATABASE -U $POSTGRES_USER -a -f ./insert-data.sql

# indicate migrations have completed
touch FINISHED

# Keeps this container alive
tail -f /dev/null