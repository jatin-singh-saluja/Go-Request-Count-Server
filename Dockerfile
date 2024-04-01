FROM golang:latest
WORKDIR /app
# Add the source code:
ADD . .
RUN go mod download

ENTRYPOINT go run .
