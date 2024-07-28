FROM golang:1.22-alpine3.20 as builder

RUN apk add --no-cache make nodejs npm

WORKDIR /server

COPY go.mod /server/go.mod
COPY cmd /server/cmd
COPY ui /server/ui
COPY makefile /server/makefile

RUN make build

FROM scratch

COPY --from=builder /server/server /server

CMD ["./server"]
