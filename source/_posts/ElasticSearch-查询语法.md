---
title: ElasticSearch 查询语法
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: ElasticSearch
tags:
  - ElasticSearch
abbrlink: 65a50e61
date: 2024-02-02 16:55:14
img:
coverImg:
password:
summary:
---

## 查询的基本语法结构

```json

GET /{index}/{type}/_search
{
    "from": 0,  // 搜索结果的开始位置
    "size": 5,  // 分页大小，每页显示数量
    "_source": ["field_1", "field_2", "field_3"],  // 需要返回的字段数组
    "query": {
        "bool": {
            "must": [],  // must 条件，类似 sql 中的 and，代表必须匹配条件
            "must_not": [],  // must_not 条件，跟 must 相反，必须不匹配条件，相当于 sql 语句中的 `!=`
            "should": [],  // should 条件，类似 sql 中的 or，代表匹配其中一个条件
        }
    },  // query 子句
    "aggs": {},  // aggs 子句，主要用来编写统计分析语句，类似 sql 的 group by 子句
    "sort": [
        {
            "{field1}": {  // 需要排序的字段
                "order": "desc"  // 排序的规则，asc 升序，desc 降序
            }
        },
        {
            "{field2}": {
                "order": "asc"
            }
        }
    ]  // 排序子句
}


# 也可以多个索引同时查询
# 多个索引同时查询时，用逗号分隔，比如：
GET /{index1},{index2},{index3}/_doc/_search
# 或者使用前缀模糊匹配，比如：
GET /rcp_goods_img_*/_doc/_search

```

## Filter 和 Query

任何一个 document 对于 filter 来说，就是 **match 与否** 的问题，是个二值问题，0 和 1 没有 **scroing** 的过程，但是使用 query 的时候，是表示 **match 程度问题，有 scroing 过程**。

es 底层对 filter 做了很多优化，会对过滤结果进行缓存，同时，filter 没有相关性计算过程，所以 filter 比 query 快。

所以，官网推荐，作为一条比较通用的规则：***仅在全文检索时，使用 Query，其它时候都用 Filter***，但是还是得按照我们的实际情况来看，毕竟有些时候用空间换时间不一定划算。


## match 和 term

match 查询的时候，es 会根据你给定的字段提供合适的分析器，而 term 查询不会有分析器分析的过程，match 查询相当于模糊匹配，只包含其中一部分关键词就行。**match 系列匹配时，datatype 要设置为 text ，否则不会开启分词**

### match_all

表示取出所有 documents，在与 filter 结合使用时，会经常用到 match_all

```json

# 查询所有数据（查询所有的索引对应的所有文档）
GET _search
{
  "query": {
    "match_all": {}
  }
}

# 查指定索引下所有的数据
GET /rcp_goods_img_*/_doc/_search
{
  "query": {
    "match_all": {}
  }
}

```

### match

一般在全文检索时使用，首先利用 analyzer（分词器）对具体查询字符串进行分析，然后进行查询。但是在查询，比如说数值型字段、日期型字段、布尔字段或者不需要分词的字符串上进行查询时，表示精确匹配

```json

# 示例
# select goods_id, app_name from rcp_goods_img_checks._doc where goods_id = m43337929006 order by app_name desc
GET /rcp_goods_img_checks/_doc/_search
{
  "_source": [
    "goods_id",
    "app_name"
  ],
  "query": {
    "match": {
      "goods_id": "m43337929006"
    }
  },
  "sort": [
    {
      "app_name": {
        "order": "desc"
      }
    }
  ]
}

```

### multi_match

同时对查询的关键词进行多字段同时匹配，只要其中一个字段匹配到值就返回结果

```json

# 示例
# 只要查询的字段 check_status 、app_name 、site_id 这三个字段中任意一个字段为 19 那么就返回结果
# select * from rcp_goods_img_checks._doc where check_status = 19 or app_name = 19 or site_id = 19
GET /rcp_goods_img_checks/_doc/_search
{
  "query": {
    "multi_match": {
      "query": 19,
      "fields": ["check_status", "app_name", "site_id"]
    }
  }
}

```

### match_phrase_prefix

左前缀匹配，类似 sql 中的 like `field%`

