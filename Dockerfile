FROM golang:alpine AS builder
ENV GO111MODULE=on
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go clean --modcache
COPY . .
RUN go build -o main
# EXPOSE 9000
# CMD ["./main"]

# Start 2 (Reduce Size Without Golang Image)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main . 
COPY --from=builder /app/.env .     
EXPOSE 9000
CMD ["./main"]