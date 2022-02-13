#!/bin/bash -e
# Run from directory above via ./scripts/async_readme.sh
ag ./asyncapi.json @asyncapi/markdown-template -o ./ -outFilename ./NATS-API.md --force