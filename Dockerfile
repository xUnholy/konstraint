FROM golang:1.14-alpine as build

ARG TARGETPLATFORM

ENV GO111MODULE=on \
  CGO_ENABLED=0

RUN apk add --no-cache git make

WORKDIR /go/src/github.com/xunholy/konstraint/

COPY . .

RUN export GOOS=$(echo ${TARGETPLATFORM} | cut -d / -f1) && \
  export GOARCH=$(echo ${TARGETPLATFORM} | cut -d / -f2) && \
  GOARM=$(echo ${TARGETPLATFORM} | cut -d / -f3); export GOARM=${GOARM:1} && \
  go build .

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=build --chown=nonroot /go/src/github.com/xunholy/konstraint/konstraint /usr/local/bin/

USER nonroot:nonroot

ENTRYPOINT ["konstraint"]

CMD ["help"]
