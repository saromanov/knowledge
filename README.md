# knowledge
Basic backend for knowlange base

## Developers

### Running 

#### Local running

```sh
POSTGRES_HOST=localhost POSTGRES_PORT=5432 POSTGRES_USERNAME=knowledge POSTGRES_PASSWORD=knowledge POSTGRES_DB=knowledge $GOPATH/bin/go1.18beta1 run ./cmd/knowledge/main.go
```

#### Container running

```sh
make build
make run
```

### Migrations

```
migrate -source file://migrations -database postgres://username:postgres@localhost:5432/knowlage up
```

### API

Creating of the author

#### Payload

```
{
   "name":"Abcd"
}
```

```
POST /api/v1/authors
```

Creating of the page

#### Payload

