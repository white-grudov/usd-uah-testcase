# Start with the official Golang image
FROM golang:1.22.3-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Install git to allow go mod download of private repositories
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Install PostgreSQL client and CA certificates
RUN apk --no-cache add postgresql-client ca-certificates bash

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/main .
COPY --from=build /app/database/migration.sql .
COPY .env ./

# Copy the wait-for-it script
COPY wait-for-it.sh .

# Make the wait-for-it script executable
RUN chmod +x wait-for-it.sh

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the database migrations and then start the server
CMD ./wait-for-it.sh db:5432 -- psql $DATABASE_URL -f /app/migration.sql && ./main
