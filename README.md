# Go_gRPC_from_Unary_to_fullDuplex
Repositorio de estudo de Go e gRPC

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
https://www.linkedin.com/pulse/grpc-go-jos%25C3%25A9-augusto-zimmermann-negreiros/