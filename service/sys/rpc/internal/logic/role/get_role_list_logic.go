package rolelogic

import (
	"context"
	"fmt"
	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取角色列表
func (l *GetRoleListLogic) GetRoleList(in *pb.RoleListReq) (*pb.RoleListRes, error) {
	fmt.Println("-------------- GetRoleList -------------", in)
	cursor, nextCursor, total, err := l.svcCtx.SysRolesModel.FindPageByCursor(l.ctx, in.PageInfo.Cursor, in.PageInfo.PageSize)
	fmt.Println("=========== nextCursor", nextCursor)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("FindPageByCursor error: %v", err)
		return nil, err
	}

	resp := pb.RoleListRes{}
	resp.Cursor = nextCursor
	resp.Total = total
	for _, roles := range cursor {
		fmt.Println("----------------- roles", roles)
		resp.RoleInfo = append(resp.RoleInfo, &pb.RoleInfo{
			Status: roles.Status,
			Name:   roles.Name,
			Code:   roles.Code,
			Remark: roles.Remark,
			Sort:   roles.Sort,
			Model: &pb.Model{
				Id:        roles.Id,
				CreatedAt: uint64(roles.CreatedAt.UnixMilli()),
				UpdatedAt: uint64(roles.UpdatedAt.UnixMilli()),
			},
		})
	}

	fmt.Println("----------------- resp", resp)

	return &resp, nil
}
