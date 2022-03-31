FROM postgres

ENV PORT=5432

ENV POSTGRES_PASSWORD=passw0rd
ENV POSTGRES_USER=test

EXPOSE $PORT

COPY ./database.sql /docker-entrypoint-initdb.d/


#COMAND TO BUILD THE DOCKER IMAGE
#docker build -t apidb ./database/

#COMAND TO CREATE THE CONTAINER
#docker run -p5432:5432 -it --rm apidb
