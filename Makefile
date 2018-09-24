all: build ## runs build

build: ## builds stuff
	go build

run: ## runs code
	go run ./schedule-splitter/main.go

test: ## tests stuff
	go test -v ./...

deploy-pachyderm: ##deploys pachyderm
	pachctl deploy local

create-repo: ##creates test repo
	pachctl create-repo schedule

