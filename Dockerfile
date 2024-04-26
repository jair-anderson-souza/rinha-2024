FROM golang:1.22

WORKDIR /

COPY . .

RUN go mod download

RUN go build -o rinha 

EXPOSE 8080

# Run binary
CMD ["/rinha"]
