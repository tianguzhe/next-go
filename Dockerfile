FROM golang:alpine3.16 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /nextgo

FROM alpine 

WORKDIR /

COPY --from=builder /nextgo /nextgo

ENTRYPOINT ["/nextgo"]