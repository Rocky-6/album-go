FROM golang:1.20-bullseye as builder

WORKDIR /usr/src/api

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/api .


FROM debian:bullseye-slim

COPY --from=builder /usr/local/bin/api /api

CMD [ "/api" ]