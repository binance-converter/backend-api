.DEFAULT_GOAL := gen

gen:
	rm -r api
	protoc --go_out=. --go_opt=paths=import --go_opt=module=github.com/binance-converter/backend-api \
	--go-grpc_out=. --go-grpc_opt=paths=import --go-grpc_opt=module=github.com/binance-converter/backend-api \
	proto/*.proto