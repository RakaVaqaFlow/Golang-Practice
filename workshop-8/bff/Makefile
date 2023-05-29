generate_proto:
	protoc \
		--proto_path=internal/pb/api \
		--go_out=internal/pb/api \
		--go-grpc_out=internal/pb/api \
		internal/pb/api/*.proto

generate_graphql:
	go run github.com/99designs/gqlgen generate

