FROM golang:alpine AS build

RUN apk --no-cache add ca-certificates

WORKDIR /src/
ADD . /src/

RUN CGO_ENABLED=0 go build -o /bin/eth-mempool-whale-watcher

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/eth-mempool-whale-watcher /bin/eth-mempool-whale-watcher

ENTRYPOINT ["/bin/eth-mempool-whale-watcher"]
