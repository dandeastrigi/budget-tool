 FROM golang:1.14.2

ENV APP_NAME myproject
ENV PORT 8000

COPY ./src /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get -u ./...
RUN go get ./
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go build -o ${APP_NAME}

CMD ./${APP_NAME}

EXPOSE ${PORT}