FROM postgres
ENV POSTGRES_PASSWORD 12345
ENV POSTGRES_USER kopilka
ENV POSTGRES_DB kopilka
COPY ./kopilka.sql /docker-entrypoint-initdb.d/