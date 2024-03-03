FROM golang:alpine

RUN mkdir /app

WORKDIR /app

COPY . .

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

RUN chmod +x /go/src/github.com/githubnemo/CompileDaemon

EXPOSE 8000

CMD [ "CompileDaemon", "-command=./main" ]

