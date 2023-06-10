FROM golang:1.20.5-alpine3.18 as builder
WORKDIR /app
COPY --from=developer /app .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.14.2 as production
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]

FROM golang:1.20.5-alpine3.18 as developer
RUN apk update && apk add --no-cache git
WORKDIR /app
RUN adduser -D -g '' appuser && \ 
chown -R appuser /app
USER appuser
COPY --chown=appuser:appuser . .
RUN go mod tidy && \
go install -v golang.org/x/tools/gopls@latest && \
go install -v github.com/ramya-rao-a/go-outline@latest
