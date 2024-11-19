#!/bin/bash

# 🐳 Create Dockerfile
# goctl docker -go insou.go

# 🐳 Build Docker image
docker build --build-arg VERSION=$(cat VERSION) -t insou:$(cat VERSION) .


# 📜 List Docker images
docker image ls

# 🏃 Run Docker container
docker run --rm -it -p 6868:8091 loginservice:v1

# 💻 Test if Docker container is running
curl -i http://localhost:8868688/echo
