FROM golang:1.19

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /menuservice ./cmd/main.go

EXPOSE 50052

CMD [ "/menuservice" ]