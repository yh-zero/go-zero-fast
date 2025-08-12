package rolelogic

import (
	"context"

	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByIdLogic {
	return &GetRoleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过ID获取角色
func (l *GetRoleByIdLogic) GetRoleById(in *pb.IDReq) (*pb.NoDataResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.NoDataResponse{}, nil
}
