FROM golang:alpine as builder

WORKDIR build

COPY . ./

RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=false go build -o ./bin

FROM scratch as final

WORKDIR app

COPY --builder=builder /build/bin ./
COPY --builder=builder /build/configs/local.yaml ./configs/

CMD ["./bin"]