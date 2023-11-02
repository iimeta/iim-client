<div align=center>

<img src="https://iim.ai/public/images/logo.png" width="256"/>

# 智元 IM

[![Release](https://img.shields.io/github/v/release/iimeta/iim-client?color=blue)](https://github.com/iimeta/iim-client/releases)
[![License](https://img.shields.io/static/v1?label=license&message=MIT&color=green)](https://github.com/iimeta/iim-client/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/iimeta/iim-client)](https://goreportcard.com/report/github.com/iimeta/iim-client)

</div>

`智元 IM` 是一个网页版即时聊天系统, 同时拥有AI聊天对话功能, 本项目基于[Lumen IM](https://github.com/gzydong/go-chat)使用[GoFrame](https://github.com/gogf/gf)框架开发

[![Release](https://img.shields.io/static/v1?label=release&message=v1.0.0&color=blue)](https://github.com/iimeta/iim-client/releases)
[![License](https://img.shields.io/static/v1?label=license&message=MIT&color=green)](https://github.com/iimeta/iim-client)

## 手动Docker部署
0. 需提前安装好Redis和MongoDB以及Docker服务
1. 拉取镜像: `docker pull iimeta/iim-client:1.0.0`
2. 创建配置文件目录: `mkdir -p /data/iim-client/manifest/config`
3. 创建公开文件目录: `mkdir -p /data/iim-client/resource/public`
4. 创建私有文件目录: `mkdir -p /data/iim-client/resource/private`
5. 进入配置文件目录: `cd /data/iim-client/manifest/config`
6. 下载配置文件: `wget https://github.com/iimeta/iim-client/raw/docker/manifest/config/config.yaml`
7. 修改配置文件: `vi config.yaml` 按 config.yaml 中配置项的注释, 修改对应参数
8. 启动服务: 
```shell
docker run -d \
  --network host \
  --restart=always \
  -p 8000:8000 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/iim-client/manifest/config/config.yaml:/app/manifest/config/config.yaml \
  -v /data/iim-client/resource/public:/app/resource/public \
  -v /data/iim-client/resource/private:/app/resource/private \
  --name iim-client \
  iimeta/iim-client:1.0.0
```
9. 访问: http://ip:8000/ 如: http://localhost:8000/

## 自动Docker部署
0. 需提前安装好Redis和MongoDB以及Docker服务
1. 下载一键部署脚本: `wget https://github.com/iimeta/iim-client/raw/docker/bin/install.sh`
2. 执行安装: `chmod +777 install.sh && ./install.sh`
3. 修改配置文件: `vi /data/iim-client/manifest/config/config.yaml` 按 config.yaml 中配置项的注释, 修改对应参数
4. 启动服务: `chmod +777 start.sh && ./start.sh`
5. 访问: http://ip:8000/ 如: http://localhost:8000/

## 备用国内镜像地址
```shell
docker pull registry.cn-hangzhou.aliyuncs.com/iimeta/iim-client:1.0.0
```
```shell
docker run -d \
  --network host \
  --restart=always \
  -p 8000:8000 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/iim-client/manifest/config/config.yaml:/app/manifest/config/config.yaml \
  -v /data/iim-client/resource/public:/app/resource/public \
  -v /data/iim-client/resource/private:/app/resource/private \
  --name iim-client \
  registry.cn-hangzhou.aliyuncs.com/iimeta/iim-client:1.0.0
```


![Star History Chart](https://api.star-history.com/svg?repos=iimeta/iim-client&type=Date)
