FROM docker.io/golang:1.22.0-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /proxima-api

CMD ["/proxima-api"]
