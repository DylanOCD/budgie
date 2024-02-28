#
# Copyright (c) 2024 Dylan O' Connor Desmond
#

.PHONY: lint build test up down

IMAGE_NAME		?= budgie-api
IMAGE_TAG		?= dev
HADOLINT_IMAGE	?= hadolint/hadolint:latest
GOLANG_IMAGE	?= golang:1.22.0

ifeq ($(OS), Windows_NT)
	WORKDIR	?= "$(shell cygpath -w $$(pwd))"
else
	WORKDIR ?= `pwd`
endif

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

test:
	@go \
		test \
		-count=1 \
		./...

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
