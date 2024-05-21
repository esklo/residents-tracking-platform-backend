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
.PHONY: graph
RANKDIRS = LR TB
GROUPTYPE = pkg type pkg,type type,pkg
FOCUSES = github.com/esklo/residents-tracking-platform-backend/internal/app github.com/esklo/residents-tracking-platform-backend/internal/app/provider github.com/esklo/residents-tracking-platform-backend/internal/model github.com/esklo/residents-tracking-platform-backend/internal/repository/user github.com/esklo/residents-tracking-platform-backend/internal/config
graph:
	@for rankdir in $(RANKDIRS); do \
        for group in $(GROUPTYPE); do \
          for focus in $(FOCUSES); do \
            focus_file=$$(echo $$focus | sed 's/\//_/g'); \
            go-callvis -file "graphvis/$${focus_file}_gv_$${rankdir}_$${group}" -graphviz -rankdir $$rankdir -group $$group -nostd -ignore github.com/esklo/residents-tracking-platform-backend/gen/proto,google,\*google,github.com/pkg,\*github.com/pkg -focus $$focus -format png cmd/api/main.go & \
          done; \
          wait; \
        done; \
    done;