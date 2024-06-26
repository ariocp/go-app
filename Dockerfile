FROM golang:latest

WORKDIR ./cmd/app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o go-app ./cmd/app/main.go

CMD ["./go-app"]