# get go image
FROM golang:1.15-alpine

WORKDIR /go/src/app
COPY . .

# get go modules
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8888
CMD [ "server" ]