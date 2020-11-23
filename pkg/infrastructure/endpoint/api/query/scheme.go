package query

import (
  "errors"
  "fmt"
  "github.com/graphql-go/graphql"
  "github.com/mlambda-net/template/pkg/application/message"
)

func (c *control) dummyType() *graphql.Object {

  return graphql.NewObject(graphql.ObjectConfig{
    Name: "dummy",
    Fields: graphql.Fields{
      "id": &graphql.Field{
        Name: "Id",
        Type: graphql.Int,
      },
      "name": &graphql.Field{
        Name: "Name",
        Type: graphql.String,
      },
      "full": &graphql.Field{
        Type: graphql.String,

        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          if pong, ok := p.Source.(*message.Pong); ok {
            return pong.Value, nil
          }
          return nil, errors.New("the type is not correct")
        },
      },
    },
  })
}

func (c *control) dummiesQuery(token string) *graphql.Object {

  return graphql.NewObject(graphql.ObjectConfig{
    Name: "dummies",
    Fields: graphql.Fields{
      "user": &graphql.Field{
        Type:    graphql.NewList(c.dummyType()),
        Args:    c.ById(),
        Resolve: c.getDummy(token),
      },
    },
  })
}

func (c *control) ById() graphql.FieldConfigArgument {
  return graphql.FieldConfigArgument{
    "id": &graphql.ArgumentConfig{
      Type: graphql.Int,
    },
  }
}

func (c control) schema(token string) graphql.Schema  {
  var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: c.dummiesQuery(token),
  })
  return schema
}

func (c *control) exec(query string, token string)  (*graphql.Result, error) {
  result := graphql.Do(graphql.Params{
    Schema:        c.schema(token),
    RequestString: query,
  })
  if len(result.Errors) > 0 {
    return nil, fmt.Errorf("wrong result, unexpected errors: %v", result.Errors)
  }
  return result, nil
}
