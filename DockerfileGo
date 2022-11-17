FROM golang:1.19.0-buster

WORKDIR /home/
COPY ./go /home/

RUN go mod init args
RUN go build

ENTRYPOINT ["./args"]