include Makefile.env

.PHONY: build
build:
	@docker build \
		--progress=plain \
		${BUILD_ARGS} \
		$(if $(filter true,$(DEBUG)), \
			--build-arg APK_EXTRA_PACKAGES="${APK_EXTRA_PACKAGES} delve" \
			--build-arg DEBUG_BUILD_FLAGS="${DEBUG_BUILD_FLAGS}" \
			--build-arg DEBUG_CMD="${DEBUG_CMD}" \
		) \
		-f ./build/docker/Dockerfile \
		-t ${PROJECT_NAME}:${IMAGE_TAG} \
		.

.PHONY: start
start:
	@mkdir -p ./.out
	@docker run \
		-it \
		--rm \
		--name ${CONTAINER_NAME} \
		$(if $(filter true,$(DETACH)),--detach) \
		$(if $(filter true,$(DEBUG)), \
			-p ${DEBUG_PORT}:${DEBUG_PORT} \
			--security-opt="apparmor=unconfined" \
			--cap-add=SYS_PTRACE \
			-v .:/go/src \
		) \
		-p 8080:80 \
		${IMAGE_NAME}

.PHONY: stop
stop:
	@docker stop \
		${CONTAINER_NAME} \
		1> /dev/null

.PHONY: test
test-unit:
	go test \
		-race \
		--tags=unit \
		./...

### TODO Need improve this command
### Stop after unsuccessfull run
.PHONY: test
test-integration: DETACH=true
test-integration:
	@APP_HOST=$(APP_HOST) APP_PORT=$(APP_PORT) go test \
		--count=1 \
		--tags=integration \
		./...
