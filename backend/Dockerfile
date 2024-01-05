FROM golang:1.21-alpine
WORKDIR /app/backend
EXPOSE 5432

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o bin/gopherbank cmd/api/main.go

CMD ["./bin/gopherbank"]
