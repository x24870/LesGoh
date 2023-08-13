FROM golang:1.18-alpine
LABEL maintainer="Jan <jan.jiazhen@gmail.com>"

RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN go build -o demo_app main.go

EXPOSE 8080
ENTRYPOINT ["./demo_app"]