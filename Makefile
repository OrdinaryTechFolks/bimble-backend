EXCLUDE_THIRD_PARTY=--exclude-path third_party/errors --exclude-path third_party/google --exclude-path third_party/openapi --exclude-path third_party/validate

.PHONY: setup
setup:
	go mod vendor
	go install github.com/cespare/reflex@latest

.PHONY: api
api:
	buf generate ${EXCLUDE_THIRD_PARTY} --path api/v1

build:
	go build -v -o bin/app-http cmd/app-http/*.go

.PHONY: start-dev
start-dev:
	make api
	reflex -r "\.(go|yaml)" -s -- sh -c "make build && ./bin/app-http -config=./files/config/development.yaml"