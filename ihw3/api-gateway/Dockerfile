FROM golang:1.24-alpine

WORKDIR /app
COPY . .

RUN go build -o api-gateway .

CMD ["./api-gateway"]