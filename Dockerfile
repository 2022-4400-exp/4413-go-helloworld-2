FROM golang:alpine as build

WORKDIR /build

ARG TARGETOS
ARG TARGETARCH

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o helloworld

FROM gcr.io/distroless/static-debian11

COPY --from=build /build/helloworld /

CMD [ "/helloworld" ]
