FROM golang:1.19 as builder

ENV CGO_ENABLED=0

WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go build -o ./build/ main.go

#############

FROM alpine:latest

COPY --from=builder /app/build .

COPY --from=builder /app/input.yaml ./input.yaml
COPY --from=builder /app/profile.jpg ./profile.jpg

CMD ["./main"]
