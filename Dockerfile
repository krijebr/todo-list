FROM golang:1.24
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY /cmd ./cmd
COPY /config ./config
COPY /internal ./internal
RUN go build -v -o /usr/local/bin/app ./cmd
CMD ["app"]
EXPOSE 8080