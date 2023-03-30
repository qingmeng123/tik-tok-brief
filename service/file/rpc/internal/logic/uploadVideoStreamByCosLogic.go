package logic

import (
	"bytes"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"os"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/file/rpc/internal/svc"
	"tik-tok-brief/service/file/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoStreamByCosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVideoStreamByCosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoStreamByCosLogic {
	return &UploadVideoStreamByCosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadVideoStreamByCosLogic) UploadVideoStreamByCos(stream pb.File_UploadVideoStreamByCosServer) error {
	//解析存储桶地址
	u, _ := url.Parse(l.svcCtx.Config.Cos.BucketURL)
	b := &cos.BaseURL{BucketURL: u}

	//创建连接COS客户端
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{

			SecretID:  os.Getenv(l.svcCtx.Config.Cos.SecretID),
			SecretKey: os.Getenv(l.svcCtx.Config.Cos.SecretKey),
		},
	})

	var buffer bytes.Buffer

	//获取客户端传来的视频名
	res, err := stream.Recv()
	if err != nil {
		logx.Error("stream_recv err:", err)
		return errorx.NewStatusParamErr(errorx.ERRFILEUPLOAD)
	}
	//保存videoName
	videoName := res.VideoName

	//接受客户端的流式数据
	for {
		res, err = stream.Recv()
		if err == io.EOF {
			//完成读取
			playPath := l.svcCtx.Config.Cos.VideoPrefix + videoName + l.svcCtx.Config.Cos.VideoSuffix
			coverPath := l.svcCtx.Config.Cos.CoverPrefix + videoName + l.svcCtx.Config.Cos.CoverSuffix
			_, err = client.Object.Put(context.Background(), playPath, bytes.NewReader(buffer.Bytes()), nil)
			if err != nil {
				logx.Error("uploadVideoByCosLogic client_object_put err:", err)
				return errorx.NewStatusParamErr(errorx.ERRFILEUPLOAD)
			}

			return stream.SendAndClose(&pb.UploadVideoByCosResp{
				PlayUrl:  l.svcCtx.Config.Cos.BucketURL + l.svcCtx.Config.Cos.VideoPrefix + videoName + l.svcCtx.Config.Cos.VideoTranscodingSuffix,
				CoverUrl: l.svcCtx.Config.Cos.BucketURL + coverPath,
			})
		}

		if err != nil {
			logx.Error("stream recv err:", err)
			return errorx.NewStatusParamErr(errorx.ERRFILEUPLOAD)
		}
		//增加数据
		buffer.Write(res.Data)
	}

}
