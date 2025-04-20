FROM golang:1.24-alpine@sha256:7772cb5322baa875edd74705556d08f0eeca7b9c4b5367754ce3f2f00041ccee

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN /go/bin/swag init

RUN go build -o main .

CMD ["/app/main"]
