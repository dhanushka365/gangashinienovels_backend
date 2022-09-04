FROM golang:latest 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o  main services/book/cmd/book/main.go 
EXPOSE 8081
ENTRYPOINT ["./main"]

