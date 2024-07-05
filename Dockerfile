FROM golang:1.16

WORKDIR /app

COPY . .

# Fetching dependencies
RUN go mod tidy

# Building the Go app
RUN go build -o main .

# Making port 8080 available to the world outside this container
EXPOSE 8080

# Running the binary program we just built
CMD ["./main"]
