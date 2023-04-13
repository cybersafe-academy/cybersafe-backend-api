export ENV=local


build:
	go build -o bin/main main.go

run:
	docker compose up -d
	go run main.go

swag:
	export PATH=$(go env GOPATH)/bin:$PATH
	swag fmt && swag init --parseDependency --parseInternal
