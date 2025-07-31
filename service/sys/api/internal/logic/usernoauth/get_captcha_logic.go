package usernoauth

import (
	"context"

	"go-zero-fast/common/result/xerr"
	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

var store = base64Captcha.DefaultMemStore

const (
	prefixCaptcha    = "biz#captcha#ip:%s"
	expireCaptcha    = 60 * 3000 // 2分钟
	captchaImgWidth  = 105
	captchaImgHeight = 36
	captchaImgLength = 6
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取验证码
func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.CaptchaRes, err error) {
	driver := base64Captcha.NewDriverDigit(captchaImgHeight, captchaImgWidth, captchaImgLength, 0.1, 10)
	cp := base64Captcha.NewCaptcha(driver, store)

	id, b64s, _, err := cp.Generate()
	//id, b64s, answer, err := cp.Generate()
	if err != nil {
		logx.Errorf("验证码生成失败! error: %v", err)
		return nil, xerr.NewErrCode(xerr.CAPTCHA_GENERATE_ERROR)

	}

	return &types.CaptchaRes{
		CaptchaInfo: types.CaptchaInfo{
			CaptchaId: id,
			ImgPath:   b64s,
		},
	}, nil
}
