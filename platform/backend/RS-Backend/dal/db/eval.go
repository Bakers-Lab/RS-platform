package db

import (
	"RS-Backend/models/dao"
	"RS-Backend/util"
	"context"

	"github.com/sirupsen/logrus"

)

type IEvalAccesser interface {
	GetAllEvalJobs(ctx context.Context) (result []*dao.RsEvalJob, err error)
	GetEvalJobById(ctx context.Context, id int64) (result *dao.RsEvalJob, err error)
}

type EvalAccesserImpl struct {
	db IDB
}

func NewEvalAccesser(db IDB) IEvalAccesser {
	return &EvalAccesserImpl{db: db}
}

func (d *EvalAccesserImpl) GetAllEvalJobs(ctx context.Context) (result []*dao.RsEvalJob, err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("[GetAllEvalJobs] Error occurs! error=%v\n", err.Error())
		}
	}()
	var rsEvalJobs []*dao.RsEvalJob
	tx := d.db.Read(ctx).Find(&rsEvalJobs)
	return rsEvalJobs, util.GormRealError(tx)
}

func (d *EvalAccesserImpl) GetEvalJobById(ctx context.Context, id int64) (result *dao.RsEvalJob, err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("[GetEvalJobById] Error occurs! error=%v\n", err.Error())
		}
	}()
	var rsEvalJob dao.RsEvalJob
	tx := d.db.Read(ctx).First(&rsEvalJob, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &rsEvalJob, util.GormRealError(tx)
}