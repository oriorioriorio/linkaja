FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=rootpassword
ENV MYSQL_DATABASE=linkaja
ENV MYSQL_USER=goappuser
ENV MYSQL_PASSWORD=goapppassword

COPY ./migrations /docker-entrypoint-initdb.d
