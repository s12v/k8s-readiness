NAME := readiness
IMAGE := gobuild-$(shell head -c4 </dev/urandom | xxd -p)
CONTAINER := gobuild-$(shell head -c4 </dev/urandom | xxd -p)
VERSION := "latest"

.PHONY: all
all: build image

build:
	docker build -f Dockerfile.build -t ${IMAGE} --no-cache .
	docker run --rm ${IMAGE} go test -v
	docker run -e GOOS=linux -e GOARCH=amd64 --name ${CONTAINER} ${IMAGE} go build -v
	docker cp ${CONTAINER}:/go/src/app/app .
	docker rm ${CONTAINER}

image:
	docker build -f Dockerfile -t $(NAME):latest .
	docker tag $(NAME):latest $(NAME):$(VERSION)
