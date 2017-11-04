# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: setup
setup:: secrets.yaml minikube.secrets.yaml ## Setup kubernetes deployment

secrets.yaml: ## Create the secrets.yaml file if it does not exist.
	cp secrets.yaml.dist secrets.yaml

minikube.secrets.yaml: ## Create minikube.secrets.yaml file if it does not exist.
	cp secrets.yaml.dist minikube.secrets.yaml

.PHONY: clean
clean: ## Cleans the project
	rm -rf secrets.yaml minikube.secrets.yaml

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
