---
title: PHP 无限遍历数组
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
abbrlink: 65e7f3e8
date: 2021-08-17 13:38:38
img:
coverImg:
password:
summary:
---

# 无限遍历数组

```php

<?php
    function fun($arr){
        foreach($arr as $key => $value){
            if(is_array($value)){
                fun($value);
            }else{
                echo $value;
            }
        }
    }
   $array = ['a'=>'123','b'=>'456',['c'=>'666','d'=>'8888',['e'=>'4578','f'=>'484878']]]; 
   fun($array);
?>

```

## 对无限数组格式化数组中的值

```php

    public function index()
    {
        $data = [
            '1',
            '2',
            '3',
            [
                'a' => '333',
                ['b' => '111']
            ]
        ];
        $result = $this->recursion($data);
        dump($result);
    }
    
    public function recursion(array $data) : array
    {
        $res = [];
        if (is_array($data) && count($data)) {
            foreach ($data as $k => $v) {
                if (is_array($v) && count($v)) {
                    $res[$k] = $this->recursion($v);
                } elseif (is_string($v)) {
                    $res[$k] = (int)$v;
                }
            }
        }
        return $res;
    }
    
    // 或者这种
    
    public static function recursion(array &$data) : array
    {
        foreach ($data as $k => &$v) {
            if (! is_array($v)) {
                if (is_numeric($v)) {
                    $v = intval($v);
                }
            } else {
                static::recursion($v);
            }
        }
        return $data;
    }    

```

## 无限遍历与分类(一)

```php

$db = new mysqli('127.0.0.1', 'root', '123456', 'test');

$db->query("SET NAMES UTF8");

$res = $db->query('select * from menus');

$items = array();

while ($row = $res->fetch_array(MYSQLI_ASSOC)) {
    $items[$row['id']] = $row;
}

$tree = [];
foreach ($items as $key => $item) {
    $items[$item['pid']]['son'][$item['id']] = &$items[$item['id']];
}

// 所有数据堆成的数据树
$tree = $items[0]['son'];


function a ($tree , &$result = array(), &$level = 0) {
    if (is_array($tree)) {
        $level++;
        foreach ($tree as $id => $data) {
            $data['level'] = $level;
            $result[$id] = $data;
            if (isset($data['son'])) {
                a($data['son'], $result, $level);
                unset($result[$id]['son']);
            }
        }
        $level = 0;
    }
    return $result;
}

$newtree = a($tree);


echo json_encode($newtree);

```

## 无限遍历与分类(二)

```php

$db = new mysqli('127.0.0.1', 'root', '123456', 'test');

$db->query("SET NAMES UTF8");

$res = $db->query('select * from menus');

$items = array();

while ($row = $res->fetch_array(MYSQLI_ASSOC)) {
    $items[] = $row;
}

public function menulist($menu,$id=0,$level=0){
	
	static $menus = array();
	foreach ($menu as $value) {
		if ($value['pid']==$id) {
			$value['level'] = $level+1;
			if($level == 0)
			{
				$value['str'] = str_repeat('',$value['level']);
			}
			elseif($level == 2)
			{
				$value['str'] = '&emsp;&emsp;&emsp;&emsp;'.'└ ';
			}
			elseif($level == 3)
			{
				$value['str'] = '&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;'.'└ ';
			}
			else
			{
				$value['str'] = '&emsp;&emsp;'.'└ ';
			}
			$menus[] = $value;
			$this->menulist($menu,$value['id'],$value['level']);
		}
	}
	return $menus;
}

$newRows = menulist($items);	
	
```

## 无限遍历与分类 （三）

```php

<?php

$arrs = [
    [
        'id'=>1,
        'parent_id'=>0
    ],
    [
        'id'=>2,
        'parent_id'=>1
    ],
    [
        'id'=>3,
        'parent_id'=>2
    ],
    [
        'id'=>4,
        'parent_id'=>2
    ],
    [
        'id'=>5,
        'parent_id'=>0
    ],
];


function getTree($arrs,$root=0,$level=100)
{
    $tree = array();
    foreach ($arrs as $foo) {
        if ($foo['parent_id'] == $root) {

           if($level>0){               
            $foo['children'] = getTree($arrs,$foo['id'],$level-1);     
           }
            $tree[] = $foo;

        }
    }
    --$level;
    return $tree;
}


var_export ($arrs,0,0)// 一级

```

## 无限遍历与分类 （四）

对具有父子关系的数组进行重新排序，不改变数据结构

```php

    public function index()
    {

        $permissions = [
            ['id' => 1, 'pid' => '0', 'name' => 'a'],
            ['id' => 2, 'pid' => '0', 'name' => 'b'],
            ['id' => 3, 'pid' => '0', 'name' => 'c'],
            ['id' => 4, 'pid' => '1', 'name' => 'a-1'],
            ['id' => 5, 'pid' => '1', 'name' => 'a-2'],
            ['id' => 6, 'pid' => '3', 'name' => 'c-1'],
            ['id' => 7, 'pid' => '3', 'name' => 'c-2'],
            ['id' => 9, 'pid' => '4', 'name' => 'a-1-2'],
            ['id' => 8, 'pid' => '4', 'name' => 'a-1-1'],
        ];
        $data = $this->permissionTree($permissions);
        return $data;
    }

    /**
     * 根据父子关系重新排序
     *
     * @param $data  具有父子关系的二维数组
     * @param int $root  获取指定层级标识
     * @param array $result  用于保存数据的数组
     * @return array
     */
    public function permissionTree($data, $root = 0, &$result = [])
    {
        foreach ($data as $item) {
            // 排除掉非直接子集
            if ($item['pid'] != $root) {
                continue;
            }
            $result[$item['id']] = $item;
            $this->permissionTree($data, $item['id'], $result);
        }
        return $result;
    }

```
