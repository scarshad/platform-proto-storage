check:
	protoc -I ./proto \
       --go_out ./pb/go --go_opt paths=source_relative \
       --go-grpc_out ./pb/go --go-grpc_opt paths=source_relative \
       ./proto/*.proto