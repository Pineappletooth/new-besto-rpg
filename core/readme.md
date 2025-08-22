protoc --go_out=pkg/api --go_opt=paths=source_relative --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative proto/commands.proto


go build -o ./bin/main.exe .\cmd\core\main.go && ./bin/main.exe