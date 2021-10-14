FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="env GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin main.go"