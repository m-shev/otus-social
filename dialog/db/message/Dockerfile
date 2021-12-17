FROM mysql:8.0.27
COPY dialog/docker-entrypoint-initdb.d/* /docker-entrypoint-initdb.d/
COPY dialog/.my.cnf /etc/mysql/conf.d
ENV MYSQL_ROOT_PASSWORD "secret"
RUN chmod 775 /etc/mysql/conf.d/.my.cnf \
    && mkdir -p /var/log/mysql \
    && chown -R mysql:mysql /var/log/mysql \
    && chmod 1777 /var/log/mysql;