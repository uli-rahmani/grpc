echo "\e[1m\e[32mGenerating gRPC File..."

#protoc -I {{Location of .proto folder}} {{Location of .proto file}} --go_put=plugins=grpc:{{Destination Location of generated .pg.go file}}
# protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto
 protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/*.proto