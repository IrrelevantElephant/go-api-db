# Description

Simple Go API backed by Postgres DB

To run build and run the code with db migrations:

`docker compose -f docker-compose.yml -f docker-compose.infra.yml up --build`

To run the int tests:

`docker compose -f docker-compose.yml -f docker-compose.infra.yml -f docker-compose.inttest.yml up --build --exit-code-from integrationtests`

or from the root directory:

`./run-inttests.sh`