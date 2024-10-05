FROM golang:1.18-alpine

WORKDIR /app/

RUN go mod init serversample

COPY . .

RUN go mod tidy
RUN go build -o serversample ./src/serversample/main.go

CMD [ "./serversample" ]