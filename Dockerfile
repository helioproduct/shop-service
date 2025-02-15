FROM golang:1.23

WORKDIR ${GOPATH}/shop-service/
COPY . ${GOPATH}/shop-service/

RUN go build -o /bin/app ./cmd/main.go \
    && go clean -cache -modcache

EXPOSE 8080

CMD ["/bin/app"]