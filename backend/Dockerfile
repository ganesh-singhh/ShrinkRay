FROM golang:1.20

WORKDIR /app

COPY . /app/

COPY /cmd/.env /app/.env

RUN go build -o /app/shrinkray ./cmd

RUN chmod +x /app/shrinkray

RUN ls -l /app

EXPOSE 8080

CMD ["/app/shrinkray"]