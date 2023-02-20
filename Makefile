compile-tracer:
	protoc -I ./instru/tracer/proto \
       --go_out ./instru/pb --go_opt paths=source_relative \
       --go-grpc_out ./instru/pb --go-grpc_opt paths=source_relative \
       ./instrumentation/tracer/proto/*.proto