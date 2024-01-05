package dao

type RsInferJobState string

// 这些是 RsInferJobState 的可能值
const (
	RsInferJobStateFailed  RsInferJobState = "Failed"
	RsInferJobStateRuning  RsInferJobState = "Runing"
	RsInferJobStateDeleted RsInferJobState = "Deleted"
	RsInferJobStatePending RsInferJobState = "Pending"
	RsInferJobStateFinished RsInferJobState = "Finished"
)

type RsInferJob struct {
	ID          int64        // 数据集自增id
	Name        string       // 数据集名称
	Path        string       // 数据集保存路径
	State       RsInferJobState // 数据集保存状态
	Modelid		int64
	Datasetid	int64
	CreateAt	int64
	UpdatedAt	int64
	DeletedAt	int64
	FileSize	int64
}
