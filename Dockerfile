FROM golang:1.15.7-alpine AS builder
WORKDIR /go/src/scoober-sync
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine
COPY --from=builder /go/bin/scoober-sync /etc/periodic/hourly/sync
COPY --from=builder /go/bin/scoober-init /bin/scoober-init
ENV SCOOBER_CONFIG "/config"
CMD ["crond", "-f"]
