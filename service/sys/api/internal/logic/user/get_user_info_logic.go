package user

import (
	"context"
	"fmt"
	"go-zero-fast/common/ctxJwt"
	"go-zero-fast/service/sys/rpc/pb"

	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据token获取用户信息
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserBaseIDInfoRes, err error) {
	jwtUserId := ctxJwt.GetJwtDataUserId(l.ctx)
	userInfoRpc, err := l.svcCtx.UserRPC.GetUserInfoById(l.ctx, &pb.IDReq{Id: jwtUserId})
	fmt.Println("========= userInfoRpc", userInfoRpc)
	if err != nil {
		return nil, err
	}
	return &types.UserBaseIDInfoRes{
		UserInfo: types.UserInfo{
			Id:             *userInfoRpc.UserInfo.Id,
			Username:       *userInfoRpc.UserInfo.Username,
			Nickname:       *userInfoRpc.UserInfo.Nickname,
			Avatar:         *userInfoRpc.UserInfo.Avatar,
			HomePath:       *userInfoRpc.UserInfo.HomePath,
			Description:    *userInfoRpc.UserInfo.Description,
			RoleName:       userInfoRpc.UserInfo.RoleName,
			DepartmentName: *userInfoRpc.UserInfo.DepartmentName,
		},
	}, nil
}
