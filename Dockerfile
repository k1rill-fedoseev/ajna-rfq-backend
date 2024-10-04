FROM golang:1.22.5 AS build

WORKDIR /app

COPY . .

RUN mkdir out && go build -o ./out ./cmd/...

FROM ubuntu:22.04

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

COPY --from=build /app/out/ ./

EXPOSE 80

ENTRYPOINT ./server
