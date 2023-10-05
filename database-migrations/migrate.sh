#!/bin/sh

psql -h $PGHOST -d $PGDATABASE -U $POSTGRES_USER -a -f ./migration.sql
psql -h $PGHOST -d $PGDATABASE -U $POSTGRES_USER -a -f ./insert-data.sql

# indicate migrations have completed for the healthcheck
touch FINISHED

# Force alpine to respect SIGTERM so the container shuts down more quickly
trap "exit" SIGTERM

# Keeps this container alive - `tail -f /dev/null` is more standard but runs as a foreground process and blocks the above `trap` command.
while true; do
    sleep 1
done
