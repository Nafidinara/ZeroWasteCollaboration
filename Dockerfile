FROM golang:1.22-alpine AS build

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux

RUN go build -o ./main cmd/main.go

FROM gcr.io/distroless/base-debian12 AS runtime

WORKDIR /

COPY --from=build ./app/main ./main
COPY --from=build ./app/.env .

EXPOSE 8080

ENTRYPOINT ["/main"]