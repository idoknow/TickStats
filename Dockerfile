FROM golang:alpine AS builder
 
WORKDIR /build
COPY . .
# 开启 CGO_ENABLED=0 
RUN apk add --no-cache --update go gcc g++
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine
WORKDIR /build
COPY --from=builder /build/app /build/app
COPY ./frontend/tick-stats-fe/dist /build/frontend/tick-stats-fe/dist
CMD ["./app"]