FROM golang:1.21-alpine3.18 AS builder

ADD . /app
WORKDIR /app

RUN go mod tidy -v
RUN go build -o api ./main.go 

###############Application Image################
FROM scratch AS final
WORKDIR /app
COPY --from=builder /app/api .
CMD ["/app/api"]
