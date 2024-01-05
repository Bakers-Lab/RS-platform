// Infer_service.go

package service

import (
	"RS-Backend/dal/db"
	"RS-Backend/models/dao"
	"RS-Backend/models/dto"
	"context"
)

type InferService interface {
	GetAllInferJobs(ctx context.Context) ([]*dao.RsInferJob, error)
	GetInferJobById(ctx context.Context, id int64) (*dto.RsInferJob, error)
}

type InferServiceImpl struct {
	accesser db.IInferAccesser
}

func NewInferService(accesser db.IInferAccesser) InferService {
	return &InferServiceImpl{accesser: accesser}
}

func (s *InferServiceImpl) GetAllInferJobs(ctx context.Context) ([]*dao.RsInferJob, error) {
	return s.accesser.GetAllInferJobs(ctx)
}

func (s *InferServiceImpl) GetInferJobById(ctx context.Context, id int64) (*dto.RsInferJob, error) {
	return s.accesser.GetInferJobById(ctx, id)
}