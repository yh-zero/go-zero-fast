package rolelogic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go-zero-fast/service/sys/model"
	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建新角色
func (l *CreateRoleLogic) CreateRole(in *pb.RoleInfo) (*pb.NoDataResponse, error) {
	insert, err := l.svcCtx.SysRolesModel.Insert(l.ctx, &model.SysRoles{
		//Id:            0,  // 先用自增 后续用雪花算法
		Status:        in.Status,
		Name:          in.Name,
		Code:          in.Code,
		DefaultRouter: "dashboard",
		Remark:        in.Remark,
		Sort:          in.Sort,
	})
	fmt.Println("---------- insert", insert)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("CreateRole error: %v", err)
		return nil, errors.New("角色可能已存在")
	}
	return &pb.NoDataResponse{}, nil
}
