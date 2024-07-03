PROJECT_NAME=top-selection-test
GO_VERSION=1.22.1
IMAGE_TAG=local
IMAGE_NAME=${PROJECT_NAME}:${IMAGE_TAG}

.PHONY: build
build:
	@docker build \
		--progress=plain \
		-f ./build/docker/Dockerfile \
		-t ${PROJECT_NAME}:${IMAGE_TAG} \
		.

.PHONY: start
start: build
	@docker run \
		--rm \
		${IMAGE_NAME}

.PHONY: test
test:
	go test \
	--count=1 \
	./...

