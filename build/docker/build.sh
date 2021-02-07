#!/bin/bash

set -eE  # same as: `set -o errexit -o errtrace`

# cd to repo root
cd "$(dirname "$0")"
cd ../../

IMAGE="osiris-pwm"
TAG="build-$(date +'%s')"
IMAGE_TAG="${IMAGE}:${TAG}"
CONTAINER_NAME="${IMAGE}-${TAG}"

clean () {
	docker rm -f $CONTAINER_NAME && echo "cleaned up container ${CONTAINER_NAME}"
	docker rmi -f $IMAGE_TAG && echo "cleaned up image ${IMAGE_TAG}"
}
trap clean ERR

build () {
	docker build \
		--no-cache \
		-t "${IMAGE_TAG}" \
       		-f build/docker/Dockerfile .

	docker create -ti --name $CONTAINER_NAME $IMAGE_TAG bash
	docker cp $CONTAINER_NAME:/app/build/artifacts/. build/artifacts/
}

build
clean
