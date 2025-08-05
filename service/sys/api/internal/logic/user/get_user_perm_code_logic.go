package user

import (
	"context"
	"fmt"
	"go-zero-fast/common/ctxJwt"
	"go-zero-fast/common/result/xerr"
	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPermCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户权限码
func NewGetUserPermCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPermCodeLogic {
	return &GetUserPermCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPermCodeLogic) GetUserPermCode() (resp *types.PermCodeRes, err error) {
	roleIds := ctxJwt.GetJwtDataRoleIds(l.ctx)
	if len(roleIds) == 0 {
		logx.Errorf("获取权限码错误 error: %v", err)
		return nil, xerr.NewErrCode(xerr.USER_PASSWORD_ERROR)
	}
	fmt.Println("roleIds", roleIds)

	return &types.PermCodeRes{
		Data: roleIds,
	}, nil
}
