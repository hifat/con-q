FROM golang:1.21.5-alpine as builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /con-q

FROM scratch

COPY --from=builder ./con-q /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV TZ=Asia/Bangkok

ENTRYPOINT ["/con-q"]
