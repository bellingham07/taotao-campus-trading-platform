FROM centos:7

# 安装必要的组件
RUN yum install -y epel-release && \
yum update -y && \
yum install -y git wget curl which vim && \
yum clean all

# 下载golang，先设定版本，然后应用，加上官网的SHA256的校验和
# ENV GOLANG_VERSION 1.20.5
ENV GOLANG_DOWNLOAD_URL https://go.dev/dl/go1.20.5.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 d7ec48cde0d3d2be2c69203bc3e0a44de8660b9c09a6e85c4732a3f7dc442612


# 安装 Golang 运行环境
RUN set -eux
RUN wget $GOLANG_DOWNLOAD_URL -o golang.tar.gz
RUN tar -zxvf golang.tar.gz
RUN rm golang.tar.gz
RUN export PATH="/usr/local/go/bin:$PATH"
RUN go version

# 安装 protobuf 工具
RUN set -ex && yum install -y protobuf-compiler

# 安装 go-zero 框架需要的 goctl 工具
RUN go install github.com/zeromicro/go-zero/tools/goctl@latest

# 设置环境变量
ENV GOROOT /usr/local/go
ENV PATH "$GOROOT/bin:$PATH"
ENV PATH "$PATH:$GOPATH/bin"
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

  # 创建工作目录
RUN mkdir -p /go/src/app

WORKDIR /go/src/app


docker run -id --name taotao-container \
-p 10001:10001 -p 10002:10002 -p 10003:10003 -p 10004:10004 -p 10005:10005 -p 10006:10006 -p 10007:10007 \
-p 11001:11001 -p 11002:10002 -p 11003:11003 -p 11004:11004 -p 11005:10005 -p 11006:11006 -p 11007:11007 \
my-golang-image