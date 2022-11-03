# build
FROM golang:1.19 AS build

WORKDIR /app

COPY .  .

RUN CGO_ENABLED=0 GOOS=linux go build -o crudtmpl ./cmd

# deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /app .

ENV PORT 8080

EXPOSE $PORT

CMD ["./crudtmpl"]