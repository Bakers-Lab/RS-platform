package dao

type DatasetBatch struct {
	ID          int64        // 数据集自增id
	Name        string       // 数据集名称
	DatasetId   int64       // 数据集备注
	State       string       // 数据集保存路径
	samples_num int64     // 数据集保存格式
	CreateAt	int64
	UpdatedAt	int64
	DeletedAt	int64
	FileSize	int64
}