DOCKER_IMAGE=web-page-analyzer
DOCKER_TAG=latest
PORT=8080

build:
	go build -o bin/app

run:
	go run .

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./internal/services/... ./internal/utils/...
	go tool cover -html=coverage.out -o coverage.html

fmt:
	go fmt ./...

lint:
	golangci-lint run

docker-build:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-run:
	docker run -p $(PORT):$(PORT) $(DOCKER_IMAGE):$(DOCKER_TAG)
