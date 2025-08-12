package role

import (
	"context"

	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过ID获取角色
func NewGetRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByIdLogic {
	return &GetRoleByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleByIdLogic) GetRoleById(req *types.IDReq) (resp *types.RoleInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
