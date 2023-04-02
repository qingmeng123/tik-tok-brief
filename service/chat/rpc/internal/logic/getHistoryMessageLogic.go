package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

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
	chatList, err := l.svcCtx.ChatModel.FindChatList(l.ctx, in.FromUserId, in.ToUserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("ChatModel.FindChatList err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return &pb.GetHistoryMessageResp{}, nil
	}
	res := make([]*pb.Message, len(chatList))
	copier.Copy(&res, chatList)
	if err != nil {
		logx.Error("copier_copy err:", err)
		return nil, errorx.NewInternalErr()
	}

	for i, chat := range chatList {
		res[i].CreateTime = chat.CreateTime.Format("2006-01-02 15:04:05")
	}

	return &pb.GetHistoryMessageResp{MessageList: res}, nil
}
