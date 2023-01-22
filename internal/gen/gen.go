package gen

//go:generate go run github.com/ogen-go/ogen/cmd/ogen --clean --package gen --target ../../internal/gen ../../api/docs/test-transaction.yaml
//go:generate mockgen -source=oas_server_gen.go -destination=../../mocks/handler.go -package=mocks
//go:generate mockgen -source=oas_security_gen.go -destination=../../mocks/security-handler.go -package=mocks
