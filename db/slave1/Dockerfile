FROM mysql
ARG sec
ENV MYSQL_ROOT_PASSWORD $sec
COPY ./slave1/docker-entrypoint-initdb.d/* /docker-entrypoint-initdb.d/
COPY ./slave1/socialdb.sql /
COPY ./slave1/.my.cnf /etc/mysql/conf.d
COPY ./slave1/.msql /
RUN chmod 775 /etc/mysql/conf.d/.my.cnf \
    && mkdir /var/log/mysql \
    && chmod 777 /var/log/mysql