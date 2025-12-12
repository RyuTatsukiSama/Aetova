#!/bin/bash

image_name="image-aetova-serv"
container_name="aetova-serv"
container_port="51418"
machine_port="51418"

# Get the .env port
# while IFS='=' read -r key value; do
#     if [ "$key" = "PORT" ]; then
#         machine_port="$value"
#         container_port="$value"
#     fi
# done < .env

# Stop the container
docker stop "$container_name"
docker container rm "$container_name"

# Rebuild image
docker image rm "$image_name"
docker build --build-arg PORT="$container_port" -t "$image_name" .

# Relaunch the container
docker run -d --name "$container_name" --env-file .env -p "$machine_port:$container_port" "$image_name"

read