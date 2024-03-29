FROM golang:alpine as builder

RUN export GO111MODULE=on

LABEL maintainer="amoCRM dmiroshnikov"
RUN apk update && apk add --no-cache git
WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/unisender_integration

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/app/ .
COPY --from=builder /go/src/app/.env .

EXPOSE 8080

CMD ["./main"]