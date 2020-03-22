.PHONY: dev

all: generate

test:
	go test ./...

dev-build:
	go build -gcflags='-N -l' app/cmd/mangago/mangago.go

air-hook: test dev-build

base-image:
	docker build . -f docker/base/Dockerfile -t mangago-base

dev-image: base-image
	docker build . -f docker/dev/Dockerfile -t mangago-dev

dev: dev-image
	docker run -v "$(shell pwd)/app:/home/mangago/app/" -p 8080:8080 --rm mongogo-dev

tools-image: base-image
	docker build . -f docker/tools/Dockerfile -t mangago-tools

generate-image: tools-image
	docker build . -f docker/generate/Dockerfile -t mangago-generate

generate: generate-image
	docker run -v "$(shell pwd)/app:/home/mangago/app/" --rm mangago-generate
