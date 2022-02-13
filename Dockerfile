FROM golang:alpine as builder

RUN mkdir ./build

ADD . /build/

WORKDIR /build

RUN go build -o main

From alpine

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app/

EXPOSE 8081

CMD ["./main"]