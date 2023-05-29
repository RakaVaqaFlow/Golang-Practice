.PHONY: generate
generate:
	rm -rf internal/pb/api/*.go
	
	protoc \
		--proto_path=proto/ \
		--go_out=internal/pb/api \
		--go-grpc_out=internal/pb/api \
		proto/*.proto
