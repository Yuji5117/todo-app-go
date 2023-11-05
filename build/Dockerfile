FROM golang:1.21.3

WORKDIR /app
# COPY go.mod go.sum ./
COPY go.mod ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]