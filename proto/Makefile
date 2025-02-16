.PHONY: requirements api cmd msg

PROTOC_ARGS := -I=. --go_out=../internal/api --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

all: requirements msg api cmd

requirements:
	go mod download
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

msg:
	protoc ${PROTOC_ARGS} msg/endnode_uplink.proto

api:


cmd:
	

