package logic

import (
	"context"
	"io"
	"mime/multipart"
	"tik-tok-brief/common/errorx"
	filepb "tik-tok-brief/service/file/rpc/proto/pb"
	"tik-tok-brief/service/video/api/internal/svc"
	"tik-tok-brief/service/video/api/internal/types"
	videopb "tik-tok-brief/service/video/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	file   *multipart.FileHeader
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext, file *multipart.FileHeader) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		file:   file,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PublishReq) (resp *types.PublishResp, err error) {
	//获取user_id
	userId := l.ctx.Value("user_id")

	//读取文件
	file, err := l.file.Open()
	if err != nil {
		logx.Error("open file err:", err)
		return nil, errorx.NewParamErr(errorx.ERRFILEPARAM)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		logx.Error("io_read err:", err)
		return nil, errorx.NewParamErr(errorx.ERRFILEPARAM)
	}

	fileName := req.Title

	//调用上传文件服务
	stream, err := l.svcCtx.FileRPC.UploadVideoByLocal(l.ctx)
	if err != nil {
		logx.Error("fileRPC_UploadVideoByCos err:", err)
		return nil, err
	}

	//分片
	byteSlice := chunkByteSlice(bytes, 10)

	//发送文件名
	err = stream.Send(&filepb.UploadVideoByLocalReq{VideoName: fileName})
	if err != nil {
		logx.Error("steam send filename err:", err)
		return nil, err
	}

	//发送视频文件字节数据
	for _, bs := range byteSlice {
		err = stream.Send(&filepb.UploadVideoByLocalReq{Data: bs})
		if err != nil {
			logx.Error("steam send err:", err)
			return nil, err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		logx.Error("close and recv err:", err)
		return nil, errorx.NewInternalErr()
	}

	//调用video rpc保存视频信息进数据库
	_, err = l.svcCtx.VideoRPC.SaveVideo(l.ctx, &videopb.SaveVideoReq{
		Title:    req.Title,
		UserId:   userId.(int64),
		PlayUrl:  res.PlayUrl,
		CoverUrl: res.CoverUrl,
	})

	if err != nil {
		logx.Error("videoRPC_saveVideo err:", err)
		return nil, err
	}

	return &types.PublishResp{
		Status: types.Status{
			StatusCode: errorx.OK,
			StatusMsg:  errorx.SUCCESS,
		},
	}, nil
}

func chunkByteSlice(data []byte, n int) [][]byte {
	length := len(data)
	chunkSize := (length + n - 1) / n
	chunks := make([][]byte, n)
	for i := 0; i < n; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > length {
			end = length
		}
		chunks[i] = data[start:end]
	}
	return chunks
}

//
//func (l *PublishVideoLogic) PublishVideo(req *types.PublishReq) (resp *types.PublishResp, err error) {
//
//	//读取文件
//	file,err:=l.file.Open()
//	if err!=nil{
//		logx.Error("open file err:",err)
//		return nil, errorx.NewParamErr(errorx.ERRFILEPARAM)
//	}
//
//	bytes, err := io.ReadAll(file)
//	if err!=nil{
//		logx.Error("io_read err:",err)
//		return nil, errorx.NewParamErr(errorx.ERRFILEPARAM)
//	}
//
//	//压缩字节
//	zipBytes := tool.GZipBytes(bytes)
//
//	fileName:=req.Title+".mp4"
//
//	//调用上传文件服务
//	cosResp, err := l.svcCtx.FileRPC.UploadVideoByCos(l.ctx, &pb.UploadVideoByCosReq{
//		VideoName: fileName,
//		Data:      zipBytes,
//	})
//	if err!=nil{
//		logx.Error("fileRPC_UploadVideoByCos err:",err)
//		return nil, err
//	}
//
//	fmt.Println(cosResp)
//
//	return &types.PublishResp{
//		Status: types.Status{
//			StatusCode: errorx.OK,
//			StatusMsg:  errorx.SUCCESS,
//		},
//	}, nil
//}
