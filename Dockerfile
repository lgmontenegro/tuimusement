FROM golang:1.17.3

WORKDIR /go/src/app
COPY . .

RUN go mod vendor
RUN go mod tidy
RUN go build .

CMD ["tuimusement"]
