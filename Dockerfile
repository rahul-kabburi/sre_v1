FROM golang:1.21
WORKDIR /app
COPY src/*.go ./
RUN go mod init go-sre-agent && go mod tidy && go build -o app
CMD ["./app"]

