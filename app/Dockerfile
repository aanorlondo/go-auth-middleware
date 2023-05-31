FROM golang:alpine3.18
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./
RUN go mod download

# Copy server files
COPY ./app/config ./config
COPY ./app/models ./models
COPY ./app/server ./server
COPY ./app/utils ./utils
COPY ./main.go ./
COPY ./.env ./

# Build the application
RUN go build -o server .

EXPOSE 3456
CMD ["./server"]