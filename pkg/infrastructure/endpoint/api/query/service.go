package query

import (
  "github.com/graphql-go/graphql"
  "github.com/mlambda-net/template/pkg/application/message"
)

func (c *control) getDummy(token string) func (params graphql.ResolveParams) (interface{}, error) {
  return func(params graphql.ResolveParams) (interface{}, error) {
    id, ok := params.Args["id"].(int64)
    if ok {
      return c.fetchDummy(token, id)
    }
    return message.Pong{}, nil
  }

}

func (c *control) fetchDummy(token string, id int64) (*message.Pong, error) {
  result, err := c.user.Token(token).Request(&message.Ping{
    Id :id,
  }).Unwrap()

  if err != nil {
    return nil, err
  }

  rs := result.(*message.Pong)
  return rs, nil
}
