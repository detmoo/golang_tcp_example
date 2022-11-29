# syntax=docker/dockerfile:1

FROM golang:1.17-alpine
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
RUN chmod +x ./$ARTIFACT_PATH/$APP_NAME

COPY boot.sh ./
RUN chmod +x boot.sh

RUN chown -R $USER_NAME: ./
USER $USER_NAME

EXPOSE $CONTAINER_PORT
ENV EXECUTABLE=$ARTIFACT_PATH/$APP_NAME
ENTRYPOINT ["./boot.sh"]
CMD ["server", "--host", "localhost", "--port", "5000", "--duration", "60s"]
