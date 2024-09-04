# ----------------------------------------------------------------------------------------------------------------------
# ENV
# ----------------------------------------------------------------------------------------------------------------------
include ./etc/files/.env
export $(shell sed 's/=.*//' ./etc/files/.env)

BUILD_PATH := ./build/
BUILD_BIN_PATH := $(BUILD_PATH)bin/
BUILD_DOCKER_FILE_NAME := $(BUILD_PATH)build.dockerfile
BUILD_SERVICE_ENTRYPOINT_SCRIPT_NAME := $(BUILD_PATH)$(LOCAL_SCRIPT_PATH)service-entrypoint.sh

SERVICE_ORDER_IMAGE_TAG := $(CONTAINER_IMAGE_USERNAME)/$(SERVICE_ORDER_CONTAINER_NAME):$(SERVICE_ORDER_IMAGE_VERSION)
SERVICE_PRODUCT_IMAGE_TAG := $(CONTAINER_IMAGE_USERNAME)/$(SERVICE_PRODUCT_CONTAINER_NAME):$(SERVICE_PRODUCT_IMAGE_VERSION)
SERVICE_SHOP_IMAGE_TAG := $(CONTAINER_IMAGE_USERNAME)/$(SERVICE_SHOP_CONTAINER_NAME):$(SERVICE_SHOP_IMAGE_VERSION)
SERVICE_USER_IMAGE_TAG := $(CONTAINER_IMAGE_USERNAME)/$(SERVICE_USER_CONTAINER_NAME):$(SERVICE_USER_IMAGE_VERSION)
SERVICE_WAREHOUSE_IMAGE_TAG := $(CONTAINER_IMAGE_USERNAME)/$(SERVICE_WAREHOUSE_CONTAINER_NAME):$(SERVICE_WAREHOUSE_IMAGE_VERSION)

# ----------------------------------------------------------------------------------------------------------------------
# EXEC
# ----------------------------------------------------------------------------------------------------------------------
.PHONY: run
run: down up

.PHONY: up
up: build
	cd $(BUILD_PATH) && bash cmd.sh start

.PHONY: down
down:
	cd $(BUILD_PATH) && bash cmd.sh stop || true

.PHONY: build
build: prepare-build build-bin build-container
	rm -rf $(BUILD_BIN_PATH)
	rm -rf $(BUILD_DOCKER_FILE_NAME)
	rm $(BUILD_PATH)/.env.example
	rm $(BUILD_SERVICE_ENTRYPOINT_SCRIPT_NAME)

# ----------------------------------------------------------------------------------------------------------------------
# BUILDER
# ----------------------------------------------------------------------------------------------------------------------
.PHONY: prepare-build
prepare-build:
	rm -rf $(BUILD_PATH)
	mkdir -p $(BUILD_PATH)
	cp -r etc/files/ $(BUILD_PATH)

.PHONY: build-bin
build-bin:
	rm -rf $(BUILD_BIN_PATH)
	mkdir -p $(BUILD_BIN_PATH)

	@go mod tidy
	GOOS=linux GOARCH=amd64 go build -trimpath -o $(BUILD_BIN_PATH)$(SERVICE_ORDER_CONTAINER_NAME) -v -ldflags=-checklinkname=0 cmd/order/main.go
	GOOS=linux GOARCH=amd64 go build -trimpath -o $(BUILD_BIN_PATH)$(SERVICE_PRODUCT_CONTAINER_NAME) -v -ldflags=-checklinkname=0 cmd/product/main.go
	GOOS=linux GOARCH=amd64 go build -trimpath -o $(BUILD_BIN_PATH)$(SERVICE_SHOP_CONTAINER_NAME) -v -ldflags=-checklinkname=0 cmd/shop/main.go
	GOOS=linux GOARCH=amd64 go build -trimpath -o $(BUILD_BIN_PATH)$(SERVICE_USER_CONTAINER_NAME) -v -ldflags=-checklinkname=0 cmd/user/main.go
	GOOS=linux GOARCH=amd64 go build -trimpath -o $(BUILD_BIN_PATH)$(SERVICE_WAREHOUSE_CONTAINER_NAME) -v -ldflags=-checklinkname=0 cmd/warehouse/main.go

