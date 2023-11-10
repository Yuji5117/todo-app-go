FROM golang:1.21.3

WORKDIR /app
COPY go.mod go.sum ./

RUN go install github.com/cosmtrek/air@latest
RUN go mod download

COPY . ./

EXPOSE 8080

# Run
CMD ["air", "-c", ".air.toml"]