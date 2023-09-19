FROM golang:latest
WORKDIR /app
COPY . .
RUN go build main.go
CMD ["go","run,","."] 