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
RUN source ./.env

# Build the application
RUN go build -o auth_server .

EXPOSE 3456
CMD ["./auth_server"]