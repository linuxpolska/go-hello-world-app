FROM golang:alpine as builder
WORKDIR /build/app

COPY server.go go.mod ./

# Build app
RUN go build -o myapp

FROM alpine:latest
COPY --from=builder /build/app/myapp ./myapp
EXPOSE 8080
CMD ["./myapp"]
