FROM golang:1.21-alpine3.18 as base
RUN apk update 
WORKDIR /src
COPY go.mod go.sum ./
COPY . . 
RUN go build -o rinha 

FROM alpine:3.18 as binary
COPY --from=base /src/rinha .
EXPOSE 8000
CMD ["./rinha"]