#!/bin/bash

ALBUMID=${1:?"missing arg 1 for ALBUMID"}

curl http://localhost:8080/albums/$ALBUMID | jq -C
