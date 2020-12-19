FROM golang:latest

ENV GO111MODULE=on
ENV GOPROXY_PORT=5000

WORKDIR /app
ADD . .

RUN go mod download
RUN go build -o proxy .

EXPOSE 5000

CMD ["/app/proxy"]