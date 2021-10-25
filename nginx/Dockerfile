FROM nginx:1.21.3
COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY front/build/ /opt/front/public/

RUN set -x \
    && apt-get update -y \
    && apt-get upgrade -y \
    && apt-get install tzdata -y \
    && cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime \
    && echo "Europe/Moscow" > /etc/timezone \
    && apt-get install certbot -y \
    && apt-get install python-certbot-nginx -y;