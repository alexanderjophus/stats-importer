all: build ## runs build

run: build ## runs code
	./bin/stats-importer

build: ## builds stuff
	go build -o ./bin/stats-importer .

test: ## tests stuff
	go test -v ./...

docker:
	docker build -f Dockerfile -t stats-importer:latest .

deploy_local:
	kubectl apply -f ./deploy.yaml