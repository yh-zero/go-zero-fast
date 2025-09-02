package role

import (
	"context"
	"go-zero-fast/common/result/xerr"
	"go-zero-fast/service/sys/rpc/pb"

	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除角色信息
func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoleLogic) DeleteRole(req *types.IDsReq) (resp *types.MessageRes, err error) {
	if len(req.Ids) == 0 {
		logx.Errorf("ids error: %v", err)
		return nil, xerr.NewErrMsg("ids 不能为空")
	}
	_, err = l.svcCtx.RoleRpc.DeleteRole(l.ctx, &pb.IDsRequest{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.MessageRes{Message: "删除成功！"}, nil
}
