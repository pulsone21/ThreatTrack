FROM golang:1.21.3

WORKDIR /app

COPY . .

RUN go get
RUN make build

ENTRYPOINT [ "/app/bin/frontEnd" ]