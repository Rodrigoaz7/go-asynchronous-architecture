# docker build -t *nome da imagem* .
# docker run --name *nome conteiner* -p 8080:8080 *nome da imagem* 
FROM golang:1.17

RUN mkdir src/app
WORKDIR /src/app
COPY . .

RUN go mod tidy && go build main.go

EXPOSE 8090

ENTRYPOINT ["./main"]