FROM golang:1.15.6-alpine AS builder
WORKDIR /go/src/scoober-sync
COPY . .
RUN go get -d -v ./
RUN go install -v ./

FROM alpine
COPY --from=builder /go/bin/scoober-sync /bin/scoober-sync
CMD ["scoober-sync"]
