NAME := khaos
REGISTRY := rene00
VERSION := $(shell git describe --tags --always --dirty)

docker-build:
	docker build . -t $(NAME):$(VERSION)

docker-tag: docker-build
	docker tag $(NAME):$(VERSION) $(NAME):latest
	docker tag $(NAME):$(VERSION) $(REGISTRY)/$(NAME):latest

docker-push: docker-tag
	docker push $(REGISTRY)/$(NAME)
