FROM golang:1.22


WORKDIR /app

COPY . .

RUN go mod init main
RUN go mod tidy

RUN go build -o main ./cmd/server

CMD ["./main"]

# docker build -t my-go-server .
# docker run -itd -p 8080:8080 my-go-server