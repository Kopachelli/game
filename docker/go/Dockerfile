FROM golang:1.13.7-alpine3.11

RUN apk --no-cache add htop git

RUN addgroup -S www-data && adduser -S www-data -G www-data

COPY --chown=www-data . /go/src/game

WORKDIR /go/src/game

RUN go build -o /game && chown www-data:www-data /game

USER www-data

CMD /game