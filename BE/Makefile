PACKAGES := $(shell go list ./...)
name := $(shell basename ${PWD})
start:
	swag init --parseDependency github.com/volatiletech/null/v8
	swag fmt
	air -c .air.toml

.PHONY: vet
vet: 
	go vet $(PACKAGES)

.PHONY: init
init:
	go install github.com/cosmtrek/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/a-h/templ/cmd/templ@latest
	


.PHONY: test
test:
	go mod tidy
	go test -race -cover $(PACKAGES)
	
.PHONY: docker-start
docker-start:
	docker compose -f docker-compose.yaml  up -d --build --remove-orphans


.PHONY: docker-stop
docker-stop:
	docker compose -f docker-compose.yaml  down

.PHONY: docker-clean
docker-clean:
	docker system prune
	docker volume rm $(docker volume ls -q)
