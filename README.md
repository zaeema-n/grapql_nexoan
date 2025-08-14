When you update the schema, run the following,

```bash
go run github.com/99designs/gqlgen generate
```

It will update the following files:
- schema.resolvers.go
- model/models_gen.go
- generated/generated.go

To run the graphql server, run the following,

```bash
go run server.go
```

Try running the following query in graphiql to test the Entities query

```bash
query {
  entities(filter: { id: "2360-60_doc_8" }) {
    id
    name { value }
    kind { major minor }
    created
    terminated
  }
}
```

Try Try running the following query in graphiql to test the Entities query with relationships

```bash
query {
  entities(filter: {id:"gov_01"}) {
    id
    name { value }
    kind {minor major}
    created
    terminated
    relationships (relationshipsFilter: {name: "AS_DOCUMENT"}) {
      id
      name
      relatedEntityId
      startTime
      endTime
      direction
      entity{
        id
        name{
          value
        }
      }
    }
  }
}
```