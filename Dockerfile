FROM golang:1.12.0-alpine3.9
RUN mkdir /Jann-Pass
ADD . /Jann-Pass
WORKDIR /app
RUN go build -o main .
CMD ["/Jann-Pass/main"]