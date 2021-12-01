FROM golang:1.17-alpine-3.14

WORKDIR /justrun-app

COPY go.mod ./

COPY go.sum ./

RUN go.mod download

COPY . .

RUN go build -o mainfile

EXPOSE 3000

CMD ["mainfile"]