FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN make ws

RUN ls -l /app/tmp

CMD ["bash", "/app/tmp/ws"]