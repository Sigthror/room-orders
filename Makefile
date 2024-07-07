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
	@docker run \
		-it \
		--rm \
		$(if $(filter true,$(DEBUG)), \
			-p ${DEBUG_PORT}:${DEBUG_PORT} \
			--security-opt="apparmor=unconfined" \
			--cap-add=SYS_PTRACE \
			-v .:/go/src \
		) \
		-p 8080:80 \
		${IMAGE_NAME}

.PHONY: test
test:
	go test \
	--count=1 \
	./...

.PHONY: debug
debug: DEBUG=true
debug:

.PHONY: help
help: ## describes all the targets in the makefile
	@echo "Help for repository:" && echo
	@echo "Options: (more options in Makefile.env)"
	@grep -h -E '^[a-zA-Z0-9_-]+[[:blank:]]*[:?!+]?=.*?##' $(MAKEFILE_LIST) | \
		sed -E 's/([^[:blank:]]+).*=[[:blank:]]*([^#]*)##(.*)/\1#-\3. Default: \2/' | sort -k1,1 | column -t -s "#"
	@echo && echo "Commands:"
	@grep -h -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | sed -E 's/(.*):.*##/\1#-/' | sort -k1,1 | column -t -s "#"

echo:
	echo ${DEBUG_CMD}
