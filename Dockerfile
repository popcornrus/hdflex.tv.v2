FROM golang:latest

# Update apt packages
RUN apt update

# Install libvips
RUN apt install -y libvips-dev

# Install air for hot reload
RUN go install github.com/cosmtrek/air@latest

WORKDIR /srv/hdflex.tv

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Start air with custom config
CMD ["air", "-c", ".air.toml"]