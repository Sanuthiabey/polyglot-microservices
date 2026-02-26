FROM golang:1.21-alpine

WORKDIR /app
COPY . .

RUN go build -o itemservice

EXPOSE 8081
CMD ["./itemservice"]