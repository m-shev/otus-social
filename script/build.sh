#!/usr/bin/env bash
set -e
bash -c "cd .. && docker build -f nginx/Dockerfile --no-cache -t sigma-nginx:latest ."
docker save -o ./images/sigma-nginx.tar sigma-nginx:latest

## back
#bash -c "cd ../back && docker build -t sigma-social:latest ."
#docker save -o ./images/sigma-social.tar sigma-social:latest

## backend
#bash -c "cd ../db && bash ./build.sh"
#docker save -o ./images/sigma-social-db.tar sigma-social-db:latest