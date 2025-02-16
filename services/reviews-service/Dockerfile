# Use a minimal base image for Go binaries
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

RUN ls -la /app

# Copy the pre-built binary from the local `bin/` directory to the container
COPY ./bin/reviews /app/reviews

# Set the port that the service listens on
EXPOSE 5004

# Set permissions to make the binary executable
RUN chmod +x /app/reviews

# Set the command to run the service
CMD ["./reviews"]
