# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

# Variables
BUILD_DIR = build
DOCKER_IMAGE_NAME = flickrphotosearch-frontend
DOCKER_IMAGE_VERSION = v1.1

.PHONY: docker
docker: ## Build docker image from the frontend app
	docker build -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION} .

.PHONY: docker-tag
docker-tag: ## Tag docker image
	docker tag ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION} ${DOCKER_ID_USER}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION}

.PHONY: docker-push
docker-push: ## Tag docker image
	docker push ${DOCKER_ID_USER}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION}

.PHONY: run-docker
run-docker: ## Run docker app
	docker run -p 8081:80 -it --rm ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION}

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
