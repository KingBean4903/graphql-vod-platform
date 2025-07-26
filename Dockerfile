FROM golang:alpine AS builder

WORKDIR /app 

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o server


FROM scratch 

COPY --from=builder /app/server /app/server

EXPOSE 8800

ENTRYPOINT ["/app/server"]

