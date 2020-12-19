## Request method test
get:
	curl -s localhost:3000/products
get_1:
	curl -s localhost:3000/products/1
post:
	curl -s localhost:3000/products -d '{"name":"tea", "description":"A cup of tea", "price":0.25, "sku":"xyz-321-ab"}'
put:
	curl -sX PUT localhost:3000/products -d '{"id":1,"name":"milk", "description":"A cup of milk", "price":0.25, "sku":"xyz-321-ab"}'
delete:
	curl -sX DELETE localhost:3000/products -d '{"id":2}'

## Swagger
swagger:
	@echo Ensure you have the swagger CLI or this command will fail.
	@echo You can install the swagger CLI with: go get -u github.com/go-swagger/go-swagger/cmd/swagger
	@echo ....
	swagger generate spec -o ./swagger.yaml --scan-models
	