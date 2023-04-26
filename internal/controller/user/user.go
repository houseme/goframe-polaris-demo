package user

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/houseme/goframe-polaris-demo/api/user/v1"
	"github.com/houseme/goframe-polaris-demo/internal/dao"
	"github.com/houseme/goframe-polaris-demo/internal/model/do"
	"github.com/houseme/goframe-polaris-demo/internal/service"
)

// Controller user controller
type Controller struct {
	v1.UnimplementedUserServer
}

// Register register user controller
func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

// Create user info
func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	var lastInsertID int64
	if lastInsertID, err = dao.User.Ctx(ctx).Data(do.User{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	}).InsertAndGetId(); err != nil {
		return nil, err
	}
	g.Log().Debug(ctx, "lastInsertId:", lastInsertID)
	res = &v1.CreateRes{
		Id: uint64(lastInsertID),
	}
	return
}

// GetOne user info
func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	user, err := service.User().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetOneRes{
		User: user,
	}
	g.Log().Debug(ctx, "user:", user)
	return
}

// GetList user info
func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	err = dao.User.Ctx(ctx).Page(int(req.Page), int(req.Size)).Scan(&res.Users)
	return
}

// Delete user info
func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.User().DeleteById(ctx, req.Id)
	return
}
