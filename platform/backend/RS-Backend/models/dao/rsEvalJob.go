package dao

type RsEvalJobState string

// 这些是 RsEvalJobState 的可能值
const (
	RsEvalJobStateFailed  RsEvalJobState = "Failed"
	RsEvalJobStateRuning  RsEvalJobState = "Runing"
	RsEvalJobStateDeleted RsEvalJobState = "Deleted"
	RsEvalJobStatePending RsEvalJobState = "Pending"
	RsEvalJobStateFinished RsEvalJobState = "Finished"
)

type RsEvalJob struct {
	ID          int64        // 数据集自增id
	Name        string       // 数据集名称
	Path        string       // 数据集保存路径
	State       RsEvalJobState // 数据集保存状态
	Inferjobid	int64
	Datasetid	int64
	CreateAt	int64
	UpdatedAt	int64
	DeletedAt	int64
	FileSize	int64
}
