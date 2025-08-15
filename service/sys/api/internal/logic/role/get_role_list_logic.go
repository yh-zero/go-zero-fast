package role

import (
	"context"
	"fmt"
	"go-zero-fast/common/fun"
	"go-zero-fast/service/sys/rpc/pb"
	"strconv"

	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取角色列表
func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleListLogic) GetRoleList(req *types.RoleListRequest) (resp *types.RoleListRes, err error) {
	fmt.Println("-------GetRoleList", req)
	roles, err := l.svcCtx.RoleRpc.GetRoleList(l.ctx, &pb.RoleListReq{
		PageInfo: &pb.PageInfo{
			PageNo:   req.PageNo,
			PageSize: req.PageSize,
			Cursor:   req.Cursor,
		},
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.RoleListRes{
		Total:  roles.Total,
		Cursor: roles.Cursor,
		List:   make([]types.RoleInfo, 0, len(roles.RoleInfo)), // 预分配容量
	}

	for _, v := range roles.RoleInfo {
		model := v.Model // 减少字段访问开销
		fmt.Println("-0----------------------strconv.FormatUint(model.CreatedAt, 10)", strconv.FormatUint(model.CreatedAt, 10))
		resp.List = append(resp.List, types.RoleInfo{
			Model: types.Model{
				Id:        model.Id,
				CreatedAt: fun.FormatTimestampToDate(model.CreatedAt),
				UpdatedAt: fun.FormatTimestampToDate(model.UpdatedAt),
				DeletedAt: fun.FormatTimestampToDate(model.DeletedAt),
			},
			Trans:         "",
			Status:        v.Status,
			Name:          v.Name,
			Code:          v.Code,
			Remark:        v.Remark,
			Sort:          v.Sort,
			DefaultRouter: v.DefaultRouter,
		})
	}
	fmt.Println("----------- 11resp", resp)
	return resp, nil
}
