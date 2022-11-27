# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
ARG APP_NAME
ARG ARTIFACT_PATH=dist
ARG USER_NAME=$APP_NAME
ARG CONTAINER_PORT

RUN adduser -D -g "$USER_NAME" $USER_NAME

WORKDIR /home/$USER_NAME

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY cmd/ ./cmd/
COPY pkg/ ./pkg/

RUN mkdir -p $ARTIFACT_PATH
RUN go build -o ./$ARTIFACT_PATH/$APP_NAME

RUN chown -R $USER_NAME: ./
USER $USER_NAME

EXPOSE $CONTAINER_PORT

ENTRYPOINT ./$ARTIFACT_PATH/$APP_NAME
