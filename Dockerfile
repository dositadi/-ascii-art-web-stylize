# Stage 1
FROM golang:alpine3.22 AS builder

WORKDIR /ascii-art-web-stylize

COPY go.mod ./

RUN go mod download

COPY . .

EXPOSE 8081

RUN CGO_ENABLED=0 GOOS=linux go build -o /ascii-art ./cmd/

CMD [ "/ascii-art" ]

