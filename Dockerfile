## BUILDER
FROM golang:1.12-alpine as builder 
WORKDIR /go/src/github.com/sanjayyepuri/Akira

RUN apk add --update make git
COPY . .

ENV GO111MODULE=on


RUN make build 

## Akira Container 
FROM alpine:latest 
WORKDIR /root

RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/sanjayyepuri/Akira/bin .

# TODO: use kubernetes secretes to distribute keys
CMD ["./Akira", "-t", "<put-secret-here>"]
