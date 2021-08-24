---
title: laravel 使用 ORM 写复杂 sql
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - Laravel
abbrlink: 778c83e3
date: 2021-08-24 20:57:08
img:
coverImg:
password:
summary:
---

# laravel 使用 ORM 写复杂 sql

##  直接先贴出 sql 查询语句如下

```mysql

SELECT
	`assets_device`.* 
FROM
	(
		(
		SELECT
			`id`,
			`sub_category_id`,
			`number`,
			`name`,
			`location`,
			`officeBuilding`,
			`area`,
			`department`,
			`rack`,
			`rack_pos` 
		FROM
			`assets_device` 
		WHERE
			`department` IN ( 6, 7, 17, 20 ) 
			AND `assets_device`.`deleted_at` IS NULL 
		) UNION
		(
		SELECT
			`id`,
			`sub_category_id`,
			`number`,
			`name`,
			`location`,
			`officeBuilding`,
			`area`,
			`department`,
			`rack`,
			`rack_pos` 
		FROM
			`assets_device` 
		WHERE
			`area` IN ( 13, 14 ) 
			AND `assets_device`.`deleted_at` IS NULL 
		) 
	) AS assets_device
```

## 实现代码如下所示

```php

    public function getChoiceAssetsByErOrDt($request){

        $input = $request->input();

        // 机房 id 数组
        if(isset($input['er'])){
            if(is_array($input['er'])){
                $erIds = $input['er'];
            }else{
                $erIds = explode(',',$input['er']);
            }
        }else{
            $erIds = array();
        }

        // 科室 id 数组
        if(isset($input['dt'])){
            if(is_array($input['dt'])){
                $dtIds = $input['dt'];
            }else{
                $dtIds = explode(',',$input['dt']);
            }
        }else{
            $dtIds = array();
        }

        // id, 子分类id, 资产编号, 资产名称, 位置, 办公楼, 机房, 科室, 机柜, 机柜U数
        $field = ['id','sub_category_id','number','name','location','officeBuilding','area','department','rack','rack_pos'];
        
        if($erIds){
            // 根据机房 id 查询资产
            $erModel = $this->deviceModel->select($field)
                ->whereIn('area',$erIds);
        }

        if($dtIds){
            // 根据科室 id 查询资产
            $dtModel = $this->deviceModel->select($field)
                ->whereIn('department',$dtIds);
        }

        $model = $dtModel->union($erModel);

// 注意合并参数时 $subQuery 必须是 \Illuminate\Database\Query\Builder 类型
// 如果是 \Illuminate\Database\Eloquent\Builder 类型的，用 getQuery() 方法
        $data = $this->deviceModel->with('sub_category','zone','office_building','engineroom','department')
                                    ->mergeBindings($model->getQuery())
                                    ->select(["assets_device.*"])
                                    ->from(DB::raw("({$model->toSql()}) as assets_device"))
                                    ->withTrashed()
                                    ->orderBy('assets_device.id')
                                    ->paginate()
                                    ->toArray();
        
        // 以下代码如果使用 transformer 的话，就不需要写，直接可以在 transformer 里面转换
        foreach ($data['data'] as &$val){
            $val['department_name'] = isset($val['department']['name']) ? $val['department']['name'] : '';
            $val['sub_category_name'] = isset($val['sub_category']['name']) ? $val['sub_category']['name'] : '';
            $val['zone_name'] = isset($val['zone']['name']) ? $val['zone']['name'] : '';
            $val['office_building_name'] = isset($val['office_building']['name']) ? $val['office_building']['name'] : '';

            if(isset($val['engineroom']['name'])){
                $erdt = $val['engineroom']['name'];
            }elseif (isset($val['department']['name'])){
                $erdt = $val['department']['name'];
            }else{
                $erdt = null;
            }
            $val['erdt'] = $erdt;
            unset($val['department']);
            unset($val['sub_category']);
            unset($val['zone']);
            unset($val['office_building']);
            unset($val['engineroom']);
            unset($val['department']);
        }

        return $data;

    }
```
