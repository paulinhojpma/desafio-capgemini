FROM golang:latest
ENV GO111MODULE=on

ADD . /defafio-api/sequence-validator
WORKDIR /defafio-api/sequence-validator/

COPY go.mod .
COPY go.sum .

EXPOSE 8890
RUN go build

CMD ["./sequence-validator"]
