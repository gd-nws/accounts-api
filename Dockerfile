# Start from the latest golang base image
FROM golang:alpine as builder

RUN mkdir -p /app

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY ./src /app

ENV GOPATH=/app
ENV GOBIN=/app/bin

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go get

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]