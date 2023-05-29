.PHONY: generate
generate:
	buf mod update
	buf generate


MIGRATIONS_DIR=./migrations
.PHONY: migration
migration:
	goose -dir=${MIGRATIONS_DIR} create $(NAME) sql

.PHONY: .test
.test:
	$(info Running tests...)
	go test ./...

.PHONY: cover
cover:
	go test -v $$(go list ./... | grep -v -E '/Forum-test/pkg/(api)') -covermode=count -coverprofile=/tmp/c.out
	go tool cover -html=/tmp/c.out

