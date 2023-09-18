# go-orm-s

```
postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
```

# [ent](https://entgo.io/)

```shell
cd gen
go run -mod=mod entgo.io/ent/cmd/ent new ${table_name} 
# ex) go run -mod=mod entgo.io/ent/cmd/ent new EntUser
cd ..
go generate ./...
```
