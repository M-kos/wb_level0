FROM golang:1.24-alpine AS build
WORKDIR /go/app
ADD . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/order ./cmd/order-app/main.go

FROM scratch
COPY --from=build /go/app/bin/order /go/app/order

CMD ["/go/app/order"]
EXPOSE 8080
