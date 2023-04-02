package logic

import (
	"context"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/chat/model"
	"tik-tok-brief/service/chat/rpc/internal/svc"
	"tik-tok-brief/service/chat/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送消息
func (l *SendMessageLogic) SendMessage(in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	_, err := l.svcCtx.ChatModel.Insert(l.ctx, &model.Chat{
		FromUserId: in.FromUserId,
		ToUserId:   in.ToUserId,
		Content:    in.Content,
	})
	if err != nil {
		logx.Error("ChatModel.Insert err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	return &pb.SendMessageResp{}, nil
}
