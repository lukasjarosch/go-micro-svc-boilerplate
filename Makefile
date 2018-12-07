GOPATH=$(shell go env GOPATH)
GOFILES=$(wildcard *.go)
GOFILES=$(filter-out $(wildcard *_test.go), $(wildcard *.go))
GONAME=$(shell basename "$(PWD)")
COMMIT=$(shell git rev-parse --short HEAD)


DOCKER_IMAGE="derwaldemar/go-micro-srv-boilerplate"
DOCKER_TAG="v-${COMMIT}"

.PHONY: build get run start restart clean

build: clean proto
	@GOPATH=$(GOPATH) go build -o ${GONAME} ${GOFILES}
	@echo "[+] built binary: $(GONAME)"

run: proto
	bash -c "trap 'go run $(GOFILES)' EXIT"

docker: clean proto
	@echo "[+] building docker image"
	docker build . -t ${DOCKER_IMAGE}:${DOCKER_TAG}

docker-run:
	@echo "[+] starting container $(DOCKER_IMAGE):$(DOCKER_TAG)"
	docker run  \
			--rm \
			--name $(GONAME) \
			--mount type=bind,source=${CURDIR}/config.json,target=/config.json \
			-e MICRO_REGISTRY_ADDRESS=localhost:8500 \
			-e DATABASE_URI="user:pass@tcp(localhost:3306)/test-database" \
			-e ENVIRONMENT=local \
			--network host \
			$(DOCKER_IMAGE):$(DOCKER_TAG)

docker-stop:
	@echo "[+] stopping container $(DOCKER_IMAGE):$(DOCKER_TAG)"
	@docker stop $(GONAME) > /dev/null

start-local:
	@echo "[+] starting local development services"
	docker-compose -f deploy/docker-compose.dev.yml up -d

stop-local:
	@echo "[+] stopping local development services"
	docker-compose -f deploy/docker-compose.dev.yml down

proto:
	@echo "[+] compiling protobuf"
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/example/*.proto

test:
	go test -timeout 30s -v -cover

clear:
	@clear

clean:
	@echo "[+] removing leftover pixie dust..."
	@GOPATH=$(GOPATH) go clean

start:
	@echo "Starting service binary: ./$(GONAME)"
	@./$(GONAME)

tidy:
	@echo "[+] tidying up dependencies"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod tidy
