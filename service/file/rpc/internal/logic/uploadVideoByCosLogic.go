package logic

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/common/tool"

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

	//解压
	zipBytes := tool.UGZipBytes(in.Data)
	key := "video/" + in.VideoName
	_, err := client.Object.Put(context.Background(), key, bytes.NewReader(zipBytes), nil)
	if err != nil {
		logx.Error("uploadVideoByCosLogic client_object_put err:", err)
		return nil, errorx.NewStatusParamErr(errorx.ERRFILEUPLOAD)
	}
	return &pb.UploadVideoByCosResp{
		PlayUrl: fmt.Sprintf("%s", client.Object.GetObjectURL(key)),
	}, nil
}
