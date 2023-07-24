FROM golang:1.20.5

WORKDIR /go-book-tracker

COPY . .

RUN go mod download
RUN go build -o book-tracker

EXPOSE 8080

CMD ["./book-tracker"]