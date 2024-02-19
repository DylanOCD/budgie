#
# Copyright (c) 2024 Dylan O' Connor Desmond
#

.PHONY: lint build up down

IMAGE_NAME		?= budgie-api
IMAGE_TAG		?= dev
HADOLINT_IMAGE	?= hadolint/hadolint:latest
GOLANG_IMAGE	?= golang:1.22.0
WORKDIR			?= "$(shell cygpath -w $$(pwd))"

lint:
	@docker \
		run \
		--rm \
		--interactive \
		$(HADOLINT_IMAGE) \
		hadolint - < Dockerfile

build:
	@docker \
		build \
		--tag $(IMAGE_NAME):$(IMAGE_TAG) \
		.

up:
	@docker \
		compose \
		--file compose.yaml \
		--env-file backend/config/config.env \
		up

down:
	@docker \
		compose \
		--file compose.yaml \
		--env-file backend/config/config.env \
		down
