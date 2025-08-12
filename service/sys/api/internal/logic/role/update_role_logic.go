package role

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go-zero-fast/common/result/xerr"
	"go-zero-fast/service/sys/rpc/pb"

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
	if req == nil || req.ID == 0 {
		return nil, errors.New("参数错误：角色ID不能为空")
	}
	if req.Name == "" {
		return nil, errors.New("参数错误：角色名称不能为空")
	}

	_, err = l.svcCtx.RoleRpc.UpdateRole(l.ctx, &pb.RoleInfo{
		Model: &pb.Model{
			Id: uint64(req.ID),
		},
		Status: req.Status,
		Name:   req.Name,
		Code:   req.Code,
		Remark: req.Remark,
		Sort:   req.Sort,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("调用RPC更新角色失败, ID: %d, 错误: %v", req.ID, err)
		return nil, xerr.NewErrMsg(fmt.Sprintf("更新角色失败: %v", err))
	}

	return &types.MessageRes{Message: "更新成功！"}, nil
}
