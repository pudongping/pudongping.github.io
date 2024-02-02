---
title: ElasticSearch 聚合查询
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: ElasticSearch
tags:
  - ElasticSearch
abbrlink: 25480f12
date: 2024-02-02 17:04:39
img:
coverImg:
password:
summary:
---

## 聚合的基本语法结构

```json

{
  "aggregations": {
    "{aggregations_name_1}": {
      "{aggregations_type}": {
        {aggregations_body}
      }
    },
    "{aggregations_name_2}": {
      "{aggregations_type}": {
        {aggregations_body}
      }
    }    
  }
}

```

- `aggregations` 表示聚合查询语句，可以简写为 `aggs`
- `{aggregations_name_1}` 表示一个聚合计算的名称，可以随意命名，因为 es 支持一次进行多次统计分析查询，后面需要通过这个名字在查询结果中找到我们想要的计算结果
- `{aggregations_type}` 表示聚合类型，代表我们想要怎么统计数据，主要有两大类聚合类型，**桶聚合** 和 **指标聚合**，这两类聚合又包括多种聚合类型。指标聚合：sum、avg 桶聚合：terms
- `{aggregations_body}` 聚合类型的参数，选择不同的聚合类型，有不同的参数
- `{aggregations_name_2}` 表示其他聚合计算的名字，可以进行多种类型的统计

### value count 值聚合

value count 值聚合，主要用于统计文档总数，类似 sql 中的 `count` 函数

```json

GET /rcp_goods_img_checks/_search
{
  "size": 0, 
  "aggs": {
    "alex_count": { // 聚合查询的名字，可以随便取一个名字
      "value_count": {  // 聚合类型为：value_count
        "field": "app_name1"  // 计算 app_name1 这个字段值的总数
      }
    }
  }
}

# 返回值示例
{
  "took" : 0,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 2,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  },
  "aggregations" : {
    "alex_count" : {
      "value" : 2
    }
  }
}

```

### cardinality

cardinality 基数聚合，也用于统计文档的总数，跟 value count 的区别是，基数聚合会去重，不会统计重复的值，类似 sql 中的 `count(DISTINCT field)` 用法


```json

GET /rcp_goods_img_checks/_search
{
  "size": 0, 
  "aggs": {
    "alex_count": {
      "cardinality": {
        "field": "check_status"
      }
    }
  }
}

# 返回值示例
{
  "took" : 0,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 2,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  },
  "aggregations" : {
    "alex_count" : {
      "value" : 1
    }
  }
}

```

### avg

求平均值

```json

GET /rcp_goods_img_checks/_search
{
  "size": 0, 
  "aggs": {
    "alex_avg": {
      "avg": {
        "field": "check_status"
      }
    }
  }
}

# 返回值示例
{
  "took" : 0,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 3,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  },
  "aggregations" : {
    "alex_avg" : {
      "value" : 2.6666666666666665
    }
  }
}

```

### sum

求和

```json

GET /rcp_goods_img_checks/_search
{
  "size": 0, 
  "aggs": {
    "alex_sum": {
      "sum": {
        "field": "check_status"
      }
    }
  }
}

# 返回值示例
{
  "took" : 1,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 3,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  },
  "aggregations" : {
    "alex_sum" : {
      "value" : 8.0
    }
  }
}

```

### max

求最大值

```json

GET /rcp_goods_img_checks/_search
{
  "size": 0, 
  "aggs": {
    "alex_max": {
      "max": {
        "field": "check_status"
      }
    }
  }
}

# 返回值示例
{
  "took" : 7,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 3,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  },
  "aggregations" : {
    "alex_max" : {
      "value" : 3.0
    }
  }
}

```

### min

求最小值

```json

GET /rcp_goods_img_checks/_search
{
  "size": 0, 
  "aggs": {
    "alex_min": {
      "min": {
        "field": "check_status"
      }
    }
  }
}

# 返回值示例
{
  "took" : 0,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 3,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  },
  "aggregations" : {
    "alex_max" : {
      "value" : 2.0
    }
  }
}

```