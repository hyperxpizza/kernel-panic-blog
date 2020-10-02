FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .

EXPOSE 5432
EXPOSE 8888

CMD [ "/app/main" ]