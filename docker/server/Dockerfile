# Use the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download



# Copy the source from the current directory to the Working Directory inside the container
COPY . . 

CMD ["air", "-c", ".air.toml"]

# Build the Go app


# CMD ["go", "run", "main.go"]





