FROM golang:1.21

workdir /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o main ./cmd/main.go

CMD ["./main"]