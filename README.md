# starknet-indexer

Starknet Indexer provides a simple cloud service for indexing the starknet blockchain. It depends on [ent](https://entgo.io/), [gqlgen](https://github.com/99designs/gqlgen), and [caigo](https://github.com/dontpanicdao/caigo) in order generate db schemas, bind them to graphql endpoints and provide native starknet types.

To learn more about the entql binding generation see: https://entgo.io/docs/tutorial-todo-gql

## Quick start

Start indexing

```sh
go run cmd/main.go
```

Visit:

```
http://localhost:8081/playground
```

## Development

Generate schema, bindings, ect.

```sh
go generate ./...
```
