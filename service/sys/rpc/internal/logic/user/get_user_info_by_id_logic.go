package userlogic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go-zero-fast/common/fun"
	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIdLogic {
	return &GetUserInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户详细信息
func (l *GetUserInfoByIdLogic) GetUserInfoById(in *pb.IDRequest) (*pb.UserInfoResponse, error) {
	userInfo, err := l.svcCtx.SysUsersModel.FindOne(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GetUserInfoById 没有相关用户信息 error: %v", err)
		return nil, errors.New("没有相关用户信息")
	}
	roleIds, err := l.svcCtx.UserRolesModel.FindRoleIdsByUserId(l.ctx, userInfo.Id)
	fmt.Println("============== roleIds", roleIds)
	if err != nil {
		return nil, errors.New("该用户没有相关的角色信息")
	}

	var roleNames []string
	if len(roleIds) > 0 {
		roleNames, err = l.svcCtx.SysRolesModel.FindRoleNamesByIds(l.ctx, roleIds)
		fmt.Println("=========== roleNames", roleNames)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("FindCodesByIds error: %v", err)
			return nil, errors.New("没有找到角色代码")
		}
	}

	departments, err := l.svcCtx.SysDepartmentsModel.FindOne(l.ctx, userInfo.DepartmentId)
	fmt.Println("============ departments", departments)
	fmt.Println("============ departments", departments.Name)
	if err != nil {
		return nil, err
	}

	return &pb.UserInfoResponse{
		UserInfo: &pb.UserInfo{
			Id:             userInfo.Id,
			CreatedAt:      uint64(userInfo.CreatedAt.UnixMilli()),
			UpdatedAt:      uint64(userInfo.UpdatedAt.UnixMilli()),
			Status:         userInfo.Status,
			Username:       userInfo.Username,
			Password:       userInfo.Password,
			Nickname:       userInfo.Nickname,
			Description:    fun.NullStringToString(userInfo.Description),
			HomePath:       userInfo.HomePath,
			Mobile:         fun.NullStringToString(userInfo.Mobile),
			Email:          fun.NullStringToString(userInfo.Email),
			Avatar:         fun.NullStringToString(userInfo.Avatar),
			DepartmentId:   userInfo.DepartmentId,
			RoleName:       roleNames,
			RoleIds:        roleIds,
			DepartmentName: departments.Name, // 部门名称
			//PositionIds:   userInfo.PositionIds,
		},
	}, nil

	return nil, nil
}
