package repository

import (
	"github.com/mlambda-net/monads/monad"
)

type TemplateStore interface {
	Get(id int64) monad.Mono
	Close()
}
