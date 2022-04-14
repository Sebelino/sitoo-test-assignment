FROM golang:1.18 as base

WORKDIR /app

COPY . .
RUN go build -o /sitoo-test-assignment

CMD ["/sitoo-test-assignment"]
