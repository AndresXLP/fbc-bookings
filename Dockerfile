FROM golang:1.20-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
ARG EnvironmentVariable
RUN go mod download && go build -o main ./cmd
CMD /app/main
