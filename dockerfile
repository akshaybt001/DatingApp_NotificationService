FROM golang:1.21.5-bullseye AS build

RUN apt-get update

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/cmd

RUN go build -o notification-service

FROM busybox:latest

WORKDIR /notification-service/cmd

COPY --from=build /app/cmd/notification-service .

COPY --from=build /app/.env /notification-service

EXPOSE 8083

CMD ["./notification-service"]