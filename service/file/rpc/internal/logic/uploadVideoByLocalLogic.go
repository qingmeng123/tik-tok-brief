package logic

import (
	"bytes"
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"io"
	"os"
	"path/filepath"
	"strings"
	"tik-tok-brief/common/errorx"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/zeromicro/go-zero/core/logx"
	"tik-tok-brief/service/file/rpc/internal/svc"
	"tik-tok-brief/service/file/rpc/proto/pb"
)

type UploadVideoByLocalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVideoByLocalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoByLocalLogic {
	return &UploadVideoByLocalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传视频到本地
func (l *UploadVideoByLocalLogic) UploadVideoByLocal(stream pb.File_UploadVideoByLocalServer) error {
	var buffer bytes.Buffer

	//获取客户端传来的视频名
	res, err := stream.Recv()
	if err != nil {
		logx.Error("stream_recv err:", err)
		return errorx.NewStatusParamErr(errorx.ERRFILEUPLOAD)
	}
	//保存videoName
	videoName := res.VideoName + l.svcCtx.Config.Cos.VideoSuffix

	//接受客户端的流式数据
	for {
		res, err = stream.Recv()
		if err == io.EOF {
			//完成读取
			videoPath := l.svcCtx.Config.LocalVideoPath + videoName
			if err = os.MkdirAll(filepath.Dir(videoPath), os.ModePerm); err != nil {
				logx.Error("write file err:", err)
			}
			err = os.WriteFile(videoPath, buffer.Bytes(), 0666)
			if err != nil {
				logx.Error("write file err:", err)
				return errorx.NewStatusParamErr(errorx.ERRFILEUPLOAD)
			}
			coverPath := l.svcCtx.Config.LocalCoverPath + videoName
			coverPath, err = GetSnapshot(videoPath, coverPath, 1)
			if err != nil {
				logx.Error("getSnapshot err:", err)
				return errorx.NewStatusParamErr(errorx.ERRFILEUPLOAD)
			}
			return stream.SendAndClose(&pb.UploadVideoByLocalResp{
				PlayUrl:  l.svcCtx.Config.StaticFileServiceIP + videoPath,
				CoverUrl: l.svcCtx.Config.StaticFileServiceIP + coverPath,
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

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return "", err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		return "", err
	}
	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		return "", err
	}
	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}
