package dao

type User struct {
	ID          int64        // 数据集自增id
	Username    string       // 数据集名称
	Password    string       // 数据集备注
	Email       string       // 数据集保存路径
}