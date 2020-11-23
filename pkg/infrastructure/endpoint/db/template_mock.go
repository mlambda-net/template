package db

import (
	"github.com/mlambda-net/monads/monad"
	"github.com/mlambda-net/net/pkg/spec"
	"github.com/mlambda-net/template/pkg/domain/entity"
)

type templateMock struct {

}

func (t *templateMock) Single(_ spec.Spec) monad.Mono {

	return monad.ToMono(&entity.Dummy{
		Id:   1,
		Name: "Dummy",
	})

}

func (t *templateMock)  Get(id int64) monad.Mono  {
	return t.Single(nil)
}

func (t *templateMock) Close() {

}