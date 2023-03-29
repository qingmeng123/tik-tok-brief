FROM golang AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct\
	GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

#将配置文件复制为执行docker-compose时能找到的路径
COPY service/video/api/etc /service/etc
#编译后放在service下
RUN go build -ldflags="-s -w" -o /service/videoApi service/video/api/video.go

#构建小镜像
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
#将二进制文件复制到镜像中的app目录下
COPY --from=builder /service/videoApi /app/videoApi
COPY --from=builder /service/etc /app/etc
CMD ["./videoApi", "-f", "etc/video-api.yaml"]
