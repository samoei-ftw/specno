#!/bin/sh

docker build -t specno-backend .

docker run -d -p 8080:8080 --name specno-backend specno-backend