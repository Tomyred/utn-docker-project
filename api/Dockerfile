FROM golang:1.20.4

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

# RUN CGO_ENABLED=0 GOOS=linux go build -o /build

RUN go build -o /build


EXPOSE 8080

CMD ["/build"]