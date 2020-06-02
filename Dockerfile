FROM golang:1.14

WORKDIR /go/src/github.com/alonelegion/account_storage_mongo

COPY . /go/src/github.com/alonelegion/account_storage_mongo

EXPOSE 8080

ENV TZ Europe/Moscow

CMD ["go", "run", "main.go"]