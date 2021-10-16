FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -exclude="*_test.go" --build="env GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${BIN_FILE_NAME} main.go"