PROTO_REPO := https://github.com/bactruongvan17/taskhub-backend-protobuf
TEMP_DIR := proto_temp
SRC_PROTO_DIR := src/pkg/proto
PROTO_FILES := $(wildcard $(SRC_PROTO_DIR)/*.proto)

proto_clone:
	git clone $(PROTO_REPO) $(TEMP_DIR)
	mv $(TEMP_DIR)/* $(SRC_PROTO_DIR)/
	rm -rf $(TEMP_DIR)
	rm -rf $(SRC_PROTO_DIR)/.git

proto_build:
	protoc --go_out=$(SRC_PROTO_DIR) --go-grpc_out=$(SRC_PROTO_DIR) $(PROTO_FILES)
