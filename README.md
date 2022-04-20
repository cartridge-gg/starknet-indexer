# starknet-indexer

Starknet Indexer provides a simple cloud service for indexing the starknet blockchain. It depends on [ent](https://entgo.io/), [gqlgen](https://github.com/99designs/gqlgen), and [caigo](https://github.com/dontpanicdao/caigo) in order generate db schemas, bind them to graphql endpoints and provide native starknet types.

To learn more about the entql binding generation see: https://entgo.io/docs/tutorial-todo-gql

# TODO:
- [ ] Implement starknet contract binding generation (see: https://github.com/dontpanicdao/caigo/issues/1)
- [ ] Implement starknet processor generation (see: https://github.com/withtally/synceth/blob/main/codegen/processor.go)

## Quick start

Standup the service using a sqlite db

```sh
go run cmd/main.go
```

Generate schema, bindings, ect.
```sh
go generate ./...
```
