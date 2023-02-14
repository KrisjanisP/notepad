FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 6969
CMD ["./main"]