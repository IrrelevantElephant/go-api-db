#!/bin/bash

docker compose -f docker-compose.yml -f docker-compose.infra.yml -f docker-compose.inttest.yml up --build --exit-code-from integrationtests