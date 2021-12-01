export GO111MODULE="on"
go install github.com/dubbogo/tools/cmd/protoc-gen-go-triple@v1.0.3
protoc --go_out=./erp --go-triple_out=./erp erp.proto