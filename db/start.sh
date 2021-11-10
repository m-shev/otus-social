docker stop social-db
#docker run --name social-db --rm -p 3306:3306 -v social-db-data:/var/lib/mysql sigma-social-db:latest
docker run --name social-db --rm -p 3306:3306  sigma-social-db:latest