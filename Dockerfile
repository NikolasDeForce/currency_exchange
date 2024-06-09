FROM golang:latest

COPY . /app

EXPOSE 8010
WORKDIR /app
ENTRYPOINT [ "go", "run", "server.go" ]