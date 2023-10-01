FROM alpine

WORKDIR /databasemigrations

RUN apk add postgresql-client

COPY ./migration.sql .
COPY ./migrate.sh .

RUN chmod +x ./migrate.sh

ENTRYPOINT [ "./migrate.sh" ]
# PGPASSWORD=mypassword psql -h database -d postgres -U postgres -a -f ./migrate.sql
