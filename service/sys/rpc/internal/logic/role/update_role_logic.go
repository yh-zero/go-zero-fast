package rolelogic

import (
	"context"

	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新角色
func (l *UpdateRoleLogic) UpdateRole(in *pb.RoleInfo) (*pb.NoDataResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.NoDataResponse{}, nil
}
