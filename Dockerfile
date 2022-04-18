FROM golang:1.18.1-bullseye as builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -a -o golang-ms .

FROM debian:bullseye-slim
WORKDIR /app/
COPY --from=builder /app/golang-ms .
CMD /app/golang-ms
