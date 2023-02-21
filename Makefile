compile-tracer:
	protoc -I ./proto/instrumentation/tracer/v1 \
       --go_out ./pb/instrumentation/tracer/v1 --go_opt paths=source_relative \
       --go-grpc_out ./pb/instrumentation/tracer/v1 --go-grpc_opt paths=source_relative \
       ./proto/instrumentation/tracer/v1/*.proto