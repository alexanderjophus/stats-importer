all: build ## runs build

build: ## builds stuff
	go build -o ./bin/stats-importer ./stats-importer 

run: build ## runs code
	./bin/stats-importer

test: ## tests stuff
	go test -v ./...

deploy-pachyderm: ##deploys pachyderm
	pachctl deploy local

create-repo: ##creates test repo
	pachctl create-repo statsapi

