FROM golang:1.25

COPY . /app
WORKDIR /app
RUN go mod download
WORKDIR /app/src/cmd/event_api

CMD [ "bash", "run.sh" ]
