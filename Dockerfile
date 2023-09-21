FROM golang:1.21

WORKDIR /app

COPY /app/go.mod /app/go.sum ./
RUN go mod download && go mod verify

COPY /app .
RUN go build -o bin main.go

CMD ["/app/bin"]