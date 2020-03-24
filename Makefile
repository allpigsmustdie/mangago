.PHONY: dev

all: docker

test:
	go test ./...

dev-build:
	go build -gcflags='-N -l' ./app/cmd/mangago/

air-hook: test dev-build

base-image:
	docker build . -f docker/base/Dockerfile -t mangago-base

tools-image: base-image
	docker build . -f docker/tools/Dockerfile -t mangago-tools

dev-image: tools-image
	docker build . -f docker/dev/Dockerfile -t mangago-dev

dev:
	docker run -v "$(shell pwd)/app:/home/mangago/app/" -p 8080:8080 -it --rm mangago-dev

generate-image: tools-image
	docker build . -f docker/generate/Dockerfile -t mangago-generate

generate:
	docker run -v "$(shell pwd)/app:/home/mangago/app/" -it --rm mangago-generate

docker: dev-image generate-image