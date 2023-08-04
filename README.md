# grpc-project
Projeto para estudos de RPC


## Gerar codigos a partir do protoc

Navegue até a pastas *resources/proto* e execute os comandos abaixo.
Os arquivos gerados estarão na pasta *src/pb* dos respectivos modulos.

### Client
```
protoc --go_out=../../client/ --go-grpc_out=../../client/ *.proto
```
### Product Server
```
 protoc --go_out=../../servers/grpc-product-server/ --go-grpc_out=../../servers/grpc-product-server/ product.proto common.proto
```
### Category Server
```
 protoc --go_out=../../servers/grpc-category-server/ --go-grpc_out=../../servers/grpc-category-server/ category.proto common.proto
```

## Iniciando os servidores

### Category

Em um novo terminal, navegue até a pasta raiz do módulo, ``cd servers/grpc-category-server/src/``. 
Executando o comando abaixo o servidor vai roda na porta **50051**. 
```
 go run main.go
```

### Product

Em um novo terminal, navegue até a pasta raiz do módulo, ``cd servers/grpc-product-server/src/``.
Executando o comando abaixo o servidor vai roda na porta **50052**.
```
 go run main.go
```

### Client
Em um novo terminal, navegue até a pasta raiz do módulo, ``cd client/src/``.
Executando o comando abaixo o **client** se conectará aos servidores trazendo os dados de categoria e seus produtos associados.
```
 go run main.go
```