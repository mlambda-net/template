package spec

import (
	"github.com/mlambda-net/net/pkg/spec"
	"github.com/mlambda-net/template/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FilterById(t *testing.T) {

	dummy := getDummy()
	s := spec.Specify(ById(dummy.Id))

	assert.Equal(t, s.Query(), "id = '1'")
}

func getDummy() entity.Dummy {
	return entity.Dummy{
		Id:       1,
		Name:     "Roy",
	}
}
