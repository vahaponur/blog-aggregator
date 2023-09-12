# Use an official Go runtime as a parent image
FROM golang:1.21-bullseye

# Set the working directory in the container
WORKDIR /usr/src/app
COPY . .

RUN go mod tidy




