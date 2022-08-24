FROM golang:1.19
RUN mkdir /app
ADD ./server /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]