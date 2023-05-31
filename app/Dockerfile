FROM golang:alpine3.18
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./
RUN go mod download

# Copy Gp server files
COPY ./config ./config
COPY ./models ./models
COPY ./server ./server
COPY ./utils ./utils
COPY ./main.go ./

# Build the server application
RUN go build -o auth_server .

EXPOSE 3456
CMD ["./auth_server"]