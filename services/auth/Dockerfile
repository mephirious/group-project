# Use a minimal base image for Go binaries
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the pre-built binary from the local `bin/` directory to the container
COPY bin/auth /app/auth

# Set the port that the service listens on
EXPOSE 5001

# Set permissions to make the binary executable
RUN chmod +x /app/auth

# Set the command to run the service
CMD ["./auth"]
