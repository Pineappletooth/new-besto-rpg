protoc --proto_path=api --go_out=pkg/api/proto --go_opt=paths=source_relative --go-grpc_out=pkg/api/proto --go-grpc_opt=paths=source_relative commands.proto
go run cmd/core/main.go

Golangci-lint run

go build -o ./bin/main.exe ./cmd/core/main.go && ./bin/main.exe
go build ./cmd/core/main.go 
