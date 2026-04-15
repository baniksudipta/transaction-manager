#!/bin/sh
set -e

docker build -t transaction-manager .
docker run -p 8080:8080 transaction-manager