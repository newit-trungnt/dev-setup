FROM golang:alpine as builder
ENV SERVICE_PORT=8000\
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify
COPY . .
RUN go build -ldflags "-s -w" -o main .

FROM scratch
WORKDIR /app
ADD responses ./responses
COPY --from=builder /build/main .
EXPOSE ${SERVICE_PORT}
CMD ["/app/main"]