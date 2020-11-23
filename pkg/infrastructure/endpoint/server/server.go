package server

import (
	"fmt"
	"github.com/mlambda-net/net/pkg/remote"
	"github.com/mlambda-net/template/pkg/application"
	"github.com/mlambda-net/template/pkg/domain/utils"
	"github.com/sirupsen/logrus"
	"sync"
)

type Server interface {
	Start()
	Stop()
	Wait()
}

type server struct {
	wg     *sync.WaitGroup
	config *utils.Configuration
	remote remote.Server
}

func (s server) Wait() {
	s.wg.Wait()
}

func NewServer() Server {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	return server{wg: wg}
}

func (s server) Start() {

	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic occurred:", err)
		}
	}()

	s.LoadConfig()
	s.remote = remote.NewServer()
	go func() {
		logrus.Info(fmt.Sprintf("starting the server %s in the port %s", s.config.App.Name, s.config.App.Port))
		s.remote.Register("template", application.NewTemplateProps(s.config),false, []string{})
		s.remote.Start(fmt.Sprintf(":%s", s.config.App.Port))
		s.wg.Wait()
	}()
}

func (s server) Stop() {
	s.remote.Stop()
	s.wg.Done()
}
