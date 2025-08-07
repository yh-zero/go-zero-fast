package menulogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-fast/common/fun"
	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"
)

type GetMenuListByRoleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByRoleIdLogic {
	return &GetMenuListByRoleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据角色id 获取菜单 -- 目前系统用这个 可以方便用户切换角色获取不一样的菜单
func (l *GetMenuListByRoleIdLogic) GetMenuListByRoleId(in *pb.IDReq) (*pb.MenuInfoList, error) {
	menuIds, err := l.svcCtx.RoleMenusModel.FindMenuIdsByRoleId(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GetMenuListByRoleId 获取用户菜单权限失败 error: %v", err)
		return nil, errors.New("获取用户菜单权限失败")
	}

	// 如果没有菜单权限，返回空列表而不是错误
	if len(menuIds) == 0 {
		return &pb.MenuInfoList{
			Total:    0,
			MenuInfo: []*pb.MenuInfo{},
		}, nil
	}

	menuList, err := l.svcCtx.SysMenusModel.FindMenusByIds(l.ctx, menuIds)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GetMenuListByRoleId 查询菜单列表失败 error: %v", err)
		return nil, errors.New("查询菜单列表失败")
	}

	resp := &pb.MenuInfoList{
		MenuInfo: make([]*pb.MenuInfo, 0, len(menuList)),
	}

	// 用于去重
	existMap := make(map[int64]struct{})

	for _, m := range menuList {
		// 跳过已存在的菜单项
		if _, ok := existMap[m.Id]; ok {
			continue
		}

		// 只返回启用的菜单
		if m.Disabled == 1 {
			continue
		}

		resp.MenuInfo = append(resp.MenuInfo, &pb.MenuInfo{

			Id:          uint64(m.Id),
			MenuType:    m.MenuType,
			Level:       m.MenuLevel,
			ParentId:    m.ParentId,
			Path:        m.Path,
			Name:        m.Name,
			Redirect:    m.Redirect,
			Component:   m.Component,
			Sort:        m.Sort,
			ServiceName: m.ServiceName,
			Permission:  fun.NullStringToString(m.Permission),
			CreatedAt:   uint64(m.CreatedAt.UnixMilli()),
			UpdatedAt:   uint64(m.UpdatedAt.UnixMilli()),
			Meta: &pb.Meta{
				Title:              m.Title,
				Icon:               m.Icon,
				HideMenu:           fun.IntToBool(m.HideMenu),
				HideBreadcrumb:     fun.IntToBool(m.HideBreadcrumb),
				IgnoreKeepAlive:    fun.IntToBool(m.IgnoreKeepAlive),
				HideTab:            fun.IntToBool(m.HideTab),
				FrameSrc:           m.FrameSrc,
				CarryParam:         fun.IntToBool(m.CarryParam),
				HideChildrenInMenu: fun.IntToBool(m.HideChildrenInMenu),
				Affix:              fun.IntToBool(m.Affix),
				DynamicLevel:       m.DynamicLevel,
				RealPath:           m.RealPath,
			},
		})

		existMap[m.Id] = struct{}{}
	}

	resp.Total = uint64(len(resp.MenuInfo))

	return resp, nil
}

//func (l *GetMenuListByRoleIdLogic) GetMenuListByRoleId(in *pb.IDReq) (*pb.MenuInfoList, error) {
//	menuIds, err := l.svcCtx.RoleMenusModel.FindMenuIdsByRoleId(l.ctx, in.Id)
//	if err != nil {
//		logx.WithContext(l.ctx).Errorf("GetMenuListByRoleId 改用户没有菜单权限 error: %v", err)
//		return nil, errors.New("改用户没有菜单权限")
//	}
//	menuList, err := l.svcCtx.SysMenusModel.FindMenusByIds(l.ctx, menuIds)
//	fmt.Println("========== menuList", menuList)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &pb.MenuInfoList{}, nil
//}
