# PRODUCTION
#--------------------------------------
FROM golang:1.20.5-alpine3.18 as builder
LABEL maintainer="Miguel Gutierrez <gt2rz.dev@gmail.com>"
WORKDIR /app
COPY --chown=appuser:appuser . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM alpine:3.14.2 as production
LABEL maintainer="Miguel Gutierrez <gt2rz.dev@gmail.com>"
LABEL version = "1.0"
LABEL description = "API REST for authentication and authorization with JWT"
RUN apk update && apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 3000
CMD ["./main"]


# DEVELOPMENT
#--------------------------------------
FROM golang:1.20.5-alpine3.18 as developer
LABEL maintainer="Miguel Gutierrez <gt2rz.dev@gmail.com>"
RUN apk update && apk add --no-cache git
WORKDIR /app
RUN adduser -D -g '' appuser && \ 
  chown -R appuser /app
USER appuser
COPY --chown=appuser:appuser . .
RUN go mod tidy && \
  go install github.com/cosmtrek/air@latest && \
  go mod download

CMD ["air", "-c", ".air.toml"]


# DATABASES
#--------------------------------------
FROM postgres as postgres
COPY ./scripts/postgres/init.sql docker-entrypoint-initdb.d/init.sql
EXPOSE 5432
VOLUME /var/lib/postgresql/data