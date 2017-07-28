#!/bin/bash
# Base of script is from https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
# Modified to test if there were no older images to purge and updated to use
# variables instead
USERNAME="DOCKER_HUB_USERNAME"
CONTAINER_NAME="CONTAINER NAME"

docker pull $USERNAME/$CONTAINER_NAME:latest
if docker stop $CONTAINER_NAME; then docker rm $CONTAINER_NAME; fi
docker run -d -p 80:8080 --name $CONTAINER_NAME  $USERNAME/$CONTAINER_NAME
if [ $(docker images --filter "dangling=true" -q --no-trunc | wc -l) -gt 0 ]; then docker rmi $(docker images --filter "dangling=true" -q --no-trunc); fi
