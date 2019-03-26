.PHONY: all
all: build-deps build fmt vet lint test

ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
APP_EXECUTABLE="out/sample-golang"
SERVICE_DIRS=$(shell ls -d service/*)

setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/lint/golint

build-deps:
	dep ensure

update-deps:
	dep ensure

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

build: build-deps compile

install:
	go install $(ALL_PACKAGES)

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

setup-web-ui:
	go get github.com/gocraft/work/cmd/workwebui
	go install github.com/gocraft/work/cmd/workwebui

lint:
	@for p in $(ALL_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

test:
	make lint
	ENVIRONMENT=test go test $(ALL_PACKAGES) -p=1 -coverprofile=coverage.out -covermode=count
	rm -rf coverage
	mkdir coverage
	go tool cover -html=coverage.out -o coverage/coverage.html
	go tool cover -func=coverage.out | tail -1 | awk '{print("Coverage:"$$3)}'
	go tool cover -html=coverage.out

test_ci:
	ENVIRONMENT=test go test $(ALL_PACKAGES) -p=1 -race

test-cover-html:
	@echo "mode: count" > coverage-all.out
	$(foreach pkg, $(ALL_PACKAGES),\
	ENVIRONMENT=test go test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out -o out/coverage.html

copy-config:
	cp application.yml.sample application.yml

copy-config-ci:
	cp application.yml.sample application.yml

start:
	./out/sample-golang


define generate_mock
	@for p in $(1); do \
		echo "==> Generating for $$p"; \
		mockery -dir="$$p" -output="$$p/mocks" -case underscore -all; \
	done
endef

regenerate-db-mocks:
	$(call generate_mock, "internal/db")

regenerate-service-mocks:
	$(call generate_mock, $(SERVICE_DIRS))

regenerate-mocks: regenerate-service-mocks regenerate-db-mocks