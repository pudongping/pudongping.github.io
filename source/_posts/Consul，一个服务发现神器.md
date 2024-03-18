---
title: Consul，一个服务发现神器
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 分布式系统
tags:
  - 分布式系统
  - Consul
abbrlink: 74705a17
date: 2024-03-19 00:08:05
img:
coverImg:
password:
summary:
---

在使用微服务的架构中，Consul 是一个不可或缺的工具。它由 HashiCorp 公司开发，Consul 提供了服务发现、健康检查、配置管理和服务网格功能，使得在复杂的分布式系统中管理和配置服务变得简单高效。

## 服务发现：解决微服务寻址难题

在分布式系统中，服务实例通常会动态变化，这使得服务间的互相寻址变得复杂。Consul通过服务注册和服务发现机制解决了这一问题。每个服务实例在启动时向Consul注册自己，并周期性地更新其健康状态。这样，服务之间就可以通过Consul查询其他服务的当前可用实例。

```bash
# 注册服务示例
curl --location --request PUT 'http://localhost:8500/v1/agent/service/register' \
--header 'Content-Type: application/json' \
--data '{
  "ID": "my-service-1",
  "Name": "my-service",
  "Tags": ["v1"],
  "Address": "10.0.0.1",
  "Port": 8080,
  "Check": {
    "HTTP": "http://10.0.0.1:8080/health",
    "Interval": "10s"
  }
}
'
```

## 健康检查：确保服务质量和可靠性

Consul 的健康检查机制确保只有健康的服务实例被用于服务发现。每个服务可以定义一个或多个健康检查，Consul  会定期执行这些检查以确定服务实例的健康状态。

## 动态配置：灵活的服务配置管理

Consul 的 `Key/Value` 存储功能允许程序员存储和管理配置信息。服务可以在启动时从 Consul 查询这些配置，实现动态配置管理。当配置更改时，服务可以自动获取新的配置，而无需重启。

## 多数据中心支持：扩展全球服务

Consul 天然支持多数据中心，这使得构建跨地理位置的高可用系统变得简单。程序员可以配置服务在多个数据中心之间相互发现和通信，而 Consul 会处理所有复杂的网络和同步问题。

## 以下是 Consul 的几种安装方式

### 第一种方式：可以直接使用 `homebrew` 安装

```bash
brew install consul
```

### 第二种方式：使用 docker 安装 consul

```bash
# 拉取 latest 版本镜像
docker pull consul

# 拉取指定版本的 consul 镜像，这里以 1.14.1 版本示例
docker pull consul:1.14.1

# -bind 的地址为 docker0 或者 eth0 的 ip 地址都可以
docker run --name alex-consul \
--net=host \
-d consul:1.14.1 \
agent -server -bind=172.17.118.227 -bootstrap-expect=1 -client=0.0.0.0 -ui
```

### 第三种方式：直接下载 consul 二进制文件启动

> [官方下载页面](https://www.consul.io/downloads "官方下载页面") 可以按照系统，选择指定的 consul 版本  
我这里下载的是 `Mac 版本的 consul 1.14.2 版本`， [下载地址](https://releases.hashicorp.com/consul/1.14.2/consul_1.14.2_darwin_amd64.zip "下载地址") 下载完成之后是一个 `zip` 压缩包，直接解压即可。

```bash
# 下载 mac 版本的 consul 1.14.2
wget https://releases.hashicorp.com/consul/1.14.2/consul_1.14.2_darwin_amd64.zip
# 或者
curl -O https://releases.hashicorp.com/consul/1.14.2/consul_1.14.2_darwin_amd64.zip

# 解压
unzip consul_1.14.2_darwin_amd64.zip
```

直接使用二进制文件启动时：

```bash
# data 文件夹用于存放 consul 的数据文件， config 文件夹用于存放配置文件
# 需要将 client 设置成 0.0.0.0 否则无法在外网访问 consul 的 UI 界面

# 启动命令
./consul agent -data-dir=./data -config-dir=./config -server -bind=127.0.0.1 -bootstrap-expect=1 -client 0.0.0.0 -ui
# 或者
./consul agent -data-dir=./data -config-dir=./config -server -bootstrap -ui -node=1 -client=0.0.0.0
```

## 命令

```bash
# 通过 -dev 启动，不会保存配置（即重启 consul 后配置信息将消失）
consul agent -dev

# 查看集群成员。添加 `-detailed` 选项可以查看到额外的信息。
consul members

# 查看准确匹配 server 的状态信息
curl localhost:8500/v1/catalog/nodes
```