.PHONY: build-container
build-container:
	docker rmi $(SERVICE_ORDER_CONTAINER_NAME) --no-prune || true
	docker build \
		-f $(BUILD_DOCKER_FILE_NAME) \
		-t $(SERVICE_ORDER_IMAGE_TAG) \
		--build-arg BINARY_PATH_SRC=$(BUILD_BIN_PATH)$(SERVICE_ORDER_CONTAINER_NAME) \
		--build-arg BINARY_PATH_DST=$(CONTAINER_BIN_PATH)$(SERVICE_ORDER_CONTAINER_NAME) \
		--build-arg LOG_PATH=$(CONTAINER_LOG_PATH)$(SERVICE_ORDER_CONTAINER_NAME)$(LOG_EXTENSION) \
		--build-arg INIT_SCRIPT_PATH_SRC=$(BUILD_SERVICE_ENTRYPOINT_SCRIPT_NAME) \
		.

	docker rmi $(SERVICE_PRODUCT_CONTAINER_NAME) --no-prune || true
	docker build \
		-f $(BUILD_DOCKER_FILE_NAME) \
		-t $(SERVICE_PRODUCT_IMAGE_TAG) \
		--build-arg BINARY_PATH_SRC=$(BUILD_BIN_PATH)$(SERVICE_PRODUCT_CONTAINER_NAME) \
		--build-arg BINARY_PATH_DST=$(CONTAINER_BIN_PATH)$(SERVICE_PRODUCT_CONTAINER_NAME) \
		--build-arg LOG_PATH=$(CONTAINER_LOG_PATH)$(SERVICE_PRODUCT_CONTAINER_NAME)$(LOG_EXTENSION) \
		--build-arg INIT_SCRIPT_PATH_SRC=$(BUILD_SERVICE_ENTRYPOINT_SCRIPT_NAME) \
		.

	docker rmi $(SERVICE_SHOP_CONTAINER_NAME) --no-prune || true
	docker build \
		-f $(BUILD_DOCKER_FILE_NAME) \
		-t $(SERVICE_SHOP_IMAGE_TAG) \
		--build-arg BINARY_PATH_SRC=$(BUILD_BIN_PATH)$(SERVICE_SHOP_CONTAINER_NAME) \
		--build-arg BINARY_PATH_DST=$(CONTAINER_BIN_PATH)$(SERVICE_SHOP_CONTAINER_NAME) \
		--build-arg LOG_PATH=$(CONTAINER_LOG_PATH)$(SERVICE_SHOP_CONTAINER_NAME)$(LOG_EXTENSION) \
		--build-arg INIT_SCRIPT_PATH_SRC=$(BUILD_SERVICE_ENTRYPOINT_SCRIPT_NAME) \
		.

	docker rmi $(SERVICE_USER_CONTAINER_NAME) --no-prune || true
	docker build \
		-f $(BUILD_DOCKER_FILE_NAME) \
		-t $(SERVICE_USER_IMAGE_TAG) \
		--build-arg BINARY_PATH_SRC=$(BUILD_BIN_PATH)$(SERVICE_USER_CONTAINER_NAME) \
		--build-arg BINARY_PATH_DST=$(CONTAINER_BIN_PATH)$(SERVICE_USER_CONTAINER_NAME) \
		--build-arg LOG_PATH=$(CONTAINER_LOG_PATH)$(SERVICE_USER_CONTAINER_NAME)$(LOG_EXTENSION) \
		--build-arg INIT_SCRIPT_PATH_SRC=$(BUILD_SERVICE_ENTRYPOINT_SCRIPT_NAME) \
		.

	docker rmi $(SERVICE_WAREHOUSE_CONTAINER_NAME) --no-prune || true
	docker build \
		-f $(BUILD_DOCKER_FILE_NAME) \
		-t $(SERVICE_WAREHOUSE_IMAGE_TAG) \
		--build-arg BINARY_PATH_SRC=$(BUILD_BIN_PATH)$(SERVICE_WAREHOUSE_CONTAINER_NAME) \
		--build-arg BINARY_PATH_DST=$(CONTAINER_BIN_PATH)$(SERVICE_WAREHOUSE_CONTAINER_NAME) \
		--build-arg LOG_PATH=$(CONTAINER_LOG_PATH)$(SERVICE_WAREHOUSE_CONTAINER_NAME)$(LOG_EXTENSION) \
		--build-arg INIT_SCRIPT_PATH_SRC=$(BUILD_SERVICE_ENTRYPOINT_SCRIPT_NAME) \
		.

# ----------------------------------------------------------------------------------------------------------------------
# ETC
# ----------------------------------------------------------------------------------------------------------------------
.PHONY: push-services
push-services: build
	docker push $(SERVICE_WAREHOUSE_IMAGE_TAG)
