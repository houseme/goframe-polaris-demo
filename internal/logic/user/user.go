package user

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/houseme/goframe-polaris-demo/api/pbentity"
	"github.com/houseme/goframe-polaris-demo/internal/dao"
	"github.com/houseme/goframe-polaris-demo/internal/model/do"
	"github.com/houseme/goframe-polaris-demo/internal/model/entity"
	"github.com/houseme/goframe-polaris-demo/internal/service"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(&sUser{})
}

// GetByID get user by id
func (s *sUser) GetByID(ctx context.Context, uid uint64) (*pbentity.User, error) {
	var (
		user       *pbentity.User
		userEntity *entity.User
	)
	if err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Scan(&userEntity); err != nil {
		return nil, err
	}
	user = &pbentity.User{
		Id:       uint32(userEntity.Id),
		Passport: userEntity.Passport,
		Nickname: userEntity.Nickname,
		Password: userEntity.Password,
		CreateAt: timestamppb.New(time.Unix(userEntity.CreateAt.Unix(), userEntity.CreateAt.UnixNano()%10000000*100)),
		UpdateAt: timestamppb.New(time.Unix(userEntity.UpdateAt.Unix(), userEntity.UpdateAt.UnixNano()%10000000*100)),
	}
	if userEntity.DeleteAt != nil {
		user.DeleteAt = timestamppb.New(time.Unix(userEntity.DeleteAt.Unix(), userEntity.DeleteAt.UnixNano()%10000000*100))
	}
	g.Log().Debug(ctx, "user:", user)
	g.Log().Debug(ctx, "userEntity:", userEntity)
	g.Log().Debug(ctx, "userEntity.CreateAt Unix:", userEntity.CreateAt.Unix(), " UnixNano:", userEntity.CreateAt.UnixNano(), " nsec", userEntity.CreateAt.UnixNano()%10000000*100)
	return user, nil
}

// DeleteByID delete user by id
func (s *sUser) DeleteByID(ctx context.Context, uid uint64) error {
	_, err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Delete()
	return err
}
