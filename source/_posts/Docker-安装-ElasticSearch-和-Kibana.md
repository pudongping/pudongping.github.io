---
title: Docker 安装 ElasticSearch 和 Kibana
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: ElasticSearch
tags:
  - Docker
  - ElasticSearch
  - Kibana
abbrlink: 2e9eaf3a
date: 2022-08-06 18:26:07
img:
coverImg:
password:
summary:
---

# Docker 安装 ElasticSearch 和 Kibana

> 安装的时候，ElasticSearch 和 Kibana 的版本一定要一致。

## 创建一个网络

```bash
# 创建一个网络
docker network create alex-network
```

## 安装 ElasticSearch

- 拉取镜像

```bash
docker pull elasticsearch:7.9.3
```

- 创建容器

```bash
# 设置为单节点（非集群模式）
# -e "discovery.type=single-node"
# 设置 es 的初始内存和最大内存，否则导致过大启动不了 es 或者 kibana 连接 es 时，导致 es 直接死机
# -e ES_JAVA_OPTS="-Xms512m -Xmx512m"
# 挂载逻辑卷，绑定 es 的数据目录
# -v ~/es-data:/usr/share/elasticsearch/data
# 挂载逻辑卷，绑定 es 的插件目录
# -v ~/es-plugins:/usr/share/elasticsearch/plugins
# 挂载逻辑卷，绑定 es 的配置目录
# -v ~/es-config:/usr/share/elasticsearch/config
# 授权逻辑卷访问权限
# --privileged

docker run -d --name alex-es \
--net alex-network \
-p 9200:9200 \
-p 9300:9300 \
-e "discovery.type=single-node" \
-e ES_JAVA_OPTS="-Xms512m -Xmx512m" \
elasticsearch:7.9.3


# 如果要挂载目录时，可能会报错 `Exception in thread "main" java.nio.file.NoSuchFileException: /usr/share/elasticsearch/config/jvm.options`
# 解决方案是：先随便创建一个容器，然后将随便创建的容器的配置文件先复制一份到宿机上，然后删除掉这个容器，重新再创建容器，比如，如下步骤

# 1. 先随便创建一个容器
docker run -d --name es_service_1 \
--net alex-network \
-p 9200:9200 \
-p 9300:9300 \
-e "discovery.type=single-node" \
-e ES_JAVA_OPTS="-Xms512m -Xmx512m" \
elasticsearch:7.9.3

# 2. 然后复制该容器的配置文件到宿机上
docker cp {es_service_1_container_id}:/usr/share/elasticsearch/config/ ~/es-config

# 3. 然后再删除到这个临时容器
docker rm es_service_1

# 4. 最后再携带挂载参数创建新容器
docker run -d --name es_service_1 \
--net alex-network \
-p 9200:9200 \
-p 9300:9300 \
-e "discovery.type=single-node" \
-e ES_JAVA_OPTS="-Xms512m -Xmx512m" \
-v ~/es-data:/usr/share/elasticsearch/data \
-v ~/es-plugins:/usr/share/elasticsearch/plugins \
-v ~/es-config:/usr/share/elasticsearch/config \
--privileged -u root \
elasticsearch:7.9.3

```

- 也可以直接通过修改配置文件

```bash

# 编辑 es 的配置文件
vi /usr/share/elasticsearch/config/elasticsearch.yml

# 如果需要解决跨域问题，那么则在配置文件中追加写入以下内容
http.cors.enabled: true
http.cors.allow-origin: "*"


# 编辑 jvm 相关配置
vi /usr/share/elasticsearch/config/jvm.options

# 如果 es 容器老是会宕机，那么可能需要调整以下参数
# 修改以下参数（按照自己的实际需求进行调整）
-Xms512m
-Xmx512m

```

- 测试是否安装成功

```bash
# 出现版本号即表示成功
curl 127.0.0.1:9200
```

### 安装 ik 中文分词插件

> 注意 ik 分词插件的版本要和 es 的版本一致

```bash

[root@63d697bd738f elasticsearch]# pwd
/usr/share/elasticsearch
[root@63d697bd738f elasticsearch]# ./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.9.3/elasticsearch-analysis-ik-7.9.3.zip
-> Installing https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.9.3/elasticsearch-analysis-ik-7.9.3.zip
-> Downloading https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.9.3/elasticsearch-analysis-ik-7.9.3.zip
[=================================================] 100%??
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@     WARNING: plugin requires additional permissions     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
* java.net.SocketPermission * connect,resolve
See http://docs.oracle.com/javase/8/docs/technotes/guides/security/permissions.html
for descriptions of what these permissions allow and the associated risks.

Continue with installation? [y/N]y
-> Installed analysis-ik
[root@63d697bd738f elasticsearch]# ./bin/elasticsearch-plugin list
analysis-ik
[root@63d697bd738f elasticsearch]#

```

## 安装 Kibana

- 拉取镜像

```bash
docker pull kibana:7.9.3
```

- 创建容器

```bash
# 相关参数
# 设置 es 请求地址
# -e ELASTICSEARCH_HOSTS=http://localhost:9200
# 设置汉化
# -e I18N_LOCALE=zh-CN
# 设置时区，否则查询时间需要加 8 个小时
# -e TZ='Asia/Shanghai'

docker run -d --name alex-kibana \
-p 5601:5601 \
--net alex-network \
-e ELASTICSEARCH_HOSTS=http://alex-es:9200 \
-e I18N_LOCALE=zh-CN \
kibana:7.9.3

```

- 也可以直接去编辑 kibana 的配置文件

```bash

# 编辑配置文件
vi /usr/share/kibana/config/kibana.yml

# 适当修改以下配置，因为我们已经将 alex-es 容器和 alex-kibana 容器同时都加入到了 alex-network 网络中，因此这里，我们可以直接通过容器名称进行访问到容器对应的 IP 地址
elasticsearch.hosts: [ "http://alex-es:9200" ]
i18n.locale: zh-CN

```

- 测试是否安装成功

```bash
# 浏览器请求
127.0.0.1:5601
```
