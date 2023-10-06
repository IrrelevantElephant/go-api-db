FROM golang:1.21-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine

COPY --from=build /app/main /main

EXPOSE 8080

CMD ["./main"]