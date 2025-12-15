#!/bin/bash

# Build and run the Livebook container with HiGHS support
docker build -t advent-livebook .

docker run -p 8090:8090 -p 8091:8091 \
  -e LIVEBOOK_PORT=8090 \
  -e LIVEBOOK_PASSWORD="securesecret" \
  -e LIVEBOOK_IFRAME_PORT=8091 \
  -u $(id -u):$(id -g) \
  -v "$(pwd)":/data \
  advent-livebook
