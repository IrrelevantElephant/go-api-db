FROM alpine

WORKDIR /databasemigrations

# need bash for healthcheck - try to improve this
RUN apk update && apk add bash && apk add postgresql-client

COPY ./migration.sql .
COPY ./insert-data.sql .
COPY ./migrate.sh .

RUN chmod +x ./migrate.sh

ENTRYPOINT [ "./migrate.sh" ]
