FROM golang:1.19-alpine

ARG DB_USER=root
ARG DB_PASSWORD=password
ARG DB_NAME=bookkeeper

ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /bookkeeper-docker

EXPOSE 8080

CMD ["/bookkeeper-docker"]
