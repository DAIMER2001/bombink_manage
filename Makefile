run:
	go run cmd/main.go

deps:
	go install github.com/swaggo/swag/cmd/swag@latest && \
	goproxy=direct \
	GOSUMDB=off \
	go mod tidy

test:
	go test ./...  -cover -coverprofile=coverage.out 

gen-mocks:
	mockery --dir=./internal --all --output=internal/application/usecases/mocks

watch-coverage:
	go tool cover -html=coverage.out

gen-api-docs:
	swag init -g internal/infrastructure/entrypoints/apiRest/routes/bombink.routes.go --output internal/infrastructure/entrypoints/apiRest/routes/docs

check-linter:
	golangci-lint run --path-prefix=back/commerce/ms-manage-misions -v --config=../../../infra/lib/cloud-lint/golangci-lint.yaml
