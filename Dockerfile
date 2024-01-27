# Build stage
FROM golang:latest AS builder

WORKDIR /go/src/app

# Copy source files
COPY . .

# Get dependencies
RUN go get -d -v ./...

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM scratch

WORKDIR /app

# Copy the compiled executable from the builder stage
COPY --from=builder /go/src/app/main .

# Set environment variables if necessary
ENV PORT 8083

# Expose the port used by the application
EXPOSE $PORT

# Run the application upon container startup
CMD ["./main"]