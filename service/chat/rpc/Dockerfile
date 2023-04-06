FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct\
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update --no-cache && apk add --no-cache tzdata
WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

#将配置文件复制到镜像中
COPY service/chat/rpc/etc /service/etc
#编译后放在service下
RUN go build -ldflags="-s -w" -o /service/chatRpc service/chat/rpc/chat.go

#构建小镜像
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
#将二进制文件复制到镜像中的app目录下
COPY --from=builder /service/chatRpc /app/chatRpc
COPY --from=builder /service/etc /app/etc
CMD ["./chatRpc", "-f", "/app/etc/chat.yaml"]
