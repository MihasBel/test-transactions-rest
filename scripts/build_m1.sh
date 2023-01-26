#!/bin/bash
set -ex
cd `dirname $0`
docker buildx build --platform linux/amd64 -t test-transaction-rest:0.0.1 --load -f ../Dockerfile-amd --target=app ../