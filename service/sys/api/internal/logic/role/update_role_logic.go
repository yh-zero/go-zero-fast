package role

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-fast/common/fun"
	"go-zero-fast/common/result/xerr"
	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"
	"go-zero-fast/service/sys/rpc/pb"
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
	if req == nil || req.Id == 0 {
		return nil, errors.New("参数错误：角色ID不能为空")
	}
	fmt.Println("------------ UpdateRole req", req)
	fmt.Println("------------ req.Id", req.Id)

	updateReq := &pb.RoleUpdateRequest{}

	// 使用通用方法构建更新请求
	if err := fun.BuildUpdateRequest(updateReq, req); err != nil {
		return nil, err
	}
	fmt.Println("------------ updateReq.Id", updateReq.Id)

	_, err = l.svcCtx.RoleRpc.UpdateRole(l.ctx, updateReq)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("调用RPC更新角色失败, ID: %d, 错误: %v", req.Id, err)
		return nil, xerr.NewErrMsg(fmt.Sprintf("更新角色失败: %v", err))
	}

	return &types.MessageRes{Message: "更新成功！"}, nil
}
