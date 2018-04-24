# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

# Variables
BUILD_DIR = build
PACKAGE = ${shell go list .}
BINARY_NAME = ${shell echo ${PACKAGE}}
DOCKER_IMAGE_NAME = flickrphotosearch-video-search
DOCKER_IMAGE_VERSION = v1

include ../.make/common.mk
