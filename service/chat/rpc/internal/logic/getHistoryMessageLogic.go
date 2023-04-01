package logic

import (
	"context"

	"tik-tok-brief/service/chat/rpc/internal/svc"
	"tik-tok-brief/service/chat/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHistoryMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryMessageLogic {
	return &GetHistoryMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取历史消息
func (l *GetHistoryMessageLogic) GetHistoryMessage(in *pb.GetHistoryMessageReq) (*pb.GetHistoryMessageResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetHistoryMessageResp{}, nil
}
