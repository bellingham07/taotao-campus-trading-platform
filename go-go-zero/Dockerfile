FROM centos:7

RUN set -eux
# 安装必要的组件
RUN yum install -y epel-release && \
        yum update -y && \
        yum install -y git wget curl which vim && \
        yum clean all &&  \
        yum makecache
RUN wget -O /etc/yum.repos.d/CentOS-Base.repo https://mirrors.aliyun.com/repo/Centos-7.repo
RUN wget -O /etc/yum.repos.d/epel.repo https://mirrors.aliyun.com/repo/epel-7.repo

# 下载golang，先设定版本，然后应用，加上官网的SHA256的校验和
ENV GOLANG_VERSION 1.20.5
ENV GOLANG_DOWNLOAD_URL https://go.dev/dl/go1.20.5.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 d7ec48cde0d3d2be2c69203bc3e0a44de8660b9c09a6e85c4732a3f7dc442612

# 安装 Golang 运行环境
#ADD go1.20.5.linux-amd64.tar.gz /usr/local/go/
RUN mkdir /usr/local/go

# 设置环境变量
ENV GOROOT /usr/local/go
ENV PATH "$GOROOT/bin:$PATH"
ENV PATH "$PATH:$GOPATH/bin"
ENV GO111MODULE on
ENV GOPROXY=https://goproxy.cn,direct

# 下载golang
RUN wget $GOLANG_DOWNLOAD_URL -O go.tar.gz
RUN tar -C /usr/local/ -zxvf go.tar.gz

## 按装 go-zero 框架需要的 goctl 工具
#RUN go install github.com/zeromicro/go-zero/tools/goctl@latest
#
## 使用 goctl 一键安装安装 protobuf 工具（protoc、proto-gen-go、proto-gen-go-grpc）
#RUN goctl env check --install --verbose --force

# 创建工作目录
RUN mkdir -p /go/src/app

WORKDIR /go/src/app