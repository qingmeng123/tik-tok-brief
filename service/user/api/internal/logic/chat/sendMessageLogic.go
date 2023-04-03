package chat

import (
	"context"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/chat/rpc/proto/pb"

	"tik-tok-brief/service/user/api/internal/svc"
	"tik-tok-brief/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.SendMessageReq) (resp *types.SendMessageResp, err error) {
	userId := l.ctx.Value("user_id").(int64)
	_, err = l.svcCtx.ChatRPC.SendMessage(l.ctx, &pb.SendMessageReq{
		FromUserId: userId,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
	})

	if err != nil {
		logx.Error("ChatRPC.SendMessage err:", err)
		return nil, err
	}

	return &types.SendMessageResp{StatusResponse: types.StatusResponse{
		StatusCode: int32(errorx.OK),
		StatusMsg:  errorx.SUCCESS,
	}}, nil
}
