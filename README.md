# GO API

## Run postgres and redis
```shell
docker-compose up -d redis postgres
```

## Swagger generate
Download swag go

```shell
go get -u github.com/swaggo/swag/cmd/swag
```
Generate swagger files
```shell
swag init
```
