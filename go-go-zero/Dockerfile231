FROM centos:7

RUN cd /etc/yum.repos.d/

RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*

RUN yum -y makecache
RUN yum -y install vim
RUN yum -y install wget
RUN yum -y install gcc

RUN wget https://studygolang.com/dl/golang/go1.19.1.linux-amd64.tar.gz
RUN tar -zxvf go1.19.1.linux-amd64.tar.gz
RUN mkdir -p /home/go/src

ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOROOT/bin
ENV GOPATH /home/go/src
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

EXPOSE 10001 10002 10003 10004 10005 10006 10007
EXPOSE 11001 11002 11003 11004 11005 11006 11007

VOLUME /home/go

ENV MYPATH /home/go/src

WORKDIR $MYPATH
