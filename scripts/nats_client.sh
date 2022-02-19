#!/bin/bash -e
# Run from directory above via ./scripts/async_readme.sh
cd ../../go-nats-template && ag ../nats-config-proxy/nats-config-proxy/asyncapi.json ./ --force-write -o ../nats-config-proxy/nats-config-proxy/scripts/output 