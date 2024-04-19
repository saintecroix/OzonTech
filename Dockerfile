FROM golang:1.21 AS APP
RUN mkdir /app
COPY /. /app
WORKDIR /app
RUN go mod download
RUN go build ./cmd/app

FROM ubuntu:22.04
COPY --from=APP /app/app /app
EXPOSE 4000
CMD ["/app"]