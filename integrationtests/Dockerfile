FROM golang:1.21-alpine as build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go test -c -o integrationtests

FROM scratch

COPY --from=build /app/integrationtests /integrationtests

CMD ["./integrationtests", "-test.v"]