FROM golang:1.21.5-alpine as builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /rest-api ./cmd/rest


FROM scratch

COPY --from=builder ./rest-api /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# for run in local
# COPY /config/env/.env /config/env/  

ENTRYPOINT ["/rest-api"]
