$PROTOC_GEN_TS_PATH = Resolve-Path "../nimarh/node_modules/.bin/protoc-gen-ts.cmd"

protoc --grpc-web_out="import_style=typescript,mode=grpcweb:../nimarh/src" nimar.proto
protoc --go-grpc_out=. --go_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --plugin="protoc-gen-ts=$PROTOC_GEN_TS_PATH" --ts_out="../nimarh/src" --js_out="import_style=commonjs,binary:../nimarh/src" nimar.proto

