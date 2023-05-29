generate:
	rm -rf todo/pb
	mkdir -p todo/pb

	protoc \
		--proto_path=proto/ \
		--go_out=internal/app/todo/pb \
		--go-grpc_out=internal/app/todo/pb \
		proto/*.proto

