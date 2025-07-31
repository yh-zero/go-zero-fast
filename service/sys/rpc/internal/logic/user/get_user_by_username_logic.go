package userlogic

import (
	"context"
	"database/sql"
	"fmt"

	"go-zero-fast/common/fun"
	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据用户名获取用户详情
func (l *GetUserByUsernameLogic) GetUserByUsername(in *pb.UsernameReq) (*pb.UsernameRes, error) {

	userInfo, err := l.svcCtx.SysUsersModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 处理没有找到记录的情况
			return nil, errors.New("用户不存在")
		}
		// 处理其他错误
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	roleIds, err := l.svcCtx.UserRolesModel.FindRoleIdsByUserId(l.ctx, int64(userInfo.Id))
	if err != nil {
		return nil, errors.New("该用户没有相关的角色信息")
	}

	// 处理空角色情况
	var roleCodes []string
	if len(roleIds) > 0 {
		codes, err := l.svcCtx.SysRolesModel.FindCodesByIds(l.ctx, roleIds)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("FindCodesByIds error: %v", err)
			return nil, errors.New("没有找到角色代码")
		}
		roleCodes = codes
	}

	if err != nil {
		return nil, errors.New("没有找到角色代码")
	}

	createdAt := uint64(userInfo.CreatedAt.UnixMilli())
	updatedAt := uint64(userInfo.UpdatedAt.UnixMilli())

	return &pb.UsernameRes{
		UserInfo: &pb.UserInfo{
			Id:           &userInfo.Id,
			CreatedAt:    &createdAt,
			UpdatedAt:    &updatedAt,
			Status:       &userInfo.Status,
			Username:     &userInfo.Username,
			Password:     &userInfo.Password,
			Nickname:     &userInfo.Nickname,
			Description:  fun.NullStringToPtr(userInfo.Description),
			HomePath:     &userInfo.HomePath,
			RoleIds:      roleIds,
			Mobile:       fun.NullStringToPtr(userInfo.Mobile),
			Email:        &userInfo.HomePath,
			Avatar:       &userInfo.HomePath,
			DepartmentId: &userInfo.DepartmentId,
			RoleCodes:    roleCodes,
		},
	}, nil
}
