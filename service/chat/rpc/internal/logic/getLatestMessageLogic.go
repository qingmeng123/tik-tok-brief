package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/chat/model"
	"tik-tok-brief/service/chat/rpc/internal/svc"
	"tik-tok-brief/service/chat/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLatestMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLatestMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLatestMessageLogic {
	return &GetLatestMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取双方最新的一条消息
func (l *GetLatestMessageLogic) GetLatestMessage(in *pb.GetLatestMessageReq) (*pb.GetLatestMessageResp, error) {
	var err error
	message, err := l.svcCtx.ChatModel.FindOneByUsers(l.ctx, in.FromUserId, in.ToUserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("ChatModel.FindChatLimitList err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	//没查询到时会返回空指针,给其赋予空结构体,用于比较日期
	if err == sqlx.ErrNotFound {
		message = &model.Chat{}
	}

	rmessage, err := l.svcCtx.ChatModel.FindOneByUsers(l.ctx, in.ToUserId, in.FromUserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("ChatModel.FindChatLimitList err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		rmessage = &model.Chat{}
	}
	//若没有消息
	if message.CreateTime.Unix() == rmessage.CreateTime.Unix() {
		return &pb.GetLatestMessageResp{MsgType: 2}, nil
	}

	res := new(pb.Message)
	var msgType int64
	if message.CreateTime.Unix() > rmessage.CreateTime.Unix() {
		_ = copier.Copy(res, message)
		msgType = 1
	} else {
		_ = copier.Copy(res, rmessage)
		msgType = 0
	}
	return &pb.GetLatestMessageResp{Message: res, MsgType: msgType}, nil
}
