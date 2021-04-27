# Stage 1 (Build Golang Project)
FROM golang:alpine as builder
ENV GO111MODULE=on
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go clean --modcache
COPY . .
RUN go build -o main
EXPOSE 9000
CMD ["./main"]

# Start 2 (Reduce Size Without Golang Image)
# FROM alpine:latest
# WORKDIR /root/
# COPY --from=builder /app/main .     
# EXPOSE 3000
# CMD ["./main"]