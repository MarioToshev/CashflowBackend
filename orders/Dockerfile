FROM golang:1.23.1

# Set destination for COPY
WORKDIR /order-service

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Ensure all dependencies are fetched
RUN go mod tidy
RUN go mod vendor
RUN go build ./...
# Build
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -mod=vendor -a -o service main.go

# Expose port
EXPOSE 8081

# Run
CMD ["/order-service/service"]


