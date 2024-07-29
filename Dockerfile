FROM alpine

WORKDIR /build
COPY ./bin/app .

RUN chmod +x ./app

CMD ["./app"]