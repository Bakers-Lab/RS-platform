package dao

type DatasetState string

// 这些是 DatasetState 的可能值
const (
	StateFailed  DatasetState = "Failed"
	StateReady   DatasetState = "Ready"
	StateDeleted DatasetState = "Deleted"
)

type Dataset struct {
	ID          int64        // 数据集自增id
	Name        string       // 数据集名称
	Comment     string       // 数据集备注
	Path        string       // 数据集保存路径
	StoreFormat string       // 数据集保存格式
	State       DatasetState // 数据集保存状态
}
