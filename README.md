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

### Client
Em um novo terminal, navegue até a pasta raiz do módulo, ``cd client/src/``.
Executando o comando abaixo o **client** se conectará aos servidores trazendo os dados de categoria e seus produtos associados.
```
 go run main.go
```

### Fontes

- Protocol Buffer: https://protobuf.dev/
- gRPC: https://grpc.io/