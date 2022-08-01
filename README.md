# Go_gRPC_from_Unary_to_fullDuplex
Repositorio de estudo de Go e gRPC

## Executar o fluxo bi direcional

Para executar o projeto certifique-se de ter a versão 18 do Go instalada

Em um terminal execute `go run server/*`

Em outro terminal execute `go run client/*`

## Projeto inicializado com os seguintes comandos

```sh
# Passo a passo
go mod init go-grpc
mkdir server client pb
echo "package main" > server/server.go
echo "package main" > client/client.go
echo 'syntax = "proto3";' > pb/data.proto
# completar pb/data.proto
make # para gerar os arquivos grpc e pb
go mod tidy # para instalar as dependencias
```

###### REFERENCIAS

**↓ so much errors to fix but ok (be aware)**

https://www.linkedin.com/pulse/grpc-go-jos%25C3%25A9-augusto-zimmermann-negreiros/

## TODO

1. Adicionar flag para selecionar o tipo de request

2. Adicionar flag para utilizar valores personalizados