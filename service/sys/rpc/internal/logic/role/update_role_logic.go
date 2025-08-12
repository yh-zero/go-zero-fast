package rolelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-fast/service/sys/model"
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
func (l *UpdateRoleLogic) UpdateRole(in *pb.RoleInfo) (*pb.NoDataResponse, error) {
	logger := logx.WithContext(l.ctx)
	logger.Infof("开始更新角色, ID: %d", in.Model.Id)

	// 检查角色是否存在
	_, err := l.svcCtx.SysRolesModel.FindOne(l.ctx, in.Model.Id)
	if err != nil {
		logger.Errorf("查询角色失败, ID: %d, 错误: %v", in.Model.Id, err)
		return nil, errors.New("角色不存在")
	}

	// 准备更新数据
	updateData := &model.SysRoles{
		Id:     in.Model.Id,
		Status: in.Status,
		Name:   in.Name,
		Code:   in.Code,
		Remark: in.Remark,
		Sort:   in.Sort,
	}

	// 执行更新
	err = l.svcCtx.SysRolesModel.Update(l.ctx, updateData)
	if err != nil {
		logger.Errorf("更新角色失败, ID: %d, 错误: %v", in.Model.Id, err)
		return nil, errors.New("更新角色失败")
	}

	logger.Infof("角色更新成功, ID: %d", in.Model.Id)
	return &pb.NoDataResponse{}, nil
}
