FROM golang:1.15-alpine

WORKDIR /go/src/github.com/danielrcoura/go-wallet

COPY . .

RUN cd cmd && go build -o ../dist/go-wallet

CMD ["./dist/go-wallet"]

EXPOSE 3000