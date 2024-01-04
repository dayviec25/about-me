# Start from the official Golang base image
FROM golang:alpine as builder

# Set the working directory outside GOPATH to enable the support for modules.
WORKDIR /app

# Copy go.mod and go.sum and download the dependencies (for caching purposes)
COPY go.* ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Start a new stage from a smaller base image
FROM alpine:latest  

# Install Node.js and npm (needed for Tailwind CLI)
RUN apk add --no-cache nodejs npm

# Install Tailwind CLI globally
RUN npm install -g tailwindcss

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the templates directory
COPY --from=builder /app/templates ./templates

COPY --from=builder /app/dist ./dist
# Set the command to run the binary
CMD ["./main"]

