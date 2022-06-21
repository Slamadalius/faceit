FROM golang:1.17.11-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/faceit

COPY go.* ./
RUN go mod download

COPY . .

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o ./faceit ./cmd/main.go" --command=./faceit