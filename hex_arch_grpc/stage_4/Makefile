proto-msg: ## Generates Go code defining messages
proto-msg:
	protoc --go_out=internal/adapters/framework/left/grpc \
	--proto_path=internal/adapters/framework/left/grpc/proto \
	internal/adapters/framework/left/grpc/proto/number_msg.proto

proto-svc: ## Generates Go code for service
proto-svc:
	protoc --go-grpc_out=internal/adapters/framework/left/grpc \
	--proto_path=internal/adapters/framework/left/grpc/proto \
	internal/adapters/framework/left/grpc/proto/arithmetic_svc.proto

proto-all: ## Generates whole Go package based on .proto files
proto-all: proto-msg proto-svc