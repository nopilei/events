FROM golang:1.25

COPY . /app
WORKDIR /app
RUN go mod download
WORKDIR /app/src/cmd/server

CMD [ "bash", "run.sh" ]
