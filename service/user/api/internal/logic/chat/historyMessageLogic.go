package chat

import (
	"context"
	"github.com/jinzhu/copier"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/chat/rpc/proto/pb"
	"time"

	"tik-tok-brief/service/user/api/internal/svc"
	"tik-tok-brief/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistoryMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHistoryMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryMessageLogic {
	return &HistoryMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HistoryMessageLogic) HistoryMessage(req *types.HistoryMessageReq) (resp *types.HistoryMessageResp, err error) {
	userId := l.ctx.Value("user_id").(int64)
	messagesResp, err := l.svcCtx.ChatRPC.GetHistoryMessage(l.ctx, &pb.GetHistoryMessageReq{
		FromUserId: userId,
		ToUserId:   req.ToUserId,
	})

	if err != nil {
		logx.Error("ChatRPC.SendMessage err:", err)
		return nil, err
	}
	messages := make([]types.Message, len(messagesResp.MessageList))
	_ = copier.Copy(&messages, messagesResp.MessageList)
	for i, message := range messagesResp.MessageList {
		t, _ := time.Parse("2006-01-02 15:04:05", message.CreateTime)
		messages[i].CreateTime = t.Unix()
	}
	return &types.HistoryMessageResp{
		StatusResponse: types.StatusResponse{
			StatusCode: int32(errorx.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		MessageList: messages,
	}, nil
}
