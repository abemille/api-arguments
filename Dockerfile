FROM golang:alpine as builder

COPY . "/go/src/github.com/wikisophia/api-arguments"
WORKDIR "/go/src/github.com/wikisophia/api-arguments"
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .


FROM scratch
COPY --from=builder /go/src/github.com/wikisophia/api-arguments /app/
WORKDIR /app

EXPOSE 8001
CMD ["./main"]