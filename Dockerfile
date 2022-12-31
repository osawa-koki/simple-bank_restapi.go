FROM golang:1.19-bullseye

EXPOSE 80
WORKDIR /app

COPY . .
RUN go build -a -x -o main ./main.go

CMD ["./main"]
