FROM golang:1.17-buster

#download packages
COPY go.mod .
COPY go.sum .
ENV GOPATH=/
RUN go mod download

#build appliction
COPY . .
RUN go build -o finance-app ./cmd/main.go

CMD ["./finance-app"]