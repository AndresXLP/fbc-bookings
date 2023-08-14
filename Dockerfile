FROM golang:1.20-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
ARG EnvironmentVariable
RUN go mod download && go build -o main ./cmd
CMD /app/main

FROM golang

ADD ./internal/infra/resources/postgres/migrations  ./migrations

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

CMD sh -c './migrate -path ./migrations -database postgresql://${POSTGRESQL_DB_USER}:${POSTGRESQL_DB_PASSWORD}@${POSTGRESQL_DB_HOST}:${POSTGRESQL_DB_PORT}/${POSTGRESQL_DB_NAME}?sslmode=disable ${MIGRATIONS_COMMAND} ${MIGRATIONS_VALUE}'