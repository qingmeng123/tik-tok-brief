package logic

import (
	"context"

	"tik-tok-brief/service/file/rpc/internal/svc"
	"tik-tok-brief/service/file/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoByCosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVideoByCosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoByCosLogic {
	return &UploadVideoByCosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadVideoByCosLogic) UploadVideoByCos(in *pb.UploadVideoByCosReq) (*pb.UploadVideoByCosResp, error) {
	//

	return &pb.UploadVideoByCosResp{}, nil
}
