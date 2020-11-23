package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/mlambda-net/monads/monad"
	"github.com/mlambda-net/template/pkg/domain/entity"
	"github.com/mlambda-net/template/pkg/domain/repository"
	"github.com/mlambda-net/template/pkg/domain/utils"
	"os"
	"strconv"
)

type templateStore struct {
	db *pg.DB
}

func (i *templateStore) Get(id int64) monad.Mono {
	 user := &entity.Dummy{}
	_, err := i.db.QueryOne(user, `SELECT * FROM dummy Where id = ?`, id)
	if err != nil {
		return monad.ToMono(err)
	}

	return monad.ToMono(user)
}

func NewTemplateStore(config *utils.Configuration) repository.TemplateStore {

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

	return &templateStore{db: db}

}

func (i *templateStore) Close() {
	_ = i.db.Close()
}
