# knowledge
Basic backend for knowlange base

## Developers

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

