package rolelogic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"

	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除角色信息
func (l *DeleteRoleLogic) DeleteRole(in *pb.IDsReq) (*pb.NoDataResponse, error) {
	// 先查找角色id有没有相关的用户信息  有就不能删除
	userIds, err := l.svcCtx.UserRolesModel.FindUserIdsByRoleIds(l.ctx, in.Ids)
	fmt.Println("-------------- userIds", userIds)
	if err != nil {
		return nil, err
	}
	count, err := l.svcCtx.SysUsersModel.CountUsersByIds(l.ctx, userIds)
	fmt.Println("-------------- count", count)
	if err != nil {
		return nil, err
	}
	if count != 0 {
		logx.WithContext(l.ctx).Errorf("改角色下有相关的用户，不能删除 DeleteRole error: %v", err)
		return nil, errors.New("改角色下有相关的用户，不能删除")
	}
	err = l.svcCtx.SysRolesModel.DeleteByIds(l.ctx, in.Ids)
	if err != nil {
		return nil, err
	}

	return &pb.NoDataResponse{}, nil
}
