FROM golang:alpine as builder

WORKDIR /go/src/app

# Get Reflex for live reload in dev
ENV GO111MODULE=on

# Copy the necessary Go module files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Install Reflex
RUN go install github.com/cespare/reflex@latest

COPY . .

# Build your application
RUN go build -o ./run .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080
CMD [ "./run" ]
