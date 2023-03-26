package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"tik-tok-brief/common/errorx"
	userpb "tik-tok-brief/service/user/rpc/proto/pb"
	"tik-tok-brief/service/video/api/internal/svc"
	"tik-tok-brief/service/video/api/internal/types"
	videopb "tik-tok-brief/service/video/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListReq) (resp *types.PublishListResp, err error) {
	//token中保存的id
	tUserId := l.ctx.Value("user_id").(int64)

	//获取请求用户信息
	getUserResp, err := l.svcCtx.UserRPC.GetUser(l.ctx, &userpb.GetUserReq{UserID: req.UserId})
	if err != nil {
		logx.Error("userRPC_getUser err:", err)
		return nil, err
	}

	userinfo := getUserResp.User

	//是否关注
	if tUserId != req.UserId {
		userinfo.IsFollow = true
	}

	//获取用户视频
	publishListResp, err := l.svcCtx.VideoRPC.PublishList(l.ctx, &videopb.PublishListReq{UserId: req.UserId})
	if err != nil {
		logx.Error("videoRPC_publishList err:", err)
		return nil, err
	}

	//封装返回信息
	if err != nil {
		logx.Error("copier copy err:", err)
		return nil, errorx.NewInternalErr()
	}

	resp = new(types.PublishListResp)
	resp.VideoList = make([]types.Video, len(publishListResp.VideoList))

	copier.Copy(&resp.VideoList, publishListResp.VideoList)

	for i, _ := range resp.VideoList {
		copier.Copy(&resp.VideoList[i].Author, userinfo)
	}

	resp.StatusCode = errorx.OK
	resp.StatusMsg = errorx.SUCCESS

	return resp, nil
}
