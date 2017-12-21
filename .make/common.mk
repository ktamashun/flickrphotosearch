
.PHONY: install
install:: dep .env .env.test protoc ## Install FlickrPhotoSearch service

.PHONY: dep
dep: ## Install go dependencies
	@glide install

.env: ## Create .env file if it does not exist.
	cp .env.dist .env

.env.test: ## Create .env-test file if it does not exist.
	cp .env.dist .env.test

.PHONY: clean
clean: ## Cleans the project
	rm -rf build/ vendor/ .env .env.test apis/*

.PHONY: build
build: ## Build the FlickrPhotoSearch app
	go build -o ${BUILD_DIR}/${BINARY_NAME} ${PACKAGE}

.PHONY: build-docker
build-docker: ## Build app for Docker
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${BUILD_DIR}/${BINARY_NAME}-docker ${PACKAGE}

.PHONY: run
run: build .env ## Run the FlickrPhotoSearch app
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
	docker run -p 8101:8080 -it --rm ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION}

.PHONY: protoc
protoc: ## Generate protobuf code
	protoc -I=../protos --go_out=plugins=grpc:./apis ../protos/*.proto

.PHONY: test
test: ## Run tests
	go test ./app/... ${ARGS}

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
