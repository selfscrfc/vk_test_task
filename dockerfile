FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o app cmd/api/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/app /build/app

CMD ["/build/app"]