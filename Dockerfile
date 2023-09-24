FROM golang:1.21
# Create working directory under /usr/src/app
WORKDIR /usr/src/app
# Copy over all go config (go.mod, go.sum etc.)
COPY go.* ./
# Install any required modules
RUN go mod download && go mod verify

# Copy over Go source code
# COPY . .

# Copy over Go source code
COPY *.go ./

# Run the Go build and output binary under /DevOpsDemo
RUN go build -o /DevOpsDemo
# Make sure to expose the port the HTTP server is using
EXPOSE 8080
# Run the app binary when we run the container
ENTRYPOINT ["/DevOpsDemo"]
