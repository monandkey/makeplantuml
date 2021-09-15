# Makefile
DOCKER_ENV                   ?= DOCKER_BUILDKIT=1
DOCKER_BUILD_ARGS            ?= --rm

GOLANG_ZIP_DOCKER_TAG        ?= 1.0
GOLANG_ZIP_DOCKER_REGISTRY   ?= docker.io
GOLANG_ZIP_DOCKER_REPOSITORY ?= monandkey/golang-zip

DOCKER_TAG                   ?= 1.0
DOCKER_REGISTRY              ?= docker.io
DOCKER_REPOSITORY            ?= monandkey/makeplantuml

GOLANG_ZIP                   ?= ${GOLANG_ZIP_DOCKER_REGISTRY}/${GOLANG_ZIP_DOCKER_REPOSITORY}:${GOLANG_ZIP_DOCKER_TAG}
MAKEPLANTUML                 ?= ${DOCKER_REGISTRY}/${DOCKER_REPOSITORY}:${DOCKER_TAG}

build-all: build-golang-zip build-makeplantuml

.PHONY: build-golang-zip
build-golang-zip: 
	${DOCKER_ENV} docker build ${DOCKER_BUILD_ARGS} \
		--tag ${GOLANG_ZIP} \
		--file ./zip/Dockerfile \
		./

.PHONY: build-makeplantuml
build-makeplantuml:
	${DOCKER_ENV} docker build ${DOCKER_BUILD_ARGS} \
		--tag ${MAKEPLANTUML} \
		--file ./Dockerfile \
		--no-cache \
		./
	docker run -itd --name makeplantuml ${DOCKER_REPOSITORY}:${DOCKER_TAG} /bin/bash
	docker container cp makeplantuml:/home/makeplantuml-windows.zip .
	docker container stop makeplantuml
	docker container rm makeplantuml
