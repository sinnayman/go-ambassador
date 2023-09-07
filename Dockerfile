FROM golang:1.20

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

workdir /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go install github.com/cosmtrek/air@latest

CMD ["sh", "-c", "sleep 5 && dockerize -wait tcp://db:3306 -timeout 60s air"]