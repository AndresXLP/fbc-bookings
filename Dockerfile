FROM golang

ADD ./internal/infra/resources/postgres/migrations  ./migrations

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

CMD sh -c './migrate -path ./migrations -database postgresql://${POSTGRES_DB_USER}:${POSTGRES_DB_PASSWORD}@${POSTGRES_DB_HOST}:${POSTGRES_DB_PORT}/${POSTGRES_DB_NAME}?sslmode=disable ${MIGRATIONS_COMMAND} ${MIGRATIONS_VALUE}'

FROM golang:1.20-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
ARG EnvironmentVariable
RUN go mod download && go build -o main ./cmd
CMD /app/main