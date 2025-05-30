FROM golang:1.20 AS builder

WORKDIR /app

COPY . /app/

COPY /cmd/.env /app/.env

RUN CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o /app/shrinkray ./cmd


FROM arm64v8/alpine 

WORKDIR /app

COPY --from=builder /app/shrinkray /app/shrinkray
COPY --from=builder /app/.env /app/.env

RUN chmod +x /app/shrinkray

# RUN file /app/shrinkray

CMD ["/app/shrinkray"]