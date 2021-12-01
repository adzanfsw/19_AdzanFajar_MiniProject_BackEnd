FROM golang:1.17-alpine-3.14

WORKDIR /justrun-app

COPY . .

RUN go mod download

RUN go build -o mainfile

EXPOSE 3000

CMD ["./mainfile"]