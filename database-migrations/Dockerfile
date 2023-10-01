FROM alpine

WORKDIR /databasemigrations

RUN apk add postgresql-client

COPY ./migration.sql .
COPY ./insert-data.sql .
COPY ./migrate.sh .

RUN chmod +x ./migrate.sh

ENTRYPOINT [ "./migrate.sh" ]
