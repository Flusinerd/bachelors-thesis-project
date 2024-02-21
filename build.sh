# #!/usr/bin/env bash

# # Exit on error
# set -e

# # Build the api
# docker build -t bachelor-thesis-api --network=host --no-cache  ./api

# # Make sure the api is running if not start it
# docker ps | grep -q bachelor-thesis-api || (docker stop bachelor-thesis-api && docker rm bachelor-thesis-api && docker run -d -p 80:80 -e PORT=80 --name bachelor-thesis-api bachelor-thesis-api)

# # Wait for the api to start
# sleep 5

# Build isr
docker build -t bachelor-thesis-isr --network=host  ./isr

# Builds ssg
docker build -t bachelor-thesis-ssg --network=host  ./ssg

# Build csr
docker build -t bachelor-thesis-csr --network=host  ./csr

# Build ssr
docker build -t bachelor-thesis-ssr --network=host  ./ssr