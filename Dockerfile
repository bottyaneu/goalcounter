# Use the corresponding go version
FROM golang:1.22.7-alpine

# Set the working directory
WORKDIR /app

# Copy the go mod and sum files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Run the app in production mode
ENV MODE="production"
# Build and run the app
RUN go build -o bin/goal
CMD ["./bin/goal", "-"]
