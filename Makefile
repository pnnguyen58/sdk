protoc:
	protoc -I . \
		--go_out=./ \
		--go-grpc_out=./ \
		--go_opt=module=sdk \
		--go-grpc_opt=module=sdk \
		--grpc-gateway_out=logtostderr=true:. \
		--grpc-gateway_opt=module=sdk \
		./proto/example/*.proto