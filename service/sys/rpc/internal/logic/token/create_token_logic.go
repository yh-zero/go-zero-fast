package tokenlogic

import (
	"context"
	"fmt"
	"time"

	"go-zero-fast/common/snowflake"
	"go-zero-fast/service/sys/model"
	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成Token
func (l *CreateTokenLogic) CreateToken(in *pb.TokenInfoReq) (*pb.TokenInfoRes, error) {
	id, err := snowflake.GenID()
	fmt.Println("========= snowflake id ", id)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("生成id错误 error: %v", err)
		return nil, errors.New("没有找到角色代码")
	}

	fmt.Println("in.TokenInfo.ExpiredAt=", in.TokenInfo.ExpiredAt)
	insert, err := l.svcCtx.SysTokensModel.Insert(l.ctx, &model.SysTokens{
		Id:       id,
		Status:   in.TokenInfo.Status,
		UserId:   in.TokenInfo.UserId,
		Username: in.TokenInfo.Username,
		Token:    in.TokenInfo.Token,
		Source:   in.TokenInfo.Source,
		//ExpiredAt: time.UnixMilli(int64(in.TokenInfo.ExpiredAt)).UTC(),
		ExpiredAt: time.UnixMilli(int64(in.TokenInfo.ExpiredAt)).Local(),
	})
	fmt.Println("============ insert", insert)
	if err != nil {
		return nil, err
	}

	return &pb.TokenInfoRes{
		Id:    id,
		Token: "",
	}, nil
}
