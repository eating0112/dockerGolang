FROM golang:latest

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o sayHello_master .

EXPOSE 8001
CMD ["/app/sayHello_master"]