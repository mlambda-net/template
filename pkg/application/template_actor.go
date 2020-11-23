package application

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/mlambda-net/net/pkg/core"
	"github.com/mlambda-net/template/pkg/application/message"
	"github.com/mlambda-net/template/pkg/domain/services"
	"github.com/mlambda-net/template/pkg/domain/utils"
)

type templateActor struct {
	service services.TemplateService
}

func NewTemplateProps(config *utils.Configuration) *actor.Props {
	service := services.NewTemplateService(config)
	return actor.PropsFromProducer(func() actor.Actor { return &templateActor{service: service} })
}

func (a *templateActor) Receive(ctx actor.Context) {

	switch msg := ctx.Message().(type) {

	case *message.Ping:
		r, err := a.ping(msg)
		if err != nil {
			ctx.Respond(err)
		} else {
			ctx.Respond(r)
		}
	}
}

func (a *templateActor) ping(msg *message.Ping) (*message.Pong, *core.Error) {
	resp, err := a.service.Get(msg.Id)
	if err != nil {
		return nil, &core.Error{
			Status:  500,
			Message: err.Error(),
		}
	}

	return &message.Pong{Value: fmt.Sprintf("pong %s", resp.Name)}, nil
}
