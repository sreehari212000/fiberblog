FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./ 
RUN go mod download
COPY . .
# output of go build will be sored in main.
RUN go build -o main . 

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 3000
CMD [ "./main" ]