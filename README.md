# Rust drop project on Go

## To set up .env file with connections to db's:
##### MongoDB Connection
``` env
MONGODB_HOST=go-mongodb
MONGODB_PORT=27017
MONGODB_DBNAME=api
MONGODB_USER=admin
MONGODB_PASSWORD=admin
MONGODB_AUTH_MECHANISM=SCRAM-SHA-1
MONGODB_AUTH_DATABASE=admin
```
##### MySQL Connection 
``` env
MYSQL_HOST=go-mysql
MYSQL_PORT=3306
MYSQL_DBNAME=api
MYSQL_USER=root
MYSQL_PASSWORD=root
```

## Makefile commands
- `make build` - using by air plugin of live reload
- `make docker` or `make` - build and up project stack

## Documentation for creating newest data in project
- [Enum](https://github.com/popcornrus/go-boilerplate/issues/1)
- [Migrations](https://github.com/popcornrus/go-boilerplate/issues/2)
