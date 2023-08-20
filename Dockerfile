FROM golang:1.10.3-alpine

WORKDIR /go/src/github.com/apanagiotou/go-kafka-to-s3/

COPY . .

RUN apk add librdkafka

FROM scratch
COPY --from=0 /go/src/github.com/apanagiotou/go-kafka-to-s3 .
ENTRYPOINT ["./go-kafka-to-s3"]

CMD ["tail", "-f", "/dev/null"]
