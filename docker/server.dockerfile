FROM golang:1.22-alpine3.20 as builder

ENV GIN_MODE=release

RUN apk add --no-cache make nodejs npm

WORKDIR /server

COPY go.mod /server/go.mod
COPY go.sum /server/go.sum
COPY cmd /server/cmd
COPY internal /server/internal
COPY ui /server/ui
COPY makefile /server/makefile
RUN make build

FROM scratch

COPY --from=builder /server/bin/cli /cli

ENTRYPOINT ["./cli"]
