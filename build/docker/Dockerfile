ARG PROJECT_NAME
ARG GO_VERSION
ARG ALPINE_VERSION

ARG BUILD_IMAGE=golang:${GO_VERSION}-alpine${ALPINE_VERSION}

FROM ${BUILD_IMAGE} as build

ARG PROJECT_NAME
ARG DEBUG_BUILD_FLAGS

ENV GOPATH=/go \
    GOSUMDB=off \
    GOCACHE=/.cache/.go-build \
    GOMODCACHE=/.cache/.go-pkg
WORKDIR /go/src

RUN --mount=type=bind,readwrite,target=/go/src \
    --mount=type=cache,sharing=shared,id=gomod,target=/.cache/.go-pkg \
    --mount=type=cache,sharing=shared,id=gocache,target=/.cache/.go-build \
    go mod download -x \
    && go build -gcflags="${DEBUG_BUILD_FLAGS}" -v -o /usr/local/bin/${PROJECT_NAME} ./cmd


FROM alpine:${ALPINE_VERSION} as prod

ARG PROJECT_NAME
ARG DEBUG_CMD
ARG APK_EXTRA_PACKAGES
ENV PROJECT_NAME=${PROJECT_NAME}
ENV DEBUG_CMD=${DEBUG_CMD}

RUN apk add --update --no-cache ${APK_EXTRA_PACKAGES}

COPY --from=build /usr/local/bin/${PROJECT_NAME} /usr/local/bin/${PROJECT_NAME}

RUN chmod +x /usr/local/bin/${PROJECT_NAME}

ENTRYPOINT ${DEBUG_CMD} exec /usr/local/bin/${PROJECT_NAME}
