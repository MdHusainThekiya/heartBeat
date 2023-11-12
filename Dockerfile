# Use an official Go runtime as a parent image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Build the Go binary
RUN chmod +x build.sh && sh build.sh

# Expose a port for your Go service
EXPOSE 8080

# Run your Go service when the container starts
CMD ["./heartBeat_binary"]