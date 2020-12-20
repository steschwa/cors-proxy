FROM golang:alpine

ENV GO111MODULE=on
ENV CORSPROXY_PORT=5000

WORKDIR /app
ADD . .

RUN go mod download
RUN go build -o proxy .

EXPOSE 5000

CMD ["/app/proxy"]