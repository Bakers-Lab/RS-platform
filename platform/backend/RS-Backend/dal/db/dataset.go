package db

import (
	"RS-Backend/models/dao"
	"RS-Backend/util"
	"context"

	"github.com/sirupsen/logrus"
)

type IDatasetAccesser interface {
	GetAllDatasets(ctx context.Context) (result []*dao.Dataset, err error)
	GetDatasetById(ctx context.Context, id int64) (result *dao.Dataset, err error)
	InsertDataset(ctx context.Context, dataset *dao.Dataset) (err error)
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

func (d *DatasetAccesserImpl) GetDatasetById(ctx context.Context, id int64) (result *dao.Dataset, err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("[GetDatasetById] Error occurs! error=%v\n", err.Error())
		}
	}()
	var dataset dao.Dataset
	tx := d.db.Read(ctx).First(&dataset, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dataset, util.GormRealError(tx)
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
