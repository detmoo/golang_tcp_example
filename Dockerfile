# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
ARG APP_NAME
ARG USER_NAME=$APP_NAME
ARG CONTAINER_PORT

RUN useradd $USER_NAME

WORKDIR /home/$USER_NAME

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /$APP_NAME

RUN chown -R $USER_NAME:$USER_NAME ./
USER $USER_NAME

EXPOSE $CONTAINER_PORT

COPY boot.sh ./
RUN chmod +x boot.sh
ENTRYPOINT ["./boot.sh"]
