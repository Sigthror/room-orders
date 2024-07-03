ARG GO_VERSION=1.22.1
ARG ALPINE_VERSION=3.19
ARG BUILD_IMAGE=golang:${GO_VERSION}-alpine${ALPINE_VERSION}
ARG PROJECT_NAME=top-selection-test

FROM ${BUILD_IMAGE} as build

ARG PROJECT_NAME
ENV GOPATH=/go \
    GOSUMDB=off \
    GOCACHE=/.cache/.go-build \
    GOMODCACHE=/.cache/.go-pkg
WORKDIR /go/src

RUN --mount=type=bind,readwrite,target=/go/src \
    --mount=type=cache,sharing=shared,id=gomod,target=/.cache/.go-pkg \
    --mount=type=cache,sharing=shared,id=gocache,target=/.cache/.go-build \
    go mod download -x \
    && go build -v -o /usr/local/bin/${PROJECT_NAME} ./cmd


FROM alpine:${ALPINE_VERSION} as prod

ARG PROJECT_NAME
ENV PROJECT_NAME=${PROJECT_NAME}
COPY --from=build /usr/local/bin/${PROJECT_NAME} /usr/local/bin/${PROJECT_NAME}

RUN chmod +x /usr/local/bin/${PROJECT_NAME}

ENTRYPOINT exec ${PROJECT_NAME}
