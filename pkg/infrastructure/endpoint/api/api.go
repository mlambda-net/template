package api

import (
  "context"
  "fmt"
  "github.com/etherlabsio/healthcheck"
  "github.com/mlambda-net/net/pkg/metrics"
  "github.com/mlambda-net/net/pkg/net"
  "github.com/mlambda-net/template/pkg/infrastructure/endpoint/api/command"
  "github.com/mlambda-net/template/pkg/infrastructure/endpoint/api/conf"
  "github.com/mlambda-net/template/pkg/infrastructure/endpoint/api/query"
  "os"
  "time"
)

type Api interface {
  GetVersion() string
  GetHost() string
  Start()
  Path() string
  Docs() int32
}

type setup struct {
  config  *conf.Configuration
  command command.Command
  query   query.Query
}

func (s *setup) Docs() int32 {
  return s.config.Docs.Port
}

func (s *setup) Path() string {
   if s.config.Docs.Path != "" {
     return fmt.Sprintf("/%s", s.config.Docs.Path)
   }
   return ""
}

func NewApi() Api  {
  return &setup{config: conf.LoadConfig()}
}

func (s *setup) GetVersion() string {
  version := os.Getenv("VERSION")
  if version == "" {
    version = "0.0.0"
  }
  return version
}

func (s *setup) GetHost() string {
  if s.config.Docs.Host == "localhost" {
    return fmt.Sprintf("%s:%d", s.config.Docs.Host, s.config.App.Port)
  }
  return s.config.Docs.Host
}

func (s *setup) Start() {
  client := net.NewClient(s.config.Remote.Server, s.config.Remote.Port)
  template := client.Actor("template")

  s.command = command.NewCommand(template)
  s.query = query.NewQuery(template)


  local := net.NewApi(s.config.App.Port, s.config.Metric.Port)
  local.Metrics(func(mc *metrics.Configuration) {
    mc.App.Name = s.config.App.Name
    mc.App.Env = s.config.Env
    mc.App.Path = "/check/template"
    mc.App.Version = s.config.App.Version
    mc.Metric.Namespace = s.config.Metric.Namespace
    mc.Metric.SubSystem = "template"
  })
  local.Register(func(r net.Route) {
    s.command.Register(r)
    s.query.Register(r)

  })

  local.Checks(
    healthcheck.WithTimeout(5 * time.Second),
    healthcheck.WithChecker("server",healthcheck.CheckerFunc(func(ctx context.Context) error {
     // client.Health()
      return nil
    })),
  )
  local.Start()

  local.Wait()

}
