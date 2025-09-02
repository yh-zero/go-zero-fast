package menulogic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go-zero-fast/common/fun"

	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取菜单列表
func (l *GetMenuListLogic) GetMenuList(in *pb.PageInfo) (*pb.MenuInfoList, error) {
	list, total, err := l.svcCtx.SysMenusModel.FindMenusList(l.ctx, in.PageNo, in.PageSize)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GetMenuList 获取菜单列表失败 error: %v", err)
		return nil, errors.New("获取菜单列表失败")
	}
	fmt.Println("GetMenuList total ", total)

	resp := &pb.MenuInfoList{}
	resp.Total = total

	for _, menu := range list {
		resp.MenuInfo = append(resp.MenuInfo, &pb.MenuInfo{
			Id:          uint64(menu.Id),
			Level:       menu.MenuLevel,
			ParentId:    menu.ParentId,
			Path:        menu.Path,
			Name:        menu.Name,
			Redirect:    menu.Redirect,
			Component:   menu.Component,
			Sort:        menu.Sort,
			Disabled:    fun.IntToBool(menu.Disabled),
			MenuType:    menu.MenuType,
			ServiceName: menu.ServiceName,
			Permission:  fun.NullStringToString(menu.Permission),
			Meta: &pb.Meta{
				Title:              menu.Title,
				Icon:               menu.Icon,
				HideMenu:           fun.IntToBool(menu.HideMenu),
				HideBreadcrumb:     fun.IntToBool(menu.HideBreadcrumb),
				IgnoreKeepAlive:    fun.IntToBool(menu.IgnoreKeepAlive),
				HideTab:            fun.IntToBool(menu.HideTab),
				FrameSrc:           menu.FrameSrc,
				CarryParam:         fun.IntToBool(menu.CarryParam),
				HideChildrenInMenu: fun.IntToBool(menu.HideChildrenInMenu),
				Affix:              fun.IntToBool(menu.Affix),
				DynamicLevel:       menu.DynamicLevel,
				RealPath:           menu.RealPath,
			},
			CreatedAt: uint64(menu.CreatedAt.UnixMilli()),
			UpdatedAt: uint64(menu.CreatedAt.UnixMilli()),
		})
	}

	return resp, nil
}
