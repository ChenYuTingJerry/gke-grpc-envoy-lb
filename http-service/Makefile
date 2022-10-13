PROJ="17media"
COMPOSE_FILE="deployments/docker-compose.yaml"

tidy:
	go mod tidy

go.sum:
	go get -d ./...

di: go.sum
	go install github.com/google/wire/cmd/wire@v0.5.0
	wire gen ./internal/injector

mock:
	go install github.com/golang/mock/mockgen@v1.6.0
	go generate -x -run mockgen ./...

build:
	rm -Rf ssh_keys
	cp -r ~/.ssh ssh_keys
	docker build -t http-service:latest -f . .

run: build
	docker-compose -p $(PROJ) -f $(COMPOSE_FILE) up -d

down:
	docker-compose -p $(PROJ) -f $(COMPOSE_FILE) down


.PHONY: test
test:
	curl -v --location --request GET  http://localhost:8080/api/v1/demo
