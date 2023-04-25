package user

import (
	"context"

	"github.com/houseme/goframe-polaris-demo/api/pbentity"
	"github.com/houseme/goframe-polaris-demo/internal/dao"
	"github.com/houseme/goframe-polaris-demo/internal/model/do"
	"github.com/houseme/goframe-polaris-demo/internal/service"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(&sUser{})
}

func (s *sUser) GetById(ctx context.Context, uid uint64) (*pbentity.User, error) {
	var user *pbentity.User
	err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Scan(&user)
	return user, err
}

func (s *sUser) DeleteById(ctx context.Context, uid uint64) error {
	_, err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Delete()
	return err
}
