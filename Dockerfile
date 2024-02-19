FROM golang:alpine

WORKDIR /my-forum

LABEL version="1.1"
LABEL maintainer="mamadbah, belhadjs, msoumare, edieng"

# Copy the application source code to the container
COPY . .

# Enable CGO and set the operating system
ENV CGO_ENABLED=1
ENV GOOS=linux

# Install bash, gcc, and other dependencies
RUN apk update && apk add --no-cache \
    bash \
    gcc \
    musl-dev \
    libc-dev \
    make

# Expose the port the application listens on
EXPOSE  4000

# Define the command to run the application
CMD go run ./cmd/web/.
