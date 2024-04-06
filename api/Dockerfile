FROM golang:1.20.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /utn-project

EXPOSE 8080

CMD ["/utn-project"]