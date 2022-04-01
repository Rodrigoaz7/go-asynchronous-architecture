FROM golang:1.17

RUN mkdir src/mailService
WORKDIR /src/mailService
COPY . .

RUN go mod tidy && go build main.go

EXPOSE 8070

ENTRYPOINT ["./main"]