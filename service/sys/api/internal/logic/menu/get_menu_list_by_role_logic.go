package menu

import (
	"context"
	"fmt"
	"go-zero-fast/common/ctxJwt"
	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取菜单列表 -- 对应的用户角色的权限
func NewGetMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByRoleLogic {
	return &GetMenuListByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListByRoleLogic) GetMenuListByRole() (resp *types.MenuListRes, err error) {
	jwtRoleId := ctxJwt.GetJwtDataRoleId(l.ctx)
	menuList, err := l.svcCtx.MenuRPC.GetMenuListByRoleId(l.ctx, &pb.IDReq{Id: jwtRoleId})
	fmt.Println("====== menuList", menuList)
	if err != nil {
		return nil, err
	}

	resp = &types.MenuListRes{}
	fmt.Println("---------------- 1111")

	for _, v := range menuList.MenuInfo {
		resp.List = append(resp.List, types.MenuInfo{
			Model:       types.Model{Id: v.Id},
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
	fmt.Println("---------------- 22222")

	return resp, nil
}
