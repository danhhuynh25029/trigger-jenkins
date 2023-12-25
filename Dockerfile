FROM golang:alpine

COPY . /app

WORKDIR /app

RUN go build -o myapp

EXPOSE 8080

CMD ["./myapp"]
