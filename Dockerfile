FROM golang:alpine

WORKDIR /src
COPY . /src

EXPOSE 8000

RUN cd /src
RUN go mod tidy
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/greeting.elf ./server

RUN apk --no-cache add curl