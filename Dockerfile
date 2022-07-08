FROM golang:alpine as gobuilder
COPY . /app
WORKDIR /app
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true
RUN mkdir app && \
    mkdir app/resources && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -extldflags "-static"' -o /app/approval
EXPOSE 1323
CMD ["./approval"]