# grpc-project
Projeto para estudos de RPC

# RPC / gRPC

Remote Procedure Call (RPC), Chamada de Procedimento Remoto, é um protocolo de alto nível que utiliza uma camada mais baixo nível de transporte, como o TCP/IP ou UDP.
 Diferente do REST, que utiliza o protocolo HTTP e transferência de dados em texto, o RPC trafega dados em formato binário, 
fazendo com que a transmissão seja bem mais rápida e eficiente.

O gRPC é um framework RPC, open source, que permite o desenvolimento de aplicações em diversas linguagens, utilizando do Protocol Buffer para serialização das estruturas de dados.

Os principais casos de uso são:

- Uma conexão eficiente e poliglota em uma arquitetura de microserviços
- Conectando clientes mobile e browsers com os serviços backend
- Streaming bi-direcional usando http/2

## Gerar codigos apartir do .proto

Para gerar o código de um arquivo ```.proto``` é utilizado o Protocol Buffer Compiler ```protoc```. 

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

## Rodando o Client e obtendo os resultados
Em um novo terminal, navegue até a pasta raiz do módulo, ``cd client/src/``.
Executando o comando abaixo o **client** se conectará aos servidores trazendo os dados de categoria e seus produtos associados.
```
 go run main.go
```
O Resultado esperado são duas categorias, das quais na primeira, existem dois produtos relacionados.

```
Category: id:"5001a184-1242-42e2-a1da-200d3017fb4e" name:"Category 1" description:"Category 1"
  Product: id:"e377227e-79a2-4c3d-92e6-133a79baea1d" name:"Product 1" description:"Product 1"
  Product: id:"e9b03e26-1858-4678-8d87-bf3280029159" name:"Product 2" description:"Product 2"
Category: id:"e70e27b3-7ced-47c4-acbc-875de4640131" name:"Category 2" description:"Category 2"

```

### Fontes

- Protocol Buffer: https://protobuf.dev/
- gRPC: https://grpc.io/