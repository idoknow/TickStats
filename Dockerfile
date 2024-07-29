FROM golang:alpine AS builder
 
WORKDIR /build
COPY . .
# 开启 CGO_ENABLED=0 
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine
 
WORKDIR /build
COPY --from=builder /build/app /build/app
 
CMD ["./app"]