```json

# 示例
# select * from rcp_goods_img_checks._doc where goods_id like 'm%'
GET /rcp_goods_img_checks/_doc/_search
{
  "query": {
    "match_phrase_prefix": {
      "goods_id": "m"
    }
  }
}

```

### term

term 用于精确查找，可用于数值、日期、布尔值或者不需要分词的字符串，当使用 term 时，不会对查询字符串进行分析，进行的是精确查找

```json

# 示例

# 字段只有一个值时，精确查询
# select * from rcp_goods_img_checks._doc where site_id = 18
GET /rcp_goods_img_checks/_search
{
  "query": {
    "term": {
      "site_id": "18"
    }
  }
}

# 因为是精确查询，因此不需要查询进行评分计算
# 使用 constant_score 查询以非评分模式来执行 term 查询，并作为统一评分
GET /rcp_goods_img_checks/_search
{
  "query": {
    "constant_score": {
      "filter": {
        "term": {
          "site_id": "18"
        }
      }
    }
  }
}

```

### terms

terms 和 term 类似，terms 可以指定多个值，只要 doc 满足 terms 里的任意值，就是满足查询条件的，属于精确查找

**terms 表示的是 contains 关系，而不是 equals 关系**

```json

# 示例

# 字段有多个值时
# select * from rcp_goods_img_checks._doc where site_id in ("18", "20")
GET /rcp_goods_img_checks/_search
{
  "query": {
    "terms": {
      "site_id": ["18", "20"]
    }
  }
}

# 使用 constant_score 查询以非评分模式来执行 terms 查询，并作为统一评分
GET /rcp_goods_img_checks/_search
{
  "query": {
    "constant_score": {
      "filter": {
        "term": {
          "site_id": ["18", "20"]
        }
      }
    }
  }
}

```

### range

范围查找

操作符可以是：
- **gt**：大于
- **gte**：大于等于
- **lt**：小于
- **lte**：小于等于

```json

# select * from rcp_goods_img_checks._doc where site_id >= 19 and site_id < 30
GET /rcp_goods_img_checks/_doc/_search
{
  "query": {
    "range": {
      "site_id": {
        "gte": 19,
        "lt": 30
      }
    }
  }
}

```

### exists 和 missing

exists 用于查找某个字段含有一个或者多个值对应的 document，而 missing 用于查找某个字段不存在值对应的 document，类比关系数据库中的 **is not null 等同于 exists** 和 **is null 等同于 missing**

```json

# select * from rcp_goods_img_checks._doc where title is not null

{
    "exists": {
        "field": "title"
    }
}


# select * from rcp_goods_img_checks._doc where title is null

{
    "missing": {
        "field": "title"
    }
}

```

### bool

使用 bool 子句来将各种子查询关联起来，组建布尔表达式，实现复合查询，bool 子句可以随意组合、嵌套。

bool 子句主要包括：

- must：表示必须匹配，与 and 等价（贡献算分）
- must_not：表示一定不能匹配，与 not 等价（不贡献算分）
- should：表示可以匹配，类似于布尔运算里面的**逻辑或** 与 or 等价（贡献算分）
- filter：过滤子句，必须匹配（不贡献算分）

如果 bool 子句里，没有 must 子句，那么，should 子句里至少匹配一个，如果有 must 子句，那么，should 子句至少匹配 0 个。可以使用 `minimum_should_match` 来对最小匹配数进行设置。

```json

{
    "bool" : {
        "must" : {
            "term" : { "name" : "kimchy" }
        },
        "must_not" : {
            "range" : {
                "age" : { "from" : 10, "to" : 20 }
            }
        },
        "should" : [
            {
                "term" : { "tag" : "v1" }
            },
            {
                "term" : { "tag" : "v1.0.2" }
            }
        ],
        "minimum_should_match" : 1,
        "boost" : 1.0
    }
}

# select * from rcp_goods_img_checks._doc where app_name = 2 and site_id = 18 and updated_at between 0 and 1659695794 order by updated_at desc limit 100
GET /rcp_goods_img_checks/_doc/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "term": {
            "app_name": 2
          }
        },
        {
          "term": {
            "site_id": 18
          }
        }
      ],
      "filter": [
        {
          "range": {
            "updated_at": {
              "from": 0,
              "to": 1659695794
            }
          }
        }
      ]
    }
  },
  "size": 100,
  "sort": [
    {
      "updated_at": "desc"
    }
  ]
}

```