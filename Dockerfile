FROM golang:1.24
RUN apt update
RUN apt install ffmpeg -y

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./ .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/api/main.go

EXPOSE 8080

# Run
CMD ["/main"]