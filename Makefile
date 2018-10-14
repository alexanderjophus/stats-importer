all: run ## runs build

run: build ## runs code
	./bin/stats-importer

build: ## builds stuff
	go build -o ./bin/stats-importer ./stats-importer

test: ## tests stuff
	go test -v ./...

deploy-pachyderm: ##deploys pachyderm
	pachctl deploy local

create-repo: ##creates test repo
	pachctl create-repo statsapi

dd:
	eval $(minikube docker-env)
