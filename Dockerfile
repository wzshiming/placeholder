FROM golang:alpine AS builder
WORKDIR /go/src/github.com/wzshiming/placeholder
COPY . .
RUN go install github.com/wzshiming/placeholder/cmd/...

FROM wzshiming/upx AS upx
COPY --from=builder /go/bin/ /go/bin/
RUN upx /go/bin/*

FROM alpine
COPY --from=upx /go/bin/ /usr/local/bin/
EXPOSE 8080
CMD placeholder
