FROM alpine

WORKDIR /build
COPY ./bin/app .

CMD ["./app"]