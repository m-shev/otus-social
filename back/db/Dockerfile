FROM mysql
ENV MYSQL_ROOT_PASSWORD 1234
COPY docker-entrypoint-initdb.d/* /docker-entrypoint-initdb.d/
COPY .my.cnf /etc/mysql/conf.d
RUN chmod 0600 /etc/mysql/conf.d/.my.cnf