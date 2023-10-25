# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Filipeb99"

# Set current working directory inside the container
WORKDIR /app

# Copy over all go config (go.mod, go.sum etc.)
COPY go.* ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download && go mod verify

# Copy the source from current directory to working directory inside the container
COPY . .

# Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN go build -o main

# Start new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose the port the HTTP server is using
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
