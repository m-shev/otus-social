FROM mysql
ARG sec
ENV MYSQL_ROOT_PASSWORD $sec
COPY docker-entrypoint-initdb.d/* /docker-entrypoint-initdb.d/
COPY .my.cnf /etc/mysql/conf.d
COPY .msql /
RUN chmod 0600 /etc/mysql/conf.d/.my.cnf