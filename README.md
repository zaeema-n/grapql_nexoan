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
  entities(entitiesFilter: { id: "2360-60_doc_8" }) {
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
  entities(entitiesFilter: {id:"gov_01"}) {
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
      entities{
        id
        name{
          value
        }
      }
    }
  }
}
```

For an orgchart example, try running the following query which will retrieve
the departments under the "State Minister of Sports and Youth Affairs" during Ranil's tenure:

```bash
query {
  entities(entitiesFilter: {id:"gov_01"}) {
    id
    name { value }
    kind {minor major}
    created
    terminated
    relationships (relationshipsFilter: {name: "AS_PRESIDENT"}) {
      id
      name
      relatedEntityId
      startTime
      endTime
      direction
      entities (entitiesFilter: {name: {value: "Ranil Wickremesinghe"}}){
        id
        name{
          value
        }
        relationships(relationshipsFilter: {name: "AS_MINISTER"}){
          id
          name
          entities(entitiesFilter: {name: {value: "State Minister of Sports and Youth Affairs"}}){
            id
            name{
              value
            }
            relationships(relationshipsFilter: {name: "AS_DEPARTMENT"}){
              id
              name
              entities{
                name{value}
              }
            }
          }
        }
      }
    }
  }
}

```