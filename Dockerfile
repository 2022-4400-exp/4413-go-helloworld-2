FROM --platform=${BUILDPLATFORM} golang:alpine as build

WORKDIR /build

ARG TARGETOS
ARG TARGETARCH

COPY main.go go.mod ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod tidy
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w -extldflags" -o helloworld


FROM scratch

COPY --from=build /build/helloworld /

CMD [ "/helloworld" ]
