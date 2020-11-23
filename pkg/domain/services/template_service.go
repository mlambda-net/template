package services

import (

	"github.com/mlambda-net/template/pkg/domain/entity"
	"github.com/mlambda-net/template/pkg/domain/repository"
	"github.com/mlambda-net/template/pkg/domain/utils"
	"github.com/mlambda-net/template/pkg/infrastructure/endpoint/db"

)

type TemplateService interface {
	Get(id int64) (*entity.Dummy, error)
}

type templateService struct {
	repo  repository.TemplateStore
	query repository.TemplateQuery
}


func (s *templateService) Get(id int64) (*entity.Dummy, error) {
  rsp, err := s.repo.Get(id).Unwrap()
  if err != nil {
    return nil, err
  }

  dummy := rsp.(*entity.Dummy)
  return dummy, nil
}

func NewTemplateService(config *utils.Configuration) TemplateService {
	repo := db.NewTemplateStore(config)
	query := db.NewTemplateQuery(config)
	return &templateService{
		repo:  repo,
		query: query,
	}
}