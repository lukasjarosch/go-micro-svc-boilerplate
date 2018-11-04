GOPATH=$(shell go env GOPATH)
GOFILES=$(wildcard *.go)
GOFILES=$(filter-out $(wildcard *_test.go), $(wildcard *.go))
GONAME=$(shell basename "$(PWD)")

DOCKER_IMAGE="derwaldemar/template-srv"
DOCKER_TAG="latest"

.PHONY: build get run start restart clean

build: clean proto
	@GOPATH=$(GOPATH) go build -o ${GONAME} ${GOFILES}
	@echo "[+] built binary: $(GONAME)"

run: proto
	bash -c "trap 'go run $(GOFILES)' EXIT"

dep:
	@echo "[+] initializing dependency tree"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) dep init -v

ensure:
	@echo "[+] ensuring dependencies"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) dep ensure -v

docker: clean ensure proto
	@echo "[+] building static binary"
	@GOPATH=$(GOPATH) CGO_ENABLED=0 GOOS=linux go build -o $(GONAME) -a -installsuffix cgo .
	@echo "[+] building docker image"
	docker build . -t $(DOCKER_IMAGE):$(DOCKER_TAG)

docker-run:
	@echo "[+] starting container $(DOCKER_IMAGE):$(DOCKER_TAG)"
	docker run -d \
			--rm \
			--name $(GONAME) \
			--mount type=bind,source=${CURDIR}/config.json,target=/config.json \
			-e MICRO_REGISTRY_ADDRESS=localhost:8500 \
			--network host \
			$(DOCKER_IMAGE):$(DOCKER_TAG)

docker-stop:
	@echo "[+] stopping container $(DOCKER_IMAGE):$(DOCKER_TAG)"
	@docker stop $(GONAME) > /dev/null

proto:
	@echo "[+] compiling protobuf"
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/example/example.proto

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