package spec

import (
	"github.com/mlambda-net/net/pkg/spec"
	"strconv"
)


func ById(id int64) spec.Expression {
	return spec.NewEval("id", strconv.Itoa(int(id)), "=")
}

