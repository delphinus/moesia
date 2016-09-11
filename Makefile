.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

gom: ## Install gom
	go get github.com/mattn/gom

gom-update: ## Update gom
	go get -u github.com/mattn/gom

generate-template-bindata: ## Generate bindata for templates
	gom exec go-bindata -pkg vacancy -o vacancy/templates.go ./templates/...

test: generate-template-bindata ## Run tests only
	gom test -v `go list ./... | grep -v vendor`

build: generate-template-bindata ## Build binary
	gom build ./cmd/moesia

run: generate-template-bindata ## Build binary and run
	gom run ./cmd/moesia/main.go

install-dependencies: ## Install packages for dependencies
	gom install

install-test-dependencies: ## Install packages for dependencies with the `test` group
	gom -test install

assert-on-travis: ## Assert that this task is executed in Travis CI
ifeq ($(TRAVIS_BRANCH),)
	$(error No Travis CI)
endif

travis-test: assert-on-travis gom install-test-dependencies build test ## Run tests in Travis CI
