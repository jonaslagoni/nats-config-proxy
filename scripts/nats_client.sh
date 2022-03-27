#!/bin/bash -e
# Run from directory above via ./scripts/async_readme.sh
ag ../definitions/config-proxy.asyncapi.json @lagoni/go-nats-template --force-write -o ./output 