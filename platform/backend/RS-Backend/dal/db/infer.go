package db

import (
	"RS-Backend/models/dao"
	"RS-Backend/models/dto"
	"RS-Backend/util"
	"context"

	"github.com/sirupsen/logrus"

)

type IInferAccesser interface {
	GetAllInferJobs(ctx context.Context) (result []*dao.RsInferJob, err error)
	GetInferJobById(ctx context.Context, id int64) (result *dto.RsInferJob, err error)
}

type InferAccesserImpl struct {
	db IDB
}

func NewInferAccesser(db IDB) IInferAccesser {
	return &InferAccesserImpl{db: db}
}

func (d *InferAccesserImpl) GetAllInferJobs(ctx context.Context) (result []*dao.RsInferJob, err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("[GetAllInferJobs] Error occurs! error=%v\n", err.Error())
		}
	}()
	var rsInferJobs []*dao.RsInferJob
	tx := d.db.Read(ctx).Find(&rsInferJobs)
	return rsInferJobs, util.GormRealError(tx)
}

func (d *InferAccesserImpl) GetInferJobById(ctx context.Context, id int64) (result *dto.RsInferJob, err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("[GetInferJobById] Error occurs! error=%v\n", err.Error())
		}
	}()
	var rsInferJob dao.RsInferJob
	var dtoRsInferJob dto.RsInferJob
	var rsModel dao.RsModel
	var dataset dao.Dataset

	tx := d.db.Read(ctx).First(&rsInferJob, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	dtoRsInferJob.ID=rsInferJob.ID
	dtoRsInferJob.Name=rsInferJob.Name
	dtoRsInferJob.Path=rsInferJob.Path
	dtoRsInferJob.State=rsInferJob.State
	dtoRsInferJob.Modelid=rsInferJob.Modelid
	dtoRsInferJob.Datasetid=rsInferJob.Datasetid
	dtoRsInferJob.CreateAt=rsInferJob.CreateAt
	dtoRsInferJob.UpdatedAt=rsInferJob.UpdatedAt
	dtoRsInferJob.DeletedAt=rsInferJob.DeletedAt
	dtoRsInferJob.FileSize=rsInferJob.FileSize

	tx = d.db.Read(ctx).First(&rsModel, "id = ?", rsInferJob.Modelid)
	if tx.Error == nil {
		dtoRsInferJob.ModelName=rsModel.Name
	} 
	tx = d.db.Read(ctx).First(&dataset, "id = ?", rsInferJob.Datasetid)
	if tx.Error == nil {
		dtoRsInferJob.DatasetName=dataset.Name
	} 
	return &dtoRsInferJob, util.GormRealError(tx)
}
