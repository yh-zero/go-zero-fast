package role

import (
	"context"
	"fmt"
	"go-zero-fast/service/sys/rpc/pb"

	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建角色
func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.RoleInfo) (resp *types.MessageRes, err error) {
	role, err := l.svcCtx.RoleRpc.CreateRole(l.ctx, &pb.RoleInfo{
		Status: req.Status,
		Name:   req.Name,
		Code:   req.Code,
		Remark: req.Remark,
		Sort:   req.Sort,
	})
	fmt.Println("-------------- role", role)
	if err != nil {
		return nil, err
	}

	return &types.MessageRes{
		Message: "添加角色成功！",
	}, nil
}
