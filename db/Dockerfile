FROM mysql:8.0.27
ARG sec
ENV MYSQL_ROOT_PASSWORD $sec
COPY docker-entrypoint-initdb.d/* /docker-entrypoint-initdb.d/
COPY .my.cnf /etc/mysql/conf.d
COPY .msql /
RUN chmod 775 /etc/mysql/conf.d/.my.cnf \
    && mkdir -p /var/log/mysql \
    && chown -R mysql:mysql /var/log/mysql \
    && chmod 1777 /var/log/mysql;