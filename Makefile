.PHONY: mocks
MODULE:="github.com/esklo/residents-tracking-platform-backend"
generate:
	rm -rf gen/*
	mkdir gen/proto gen/openapi
	cd api && protoc \
		--go_out ../ \
		--go_opt=module=${MODULE} \
        --go-grpc_out ../ \
        --go-grpc_opt=module=${MODULE} \
        --grpc-gateway_out ../ \
        --grpc-gateway_opt=generate_unbound_methods=true \
        --grpc-gateway_opt=module=${MODULE} \
        --openapiv2_out ../gen/openapi \
        --openapiv2_opt=generate_unbound_methods=true \
        --openapiv2_opt allow_merge=true \
        */*.proto
mocks:
	mockgen -source=internal/repository/repository.go -destination=mocks/repository.go -package=mocks
	mockgen -source=internal/service/service.go -destination=mocks/service.go -package=mocks