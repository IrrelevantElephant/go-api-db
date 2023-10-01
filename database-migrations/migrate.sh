#!/bin/sh

psql -h $PGHOST -d $PGDATABASE -U $POSTGRES_USER -a -f ./migration.sql
psql -h $PGHOST -d $PGDATABASE -U $POSTGRES_USER -a -f ./insert-data.sql