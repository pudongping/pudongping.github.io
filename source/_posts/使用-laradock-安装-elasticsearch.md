---
title: 使用 laradock 安装 elasticsearch
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: ElasticSearch
tags:
  - Docker
  - Laradock
  - ElasticSearch
abbrlink: deba288d
date: 2021-08-26 09:20:49
img:
coverImg:
password:
summary:
---

# 使用 Laradock 安装 ElasticSearch

- [ElasticSearch 可视化工具 ElasticHQ](https://github.com/ElasticHQ/elasticsearch-HQ) / [官网地址](http://www.elastichq.org/)

## 安装和使用

1. 使用 `docker-compose up` 命令运行 `ElasticSearch` 容器

```sh

docker-compose up -d elasticsearch

```

2. 打开浏览器并通过端口 `9200` 访问本地主机 [http://localhost:9200](http://localhost:9200)

> 默认用户是 user ，默认密码是 changeme

## 如果是在 laradock 中使用时

```sh

curl http://elasticsearch:9200

```

## 安装 ElasticSearch 插件

```sh

# 安装一个 ElasticSearch 插件
docker-compose exec elasticsearch /usr/share/elasticsearch/bin/elasticsearch-plugin install {plugin-name}

# 重启容器
docker-compose restart elasticsearch

```

### 安装 [elasticsearch-analysis-ik](https://github.com/medcl/elasticsearch-analysis-ik) 中文分词插件

比如，此时需要安装 [elasticsearch-analysis-ik](https://github.com/medcl/elasticsearch-analysis-ik) 中文分词插件，需要下载 ik 的 [releases](https://github.com/medcl/elasticsearch-analysis-ik/releases) 源码 zip 包

```sh

# 方式1，你可以直接在 elasticsearch 容器外，执行以下命令
docker-compose exec elasticsearch /usr/share/elasticsearch/bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.9.1/elasticsearch-analysis-ik-7.9.1.zip

# 方式2，你可以直接进入到 elasticsearch 容器内，然后执行以下命令
./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.9.1/elasticsearch-analysis-ik-7.9.1.zip

```

**需要注意的是：如果你的 elasticsearch 的版本是 7.9.1 那么，你安装的 ik 插件也必须是 7.9.1 的版本**，elasticsearch 的版本号可以通过访问 [http://localhost:9200/](http://localhost:9200/) 查看 `version.number` 字段查看，然后 `docker-compose restart elasticsearch` 重启 elasticsearch 容器即可

#### 安装 elasticsearch-analysis-ik 过程如下所示

```sh

[root@f1831cb3b4dd elasticsearch]# ./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.9.1/elasticsearch-analysis-ik-7.9.1.zip
-> Installing https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.9.1/elasticsearch-analysis-ik-7.9.1.zip
-> Downloading https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.9.1/elasticsearch-analysis-ik-7.9.1.zip
[=================================================] 100%??
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@     WARNING: plugin requires additional permissions     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
* java.net.SocketPermission * connect,resolve
See http://docs.oracle.com/javase/8/docs/technotes/guides/security/permissions.html
for descriptions of what these permissions allow and the associated risks.

Continue with installation? [y/N]y
-> Installed analysis-ik
[root@f1831cb3b4dd elasticsearch]# ./bin/elasticsearch-plugin list
analysis-ik

```

查看插件列表

```sh

./bin/elasticsearch-plugin list

```

## ElasticSearch 和 mysql 数据库的概念对比

MySQL |	Elasticsearch
--- | ----
表（Table）| 索引（Index）
记录（Row）| 文档（Document）
字段（Column） | 字段（Fields）

## ElasticSearch 的简单使用

### 新建索引 index （创建表）

```sh

curl -XPUT http://localhost:9200/test_index  

# 在 Elasticsearch 的返回中如果包含了 "acknowledged" : true, 则代表请求成功。
{"acknowledged":true,"shards_acknowledged":true,"index":"test_index"}

```

### 查看

```sh

curl http://localhost:9200/test_index

{"test_index":{"aliases":{},"mappings":{},"settings":{"index":{"creation_date":"1617069458624","number_of_shards":"1","number_of_replicas":"1","uuid":"XKjqatZTSOu9I_PiwzaNOQ","version":{"created":"7090199"},"provided_name":"test_index"}}}}

# 可以加上 pretty 参数，返回比较人性化的结构
curl http://localhost:9200/test_index\?pretty                                                                           
{
  "test_index" : {
    "aliases" : { },
    "mappings" : { },
    "settings" : {
      "index" : {
        "creation_date" : "1617069458624",
        "number_of_shards" : "1",
        "number_of_replicas" : "1",
        "uuid" : "XKjqatZTSOu9I_PiwzaNOQ",
        "version" : {
          "created" : "7090199"
        },
        "provided_name" : "test_index"
      }
    }
  }
}

```

### 创建类型

对应的接口地址是 /{index_name}/_mapping

```sh

curl -H'Content-Type: application/json' -XPUT http://localhost:9200/test_index/_mapping?pretty -d'{
  "properties": {
    "title": { "type": "text", "analyzer": "ik_smart" }, 
    "description": { "type": "text", "analyzer": "ik_smart" },
    "price": { "type": "scaled_float", "scaling_factor": 100 }
  }
}'

# 会返回

{
  "acknowledged" : true
}


curl -H'Content-Type: application/json' -XPUT http://localhost:9200/products/_mapping/?pretty -d'{
  "properties": {
    "brand_id": { "type": "integer" },
    "type": { "type": "integer" },
    "title": { "type": "text", "analyzer": "ik_smart" }, 
    "unit": { "type": "keyword" },
    "sketch": { "type": "text", "analyzer": "ik_smart" }, 
    "keywords": { "type": "text", "analyzer": "ik_smart" },
    "tags": { "type": "keyword" },
    "barcode": { "type": "keyword" },
    "price": { "type": "scaled_float", "scaling_factor": 100 },
    "market_price": { "type": "scaled_float", "scaling_factor": 100 },
    "rating": { "type": "float" },
    "sold_count": { "type": "integer" },
    "review_count": { "type": "integer" },    
    "virtual_retail_num": { "type": "integer" },
    "description": { "type": "text", "analyzer": "ik_smart" },
    "stock": { "type": "integer" },    
    "warning_stock": { "type": "integer" },   
    "main_image": { "type": "keyword" },
    "slider_image": { "type": "keyword" },
    "status": { "type": "integer" },
    "is_hot": { "type": "integer" },
    "sort": { "type": "integer" },
    "categories": {
      "type": "nested",
      "properties": {
        "id": { "type": "integer", "copy_to": "categories_id" },
        "pid": { "type": "integer" },
        "name": { "type": "text", "analyzer": "ik_smart", "copy_to": "categories_name" }, 
        "description": { "type": "text", "analyzer": "ik_smart", "copy_to": "categories_description" },
        "status": { "type": "integer" },
        "level": { "type": "integer" },
        "img": { "type": "keyword" }
      }
    },    
    "brand": {
      "type": "nested",
      "properties": {
        "id": { "type": "integer" },
        "name": { "type": "text", "analyzer": "ik_smart", "copy_to": "brand_name" }, 
        "description": { "type": "text", "analyzer": "ik_smart", "copy_to": "brand_description" },
        "log_url": { "type": "keyword" },
        "img": { "type": "keyword" }
      }
    },      
    "attrs": {
      "type": "nested",
      "properties": {
        "id": { "type": "integer" },
        "name": { "type": "keyword", "copy_to": "attrs_name" }
      }
    },  
    "skus": {
      "type": "nested",
      "properties": {
        "id": { "type": "integer" },
        "name": { "type": "text", "analyzer": "ik_smart"}, 
        "main_url": { "type": "keyword" },
        "price": { "type": "scaled_float", "scaling_factor": 100 },
        "sold_count": { "type": "integer" }
      }
    }
  }
}'

```

- 提交数据中的 `properties` 代表这个索引中各个字段的定义，其中 key 为字段名称，value 是字段的类型定义
- `type` 定义了字段的数据类型，常用的有 `text` / `integer` / `date` / `boolean` ，还有[更多类型](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-types.html)
    - `keyword`，这是字符串类型的一种，这种类型是告诉 Elasticsearch 不需要对这个字段做分词，通常用于邮箱、标签、属性等字段。
    - `scaled_float` 代表一个小数位固定的浮点型字段，与 Mysql 的 decimal 类型类似。
    - `scaling_factor` 用来指定小数位精度，100 就代表精确到小数点后两位。
    - `nested` 代表这个字段是一个复杂对象，由下一级的 properties 字段定义这个对象的字段。
- `analyzer`是一个新的概念，这是告诉 Elasticsearch 应该用什么方式去给这个字段做分词，这里我们用了 `ik_smart`，是一个中文分词器。
- `copy_to`，Elasticsearch 的多字段匹配查询是不支持查询 Nested 对象的字段，但是我们又必须查询 `categories.name` 字段，因此我们可以使用 `copy_to` 参数，可以将  `categories.name` 字段复制到上层，我们就可以通过 `categories_name` 字段做多字段匹配查询


### 创建文档

对应的接口地址是 /{index_name}/_doc/{id} 这里的 id 和 mysql 中的 id 不一样，不是自增的，需要我们手动指定。

```sh

# 创建 id 为 1 的文档
curl -H'Content-Type: application/json' -XPUT http://localhost:9200/test_index/_doc/1?pretty -d'{
    "title": "iPhone 7P",
    "description": "iphone 第一批双摄像头",
    "price": 6799
}'

# 会返回如下内容
{
  "_index" : "test_index",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 1,
  "result" : "created",
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 0,
  "_primary_term" : 2
}


# 创建 id 为 2 的文档
curl -H'Content-Type: application/json' -XPUT http://localhost:9200/test_index/_doc/2?pretty -d'{
    "title": "OPPO find x",
    "description": "高清像素",
    "price": 3499
}'

# 会返回如下内容
{
  "_index" : "test_index",
  "_type" : "_doc",
  "_id" : "2",
  "_version" : 1,
  "result" : "created",
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 1,
  "_primary_term" : 2
}

```

### 读取文档数据

```sh

curl http://localhost:9200/test_index/_doc/1\?pretty

# 会返回如下内容
{
  "_index" : "test_index",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 1,
  "_seq_no" : 0,
  "_primary_term" : 2,
  "found" : true,
  "_source" : {
    "title" : "iPhone 7P",
    "description" : "iphone 第一批双摄像头",
    "price" : 6799
  }
}

```

### 查看 Elasticsearch 索引中有多少条数据

对应的接口地址为 /{index_name}/_doc/_count

```sh

curl http://localhost:9200/test_index/_doc/_count\?pretty

# 会返回如下内容
{
  "count" : 3,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  }
}

```

### 简单搜索

```sh

curl -XPOST -H'Content-Type:application/json' http://localhost:9200/test_index/_doc/_search\?pretty -d'
{
    "query" : { "match" : { "description" : "iphone" }}
}'

# 会返回如下内容
{
  "took" : 16,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 1,
      "relation" : "eq"
    },
    "max_score" : 0.60996956,
    "hits" : [
      {
        "_index" : "test_index",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.60996956,
        "_source" : {
          "title" : "iPhone 7P",
          "description" : "iphone 第一批双摄像头",
          "price" : 6799
        }
      }
    ]
  }
}

```
