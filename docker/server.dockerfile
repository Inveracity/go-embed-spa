FROM golang:1.22-alpine3.20 as builder

RUN apk add --no-cache make nodejs npm curl bash sudo

WORKDIR /server

COPY go.mod /server/go.mod
COPY go.sum /server/go.sum
COPY cmd /server/cmd
COPY internal /server/internal
COPY ui /server/ui
COPY makefile /server/makefile
COPY --chmod=755 scripts/install_upx.sh /server/scripts/install_upx.sh
RUN make build
RUN make install-upx
RUN make compress

FROM scratch

COPY --from=builder /server/bin/cli /cli

ENTRYPOINT ["./cli"]
