FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./Etc /app/Etc
RUN go build -ldflags="-s -w" -o /app/bookstore ./bookstore.go


FROM alpine

RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add curl

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app

COPY --from=builder /app/bookstore /app/bookstore
COPY --from=builder /app/Etc /app/Etc

CMD ["./bookstore"]
