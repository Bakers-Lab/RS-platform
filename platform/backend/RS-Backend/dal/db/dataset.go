package db

import (
	"RS-Backend/models/dao"
	"RS-Backend/models/dto"
	"RS-Backend/util"
	"context"

	"github.com/sirupsen/logrus"

)

type IDatasetAccesser interface {
	GetAllDatasets(ctx context.Context) (result []*dao.Dataset, err error)
	GetDatasetById(ctx context.Context, id int64) (result *dto.Dataset, err error)
	InsertDataset(ctx context.Context, dataset *dao.Dataset) (err error)
	FindDatasetPath(ctx context.Context, id int64, name string) (result string, err error)
}

type DatasetAccesserImpl struct {
	db IDB
}

func NewDatasetAccesser(db IDB) IDatasetAccesser {
	return &DatasetAccesserImpl{db: db}
}

func (d *DatasetAccesserImpl) GetAllDatasets(ctx context.Context) (result []*dao.Dataset, err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("[GetAllDatasets] Error occurs! error=%v\n", err.Error())
		}
	}()
	var datasets []*dao.Dataset
	tx := d.db.Read(ctx).Find(&datasets)
	return datasets, util.GormRealError(tx)
}

func (d *DatasetAccesserImpl) GetDatasetById(ctx context.Context, id int64) (result *dto.Dataset, err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("[GetDatasetById] Error occurs! error=%v\n", err.Error())
		}
	}()
	var dataset dao.Dataset
	var datasetBatches []dao.DatasetBatch
	var dtoDataset dto.Dataset
	tx := d.db.Read(ctx).First(&dataset, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = d.db.Read(ctx).Find(&datasetBatches, "datasetid = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// CreaterUserId int64
	// StoreFormat string       // 数据集保存格式
	// State       DatasetState // 数据集保存状态
	// Batches		[]DatasetBatch
	dtoDataset.ID = dataset.ID
	dtoDataset.Name = dataset.Name
	dtoDataset.Comment = dataset.Comment
	dtoDataset.Path = dataset.Path
	dtoDataset.CreaterUserId = 0
	dtoDataset.StoreFormat = dataset.StoreFormat
	dtoDataset.State = dataset.State
	dtoDataset.Batches = datasetBatches
	return &dtoDataset, util.GormRealError(tx)
}

func (d *DatasetAccesserImpl) InsertDataset(ctx context.Context, dataset *dao.Dataset) (err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("[InsertDataset] Error occurs! error=%v\n", err.Error())
		}
	}()
	tx := d.db.Write(ctx).Create(dataset)
	return util.GormRealError(tx)
}

func (d *DatasetAccesserImpl) FindDatasetPath(ctx context.Context, id int64, name string) (result string, err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("[UploadFile] Error occurs! error=%v\n", err.Error())
		}
	}()
	var dataset dao.Dataset
	tx := d.db.Read(ctx).First(&dataset, "id = ?", id)
	if tx.Error != nil {
		return "", tx.Error
	}
	
	path :=dataset.Path


	return path, util.GormRealError(tx)
}