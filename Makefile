compile-tracer:
	protoc -I ./instrumentation/tracer/proto \
       --go_out ./instrumentation/tracer/pb --go_opt paths=source_relative \
       --go-grpc_out ./instrumentation/tracer/pb --go-grpc_opt paths=source_relative \
       ./instrumentation/tracer/proto/*.proto