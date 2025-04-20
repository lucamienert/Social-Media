FROM golang:1.24-alpine@sha256:7772cb5322baa875edd74705556d08f0eeca7b9c4b5367754ce3f2f00041ccee

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download -x

COPY . .

RUN go build -o main .

CMD ["/app/main"]
