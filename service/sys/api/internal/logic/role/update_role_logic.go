package role

import (
	"context"

	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新角色
func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.RoleInfo) (resp *types.MessageRes, err error) {
	// todo: add your logic here and delete this line

	return
}
