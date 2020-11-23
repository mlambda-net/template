package repository

import (
	"github.com/mlambda-net/monads/monad"
	"github.com/mlambda-net/net/pkg/spec"
)

type TemplateQuery interface {
	Single(spec spec.Spec) monad.Mono
	Close()
}
