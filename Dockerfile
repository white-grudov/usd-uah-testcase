FROM golang:1.22.3-alpine as build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]