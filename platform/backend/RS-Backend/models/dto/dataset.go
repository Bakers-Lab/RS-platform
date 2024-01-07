package dto

import (
	"RS-Backend/models/dao"
)

type Dataset struct {
	ID          int64        // 数据集自增id
	Name        string       // 数据集名称
	Comment     string       // 数据集备注
	Path        string       // 数据集保存路径
	CreaterUserId int64
	StoreFormat string       // 数据集保存格式
	State       dao.DatasetState // 数据集保存状态
	Batches		[]dao.DatasetBatch
}
