#!/usr/bin/env bash
set -e

server=sigma@77.244.214.34
imagesPath=/home/sigma/images

scp -P 5544 ./images/sigma-nginx.tar $server:$imagesPath
#scp -P 5544 ./images/sigma-social.tar $server:$imagesPath
#scp -P 5544 ./images/sigma-social-db.tar $server:$imagesPath
scp -P 5544 ./docker-compose.yml $server:/home/sigma

ssh -p 5544 $server docker load -i $imagesPath/sigma-nginx.tar
ssh -p 5544 $server docker load -i $imagesPath/sigma-social.tar
ssh -p 5544 $server docker load -i $imagesPath/sigma-social-db.tar


ssh -p 5544 $server docker-compose up