FROM golang

WORKDIR /app/src

ENV GOPATH=/app

COPY . /app/src

RUN go build main.go

ENTRYPOINT ["./main"]

EXPOSE 8080