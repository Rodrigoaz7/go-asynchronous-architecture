FROM golang:1.17

RUN mkdir src/publisher
WORKDIR /src/publisher
COPY . .

RUN go mod tidy && go build main.go

EXPOSE 8090

ENTRYPOINT ["./main"]