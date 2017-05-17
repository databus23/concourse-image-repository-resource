IMAGE:=databus23/repository-image-resource
TAG:=0.1

ifneq ($(http_proxy),)
BUILD_ARGS+= --build-arg http_proxy=$(http_proxy) --build-arg https_proxy=$(https_proxy) --build-arg no_proxy=$(no_proxy)
endif

build:
	go build -o assets/check ./cmd/check
	go build -o assets/in ./cmd/in

image: export GOOS=linux
image: build
	docker build $(BUILD_ARGS) -t $(IMAGE):$(TAG) .

.PHONY: test
test:
	docker run --rm -i $(IMAGE):$(TAG) /opt/resource/check < $(PWD)/test/$${WHAT:-gcr.io}.json

push:
	docker push $(IMAGE):$(TAG)
	docker tag $(IMAGE):$(TAG) $(IMAGE)
	docker push $(IMAGE)
