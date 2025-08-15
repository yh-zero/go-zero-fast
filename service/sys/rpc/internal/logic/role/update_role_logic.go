package rolelogic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-fast/common/fun"
	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"
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
func (l *UpdateRoleLogic) UpdateRole(in *pb.RoleUpdateRequest) (*pb.NoDataResponse, error) {
	fmt.Println("------------- UpdateRole in", in)
	logger := logx.WithContext(l.ctx)
	logger.Infof("开始更新角色, ID: %d", in.Id)

	// 检查角色是否存在
	_, err := l.svcCtx.SysRolesModel.FindOne(l.ctx, in.Id)
	if err != nil {
		logger.Errorf("查询角色失败, ID: %d, 错误: %v", in.Id, err)
		return nil, errors.New("角色不存在")
	}

	updateData := fun.UpdateFieldsByReflect(in)

	// 动态生成SQL
	if len(updateData) == 0 {
		return nil, errors.New("无有效更新字段")
	}

	// 执行更新
	err = l.svcCtx.SysRolesModel.UpdateWithMap(l.ctx, in.Id, updateData)
	if err != nil {
		logger.Errorf("更新角色失败, ID: %d, 错误: %v", in.Id, err)
		return nil, errors.New("更新角色失败")
	}

	logger.Infof("角色更新成功, ID: %d", in.Id)
	return &pb.NoDataResponse{}, nil
}
