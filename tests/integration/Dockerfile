FROM golang:1.23

WORKDIR /app

COPY . ./
# COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENTRYPOINT ["go", "test", "-v", "./tests/integration"]
