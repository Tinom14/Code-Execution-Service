FROM golang:1.23.6-alpine AS build

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk add --no-cache make

RUN go build -o processor-bin ./processor/main.go

FROM alpine AS runner

WORKDIR /app

RUN apk add --no-cache docker clang gcc python3 curl

COPY --from=build /build/processor-bin ./processor
COPY ./processor/config/config.yml ./config/config.yml

CMD ["/app/processor", "--config=./config/config.yml"]