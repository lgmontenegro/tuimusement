FROM golang:1.17.3

WORKDIR /go/src/tuimusement
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["tuimusement"]
