FROM golang:latest
ENV GO111MODULE=on

ADD . /defafio-cap/sequence-validator
WORKDIR /defafio-cap/sequence-validator/

COPY go.mod .
COPY go.sum .

EXPOSE 8890
RUN go build

CMD ["./sequence-validator"]
