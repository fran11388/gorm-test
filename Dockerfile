FROM mysql:8

COPY ./init_db.sql /docker-entrypoint-initdb.d
#ADD default-configs.tar.gz /usr/src/app/

#WORKDIR /usr/src/app

#RUN npm install && \
#    mkdir -p /usr/logs && \
#    chown -R "1001:0" /usr/logs && \
#    chmod -R u+w /usr/logs

ENV  MYSQL_ROOT_PASSWORD password
#USER 1001

EXPOSE 3308:3306

#CMD ["node","hello-http.js"]