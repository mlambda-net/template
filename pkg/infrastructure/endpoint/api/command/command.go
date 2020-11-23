package command

import (
  "github.com/mlambda-net/net/pkg/net"
)

type Command interface {
  Register(r net.Route)
}

type control struct {
	template   net.Request
}

func NewCommand(template net.Request) Command {
  return &control{
    template: template,
  }
}

func (c *control) Register(r net.Route) {
  r.AddRoute("ping", "/template/ping/{id}", false, "GET", c.getPing)
}
