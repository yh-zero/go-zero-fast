package menu

import (
	"context"
	"fmt"
	"go-zero-fast/common/fun"
	"go-zero-fast/service/sys/rpc/pb"

	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取菜单列表
func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListLogic) GetMenuList(req *types.PageInfo) (resp *types.MenuInfoListResponse, err error) {
	menuListRpc, err := l.svcCtx.MenuRPC.GetMenuList(l.ctx, &pb.PageInfo{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("1111111111111111111")
	resp = &types.MenuInfoListResponse{} // 不初始化会报错
	resp.Total = menuListRpc.Total
	fmt.Println("2222222222222222222222")

	for _, v := range menuListRpc.MenuInfo {
		resp.List = append(resp.List, types.MenuInfo{
			//Trans:       "",
			Level:       v.Level,
			ParentId:    v.ParentId,
			Path:        v.Path,
			Name:        v.Name,
			Redirect:    v.Redirect,
			Component:   v.Component,
			Sort:        v.Sort,
			Disabled:    v.Disabled,
			MenuType:    v.MenuType,
			ServiceName: v.ServiceName,
			Permission:  v.Permission,
			Model: types.Model{
				Id:        v.Id,
				CreatedAt: fun.FormatTimestampToDate(v.CreatedAt),
				UpdatedAt: fun.FormatTimestampToDate(v.UpdatedAt),
			},
			Meta: types.Meta{
				Title:              v.Meta.Title,
				Icon:               v.Meta.Icon,
				HideMenu:           v.Meta.HideMenu,
				HideBreadcrumb:     v.Meta.HideBreadcrumb,
				IgnoreKeepAlive:    v.Meta.IgnoreKeepAlive,
				HideTab:            v.Meta.HideTab,
				FrameSrc:           v.Meta.FrameSrc,
				CarryParam:         v.Meta.CarryParam,
				HideChildrenInMenu: v.Meta.HideChildrenInMenu,
				Affix:              v.Meta.Affix,
				DynamicLevel:       v.Meta.DynamicLevel,
				RealPath:           v.Meta.RealPath,
			},
		})

	}

	return
}
