package dto

import (
	"RS-Backend/models/dao"
)

type RsInferJob struct {
	ID          int64        // 数据集自增id
	Name        string       // 数据集名称
	Path        string       // 数据集保存路径
	State       dao.RsInferJobState // 数据集保存状态
	Modelid		int64
	ModelName	string
	Datasetid	int64
	DatasetName	string
	CreateAt	int64
	UpdatedAt	int64
	DeletedAt	int64
	FileSize	int64
}