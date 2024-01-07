package db

import (
	"RS-Backend/models/dao"
	"RS-Backend/models/dto"
	"RS-Backend/util"
	"context"

	"github.com/sirupsen/logrus"

)

type IUserAccesser interface {
	Register(ctx context.Context, user *dao.User) (err error)
	LogIn(ctx context.Context, email string, password string) (result *dto.User, err error)
}

type UserAccesserImpl struct {
	db IDB
}

func NewUserAccesser(db IDB) IUserAccesser {
	return &UserAccesserImpl{db: db}
}

func (d *UserAccesserImpl) Register(ctx context.Context, user *dao.User) (err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("[Register] Error occurs! error=%v\n", err.Error())
		}
	}()
	tx := d.db.Write(ctx).Create(user)
	return util.GormRealError(tx)
}
func (d *UserAccesserImpl) LogIn(ctx context.Context,  email string, password string) (result *dto.User, err error) {
	defer func() {
		if err != nil {
			logrus.Errorf("[LogIn] Error occurs! error=%v\n", err.Error())
		}
	}()
	var user dto.User
	tx := d.db.Read(ctx).Where("email = ? AND password = ?", email, password).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, util.GormRealError(tx)
}