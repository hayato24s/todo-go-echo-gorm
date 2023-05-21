# Development
FROM golang:1.20-bullseye as dev

WORKDIR /go/src

COPY go.mod go.sum ./

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz \
  && mv ./migrate /usr/local/bin/ \
  && go mod download \
  && go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

# Build
FROM golang:1.20-bullseye as build

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# Deploy
FROM debian:bullseye-slim as deploy

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz \
  && mv ./migrate /usr/local/bin/

COPY --from=build /go/src/app .

EXPOSE 8080

CMD ["./app"]