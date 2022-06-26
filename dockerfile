#build stage
FROM golang:1.18 AS builder

WORKDIR /app

COPY . .

RUN go build -tags netgo -o main.app

#final stage
FROM alpine:latest

WORKDIR /release

COPY --from=builder /app/main.app .

CMD [ "/release/main.app" ]