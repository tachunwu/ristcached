
FROM golang:1.20.1-bullseye AS build


WORKDIR /ristcached
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /ristcached/ristcached ./cmd/server


FROM alpine:3.14
WORKDIR /ristcached
COPY --from=build /ristcached/ristcached .


EXPOSE 11212


CMD ["./ristcached"]


LABEL maintainer="tachunwu <tachunwu.go@gmail.com>"
