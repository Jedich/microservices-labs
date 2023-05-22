FROM golang:1.19-alpine3.16 AS server_builder

RUN apk add gcc libc-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags "-w -s -linkmode external -extldflags -static" -o /user-service

FROM scratch
EXPOSE 8083

COPY --from=server_builder /user-service .
COPY ./migrations ./migrations

CMD ["/user-service"]