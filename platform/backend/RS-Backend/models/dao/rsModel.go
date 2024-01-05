package dao

type RsModel struct {
	ID          int64        // 数据集自增id
	Name        string       // 数据集名称
	Taskname    string       // 数据集备注
	Comment     string       // 数据集保存路径
	RequestUrl  string       // 数据集保存格式
	Params      string // 数据集保存状态
}