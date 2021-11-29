# Go GRPC Service

________________

## Run the project localy

```
go mod vendor

docker-compose up --build

```
__________________

## structure


cmd ==> contains entry points

cmd/server==> because we will initialise a server

cmd/graphics ==> * example : if we want to initialise a graphique ui*


internal ===> contains our logic

db ==> package for the postgress database

rocket ==> package for the service

transport ==> package to define the transport layer, here the grpc

_______________________

## Help commandes (docker)
```
docker run --name rocket-inventory-db -e POSTGRES_PASSWORD=postgres -p 5432:5342 -d postgres
```

