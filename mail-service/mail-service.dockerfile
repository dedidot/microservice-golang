FROM golang:1.18-alpine as builder

RUN mkdir /app
RUN mkdir /templates

COPY . /app
COPY templates /templates

WORKDIR /app

RUN CGO_ENABLED=0 go build -o mailerApp ./cmd/api
RUN chmod +x /app/mailerApp

#Build tiny docker image
FROM alpine:latest
RUN mkdir /app
RUN mkdir /templates

COPY templates /templates

COPY --from=builder /app/mailerApp /app
CMD [ "/app/mailerApp" ]