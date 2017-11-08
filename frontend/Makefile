# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

# Variables
BUILD_DIR = build
PACKAGE = ${shell go list .}
BINARY_NAME = ${shell echo ${PACKAGE}}
DOCKER_IMAGE_NAME = flickrphotosearch-frontend
DOCKER_IMAGE_VERSION = v1.1

.PHONY: install
install:: .env .env.test ## Install FlickrPhotoSearch service

.env: ## Create .env file if it does not exist.
	cp .env.dist .env

.env.test: ## Create .env-test file if it does not exist.
	cp .env.dist .env.test

.PHONY: clean
clean: ## Cleans the project
	rm -rf build/ .env .env.test

.PHONY: build
build: ## Build the FlickrPhotoSearch app
	go build -o ${BUILD_DIR}/${BINARY_NAME} ${PACKAGE}

.PHONY: build-docker
build-docker: ## Build app for Docker
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${BUILD_DIR}/${BINARY_NAME}-docker ${PACKAGE}

.PHONY: run
run: build ## Run the FlickrPhotoSearch frontend app
	${BUILD_DIR}/${BINARY_NAME} ${ARGS}

.PHONY: docker
docker: build-docker ## Build docker image from the app
	docker build --build-arg BUILD_DIR=${BUILD_DIR} --build-arg BINARY_NAME=${BINARY_NAME}-docker -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION} .

.PHONY: docker-tag
docker-tag: ## Tag docker image
	docker tag ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION} ${DOCKER_ID_USER}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION}

.PHONY: docker-push
docker-push: ## Tag docker image
	docker push ${DOCKER_ID_USER}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION}

.PHONY: run-docker
run-docker: ## Run docker app
	docker run -p 8100:80 -it --rm ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION}

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
