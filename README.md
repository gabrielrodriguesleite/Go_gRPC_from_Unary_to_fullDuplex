# Go_gRPC_from_Unary_to_fullDuplex
Repositorio de estudo de Go e gRPC

```sh
# Passo a passo
go mod init go-grpc
mkdir server client pb
echo "package main" > server/server.go
echo "package main" > client/client.go
echo 'syntax = "proto3";' > pb/data.proto
```

###### REFERENCIAS
https://www.linkedin.com/pulse/grpc-go-jos%25C3%25A9-augusto-zimmermann-negreiros/