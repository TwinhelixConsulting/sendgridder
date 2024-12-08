FROM golang:1.23.4 as builder

WORKDIR /go/src/github.com/tomoconnor/mailgunner
COPY . /go/src/github.com/tomoconnor/mailgunner
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /go/src/github.com/tomoconnor/mailgunner/mailgunner

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/tomoconnor/mailgunner/mailgunner /app/mailgunner
ENV DOMAIN=""
ENV API_KEY=""
ENV SENDER=""
ENV SUBJECT=""
ENV RECIPIENT=""
ENTRYPOINT ["/app/mailgunner"]