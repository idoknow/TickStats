FROM golang:alpine AS builder
 
WORKDIR /build
COPY . .
RUN go build -o ./app

FROM alpine
 
WORKDIR /build
COPY --from=builder /build/app /build/app
 
CMD ["./app"]