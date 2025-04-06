FROM golang:1.23.6-alpine AS build

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk add --no-cache make

RUN go build -o server ./http_service/main.go

FROM alpine AS runner

WORKDIR /app

RUN apk add --no-cache curl netcat-openbsd


COPY --from=build /build/server ./server
COPY ./http_service/config/config.yml ./config/config.yml

CMD ["/app/server", "--config=./config/config.yml"]