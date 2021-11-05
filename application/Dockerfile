FROM golang:1.15 as build

COPY . .

RUN go build -o server .

ENTRYPOINT [ "./server" ]