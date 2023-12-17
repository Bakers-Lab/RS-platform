// dataset_service.go

package service

import (
	"RS-Backend/dal/db"
	"RS-Backend/models/dao"
	"context"
)

type DatasetService interface {
	GetAllDatasets(ctx context.Context) ([]*dao.Dataset, error)
	GetDatasetById(ctx context.Context, id int64) (*dao.Dataset, error)
	InsertDataset(ctx context.Context, dataset *dao.Dataset) error
}

type DatasetServiceImpl struct {
	accesser db.IDatasetAccesser
}

func NewDatasetService(accesser db.IDatasetAccesser) DatasetService {
	return &DatasetServiceImpl{accesser: accesser}
}

func (s *DatasetServiceImpl) GetAllDatasets(ctx context.Context) ([]*dao.Dataset, error) {
	return s.accesser.GetAllDatasets(ctx)
}

func (s *DatasetServiceImpl) GetDatasetById(ctx context.Context, id int64) (*dao.Dataset, error) {
	return s.accesser.GetDatasetById(ctx, id)
}

func (s *DatasetServiceImpl) InsertDataset(ctx context.Context, dataset *dao.Dataset) error {
	dataset.State = dao.StateReady
	return s.accesser.InsertDataset(ctx, dataset)
}
