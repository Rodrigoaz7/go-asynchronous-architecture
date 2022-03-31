FROM golang:1.17

RUN mkdir src/consumer
WORKDIR /src/consumer
COPY . .

RUN go mod tidy && go build main.go

EXPOSE 8080

ENTRYPOINT ["./main"]