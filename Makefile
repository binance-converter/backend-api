.DEFAULT_GOAL := gen

gen:
	rm -r api
	protoc --go_out=. --go_opt=paths=import \
	--go-grpc_out=. --go-grpc_opt=paths=import \
	proto/*.proto
	cp -r github.com/binance-converter/backend-api/api api
	rm -r github.com