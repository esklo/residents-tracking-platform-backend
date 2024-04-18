FROM golang:1.21


ENV DEV_MODE=true


WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /exec cmd/api/main.go

EXPOSE 8080

CMD ["/exec"]
