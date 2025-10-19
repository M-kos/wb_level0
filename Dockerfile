FROM golang:1.24-alpine AS build
WORKDIR /go/app
ADD . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/orderRepository ./cmd/main.go

FROM scratch
COPY --from=build /go/app/bin/order /go/app/orderRepository

CMD ["/go/app/orderRepository"]
EXPOSE 8080
