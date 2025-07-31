package usernoauth

import (
	"context"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"time"

	"go-zero-fast/common/ctxJwt"
	"go-zero-fast/common/fun"
	"go-zero-fast/common/result/xerr"
	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 系统用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	driver := base64Captcha.NewDriverDigit(captchaImgHeight, captchaImgWidth, captchaImgLength, 0.1, 10)
	cp := base64Captcha.NewCaptcha(driver, store)
	ok := cp.Verify(req.CaptchaId, req.Captcha, true)
	if !ok {
		logx.Errorf("验证码获取失败! error: %v", err)
		return nil, xerr.NewErrCode(xerr.CAPTCHA_ERROR)
	}
	fmt.Println("======== ok ,", ok)

	userInfoRPC, err := l.svcCtx.UserRPC.GetUserByUsername(l.ctx, &pb.UsernameReq{Username: req.Username})
	fmt.Println("======== userInfoRPC", userInfoRPC)
	fmt.Println("======== err", err)
	if err != nil {
		return nil, err
	}

	if userInfoRPC.UserInfo.Status != nil && *userInfoRPC.UserInfo.Status != 1 {
		logx.Errorf("账号已冻结 error: %v", err)
		return nil, xerr.NewErrCode(xerr.ACCOUNT_FREEZE_ERROR)
	}

	err = fun.CheckPassword(req.Password, *userInfoRPC.UserInfo.Password)
	if err != nil {
		logx.Errorf("密码错误 error: %v", err)
		return nil, xerr.NewErrCode(xerr.USER_PASSWORD_ERROR)
	}

	// 获取token
	now := time.Now().Unix()
	accessToken, err := ctxJwt.GetJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, uint64(now), uint64(l.svcCtx.Config.JwtAuth.AccessExpire), ctxJwt.JWTData{
		RoleIds: userInfoRPC.UserInfo.RoleIds,
		//RoleId:       getFirstRoleId(userInfoRPC.UserInfo.RoleIds),
		RoleId:       userInfoRPC.UserInfo.RoleIds[0],
		DepartmentId: *userInfoRPC.UserInfo.DepartmentId,
		UserId:       *userInfoRPC.UserInfo.Id,
		//Username: *userInfoRPC.UserInfo.Username,
		//ID:       *userInfoRPC.UserInfo.Id,
		//NickName: *userInfoRPC.UserInfo.Nickname,
	})
	fmt.Println("============ accessToken", accessToken)
	if err != nil {
		logx.Errorf("获取token失败 error: %v", err)
		return nil, err
	}

	// 把token 存到数据库 用于记录
	expiredAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire)).UnixMilli()
	_, err = l.svcCtx.TokenRPC.CreateToken(l.ctx, &pb.TokenInfoReq{
		TokenInfo: &pb.TokenInfo{
			Status:    userInfoRPC.UserInfo.Status, // 暂时用 用户的状态 输入
			UserId:    userInfoRPC.UserInfo.Id,
			Username:  userInfoRPC.UserInfo.Username,
			Token:     &accessToken,
			Source:    fun.GetStringLocal("local"),
			ExpiredAt: fun.Int64ToUint64Ptr(expiredAt),
		},
	})
	if err != nil {
		return nil, err
	}

	// 这里需要做一个 删除验证码的 redis 缓存的操作

	return &types.LoginRes{
		LoginInfo: types.LoginInfo{
			UserId: *userInfoRPC.UserInfo.Id,
			Token:  accessToken,
			Expire: uint64(expiredAt),
		},
	}, nil
}

func getFirstRoleId(roleIds []uint64) uint64 {
	if len(roleIds) > 0 {
		return roleIds[0]
	}
	return 0
}
