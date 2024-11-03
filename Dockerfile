FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . /app/
RUN go get api
RUN CGO_ENABLED=0 go build . 
RUN chmod +x /app/api

FROM alpine:latest

RUN mkdir /out
COPY --from=builder /app/api /out/
CMD ["/out/api"]