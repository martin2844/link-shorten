# Use the official Golang image to create a build artifact.
FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o shortlink ./cmd/short

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/shortlink .

CMD ["./shortlink"]