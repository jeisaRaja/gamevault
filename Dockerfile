FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@v1.52.3

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o cmd/gamevault/main cmd/gamevault/*.go

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
