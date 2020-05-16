FROM golang:1.13

ENV APP_NAME app
ENV PORT 8000

COPY ./ /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go build .

CMD go run server.go

EXPOSE ${PORT}