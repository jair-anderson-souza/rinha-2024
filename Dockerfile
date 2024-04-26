FROM golang:1.22
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

FROM alpine:3.14.10
EXPOSE 8080

# Copy files from builder stage
COPY --from=builder /app/bin/rinha .

# Run binary
CMD ["/rinha"]
