FROM golang:1.20.5

WORKDIR /go-book-tracker

COPY . .

RUN go mod download
RUN go build -o books

EXPOSE 8080

CMD ["./books"]