FROM golang:1.21

WORKDIR /usr/src/app

COPY app/go.mod app/go.sum ./
RUN go mod download && go mod verify

COPY app .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]