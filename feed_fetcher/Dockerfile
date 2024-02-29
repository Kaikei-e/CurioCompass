FROM golang:1.22.0

ADD ./feedFetcher /usr/src/app
WORKDIR /usr/src/app

EXPOSE 9000

CMD [ "go" , "run" , "main.go"]