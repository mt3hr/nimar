protoc.exe --go-grpc_out=. --go_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --grpc-web_out=import_style=typescript,mode=grpcwebtext:../nimarh/src nimar.proto
