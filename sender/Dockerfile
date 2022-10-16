FROM golang:latest
ENV GO111MODULE=on

ADD . /sender
WORKDIR /sender/

COPY go.mod .
COPY go.sum .

EXPOSE 8890
RUN go build

CMD ["./sender"]
