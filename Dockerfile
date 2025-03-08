# Use the official Golang image
FROM golang:1.23-alpine as builder

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download


COPY . .

# Build the application
RUN go build -o pdf-generator ./cmd/server/main.go

# Use a minimal Alpine image for the final stage
FROM chromedp/headless-shell:stable

LABEL maintainer="eliasmflilico@gmail.com"

# Copy the binary from the builder stage
COPY --from=builder /app/pdf-generator /bin/pdf-generator

# Set the working directory
WORKDIR /app

ENV PORT 80
# Run the application
# Expose the application port
ENV PORT 80
EXPOSE $PORT

# Run Chrome in the background and start the Go application
ENTRYPOINT ["sh", "-c", " pdf-generator & /headless-shell/headless-shell --no-sandbox --use-gl=angle --use-angle=swiftshader --remote-debugging-address=0.0.0.0 --remote-debugging-port=9223"]