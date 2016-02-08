# 阿里云部署指南

## 从自定义镜像创建
+  (可选) 创建RDS实例, 创建数据库, 添加账号, 设置权限
+  使用自定义镜像创建ECS实例, 将其内网IP加入RDS白名单中
+  (可选) 修改`/root/sqlServer.sh`内的账号密码部分, 以及`/root/confs/api.conf`的数据库配置部分
+  (可选) 为ECS实例创建ssh秘钥, 并把公钥添加到git服务器

## 从头开始
+  安装golang: `tar -C /usr/local -xzf go1.5.3.linux-amd64.tar.gz`
+  配置环境变量 `/root/.profile`加入:

```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/goPath
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOPATH/bin
```

+  安装依赖包:

```bash
mkdir goPath
go get github.com/astaxie/beego
go get github.com/beego/bee
go get github.com/stretchr/testify/assert
go get github.com/denverdino/aliyungo/oss
go get github.com/denverdino/aliyungo/common
go get github.com/denverdino/aliyungo/util
go get github.com/denverdino/aliyungo/slb
go get github.com/denverdino/aliyungo/ecs
go get github.com/denverdino/aliyungo/dns
apt-get install mysql-client
```

+  下载代码
