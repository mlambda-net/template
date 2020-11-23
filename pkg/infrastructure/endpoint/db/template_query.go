package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/mlambda-net/monads/monad"
	"github.com/mlambda-net/net/pkg/spec"
	"github.com/mlambda-net/template/pkg/domain/entity"
	"github.com/mlambda-net/template/pkg/domain/repository"
	"github.com/mlambda-net/template/pkg/domain/utils"
	"os"
	"strconv"
)

type templateQuery struct {
	db *pg.DB
}


func (i templateQuery) Single(spec spec.Spec) monad.Mono {
	var items []entity.Dummy
	_, err := i.db.Query(&items, fmt.Sprintf("SELECT * FROM identities where %s", spec.Query()))

	if err != nil {
		monad.ToMono(err)
	}

	if len(items) > 0 {
		return monad.ToMono(items[0])
	}

	return monad.ToMono(nil)
}

func (i templateQuery) Close() {
	_ = i.db.Close()
}

func NewTemplateQuery(config *utils.Configuration) repository.TemplateQuery {

	debug, _ := strconv.ParseBool( os.Getenv("Debug"))
	if debug {
		return &templateMock{}
	}

	db := pg.Connect(&pg.Options{
		User:     config.Db.User,
		Password: config.Db.Password,
		Addr:     fmt.Sprintf("%s:%s", config.Db.Host, config.Db.Port),
		Database: config.Db.Schema,
	})
	return &templateQuery{db: db}
}
