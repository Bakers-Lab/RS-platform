// Eval_service.go

package service

import (
	"RS-Backend/dal/db"
	"RS-Backend/models/dao"
	"context"
)

type EvalService interface {
	GetAllEvalJobs(ctx context.Context) ([]*dao.RsEvalJob, error)
	GetEvalJobById(ctx context.Context, id int64) (*dao.RsEvalJob, error)
}

type EvalServiceImpl struct {
	accesser db.IEvalAccesser
}

func NewEvalService(accesser db.IEvalAccesser) EvalService {
	return &EvalServiceImpl{accesser: accesser}
}

func (s *EvalServiceImpl) GetAllEvalJobs(ctx context.Context) ([]*dao.RsEvalJob, error) {
	return s.accesser.GetAllEvalJobs(ctx)
}
func (s *EvalServiceImpl) GetEvalJobById(ctx context.Context, id int64) (*dao.RsEvalJob, error) {
	return s.accesser.GetEvalJobById(ctx, id)
}