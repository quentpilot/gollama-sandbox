# Dockerfile
FROM golang:1.23

RUN go install github.com/cespare/reflex@latest

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o sandbox .

EXPOSE 8080

#CMD ["./sandbox"]
CMD ["go", "run", "."]