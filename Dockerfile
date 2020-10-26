FROM golang:1.15-buster

ADD ./build/hello app/hello

CMD "app/hello"