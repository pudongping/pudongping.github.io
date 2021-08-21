---
title: php 抽奖算法（适合九宫格和大转盘）
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
abbrlink: 43b98a2e
date: 2021-08-21 21:53:45
img:
coverImg:
password:
summary:
---

# php 抽奖算法（适合九宫格和大转盘）

```php
/* 
 * 不同概率的抽奖原理就是把0到*（比重总数）的区间分块
 * 分块的依据是物品占整个的比重，再根据随机数种子来产生0-* 中的某个数
 * 判断这个数是落在哪个区间上，区间对应的就是抽到的那个物品。
 * 随机数理论上是概率均等的，那么相应的区间所含数的多少就体现了抽奖物品概率的不同。
 */  


    /**
     * 抽奖方法
     * @return [array] [抽奖情况]
     */
    public function doDraw()
    {
        // 奖品数组
        $proArr = array(
            // id => 奖品等级， name => 奖品名称, v => 奖品权重
            array('id'=>1,'name'=>'超级奖品','v'=>0),
            array('id'=>2,'name'=>'特等奖','v'=>1),
            array('id'=>3,'name'=>'一等奖','v'=>5),
            array('id'=>4,'name'=>'二等奖','v'=>10),
            array('id'=>5,'name'=>'三等奖','v'=>12),
            array('id'=>6,'name'=>'四等奖','v'=>22),
            array('id'=>7,'name'=>'五等奖','v'=>50),
            array('id'=>8,'name'=>'六等奖','v'=>100),
            array('id'=>9,'name'=>'七等奖','v'=>200),
            array('id'=>10,'name'=>'八等奖','v'=>200),
            array('id'=>11,'name'=>'没中奖','v'=>500),
        );

        // 奖品等级奖品权重数组
        $arr = [];
        foreach ($proArr as $key => $val) {
            $arr[$val['id']] = $val['v'];
        }
        // 中奖 id
        $rid = $this->get_rand($arr);


        /**模拟抽奖测试**/
/*        $i = 0;
        while ( $i < 10000) {
          $rid = $this->get_rand($arr);
          $res[] = $rid;
          $i++;
        }
        // 统计奖品出现次数
        $result = array_count_values($res);
        asort($result);
        foreach ($result as $id => $times) {
            foreach ($proArr as $gifts) {
                if($id == $gifts['id']){
                    $response[$gifts['name']] = $times;
                }
            }
        }
        dump($response);
        die;*/


        $result = [];
        // 中奖礼品
        $result['yes'] = $proArr[$rid-1]['name'];
        // 从原奖品数组中剔除已经中奖礼品
        unset($proArr[$rid-1]);

        // 打乱数组排序
        shuffle($proArr);

        for ($i=0; $i < count($proArr); $i++) {
            $result['no'][] = $proArr[$i]['name'];
        }

        // foreach ($proArr as $k => $v) {
        //     // 没中奖礼品
        //     $result['no'][] = $v['name'];
        // }

        dump($result);

    }

    /**
     * 抽奖算法
     * @param  array  $proArr 奖品等级奖品权重数组
     * @return [int]         中奖奖品等级
     */
    public function get_rand($proArr = array()) {
        if(empty($proArr)) die;
        $rid = '';

        // 概率数组的总权重
        $proSum = array_sum($proArr);

        // 概率数组循环
        foreach ($proArr as $k => $proCur) {
            // 从 1 到概率总数中任意取值
            $randNum = mt_rand(1, $proSum);
            // 判断随机数是否在概率权重中
            if ($randNum <= $proCur) {
                // 取出奖品 id
                $rid = $k;
                break;
            } else {
                // 如果随机数不在概率权限中，则不断缩小总权重，直到从奖品数组中取出一个奖品
                $proSum -= $proCur;
            }
        }

        unset($proArr);
        return $rid;
    }

```

测试方法：

```php

    public function test(){

        $proArr = array(
            array('id'=>1,'name'=>'特等奖','v'=>1),
            array('id'=>2,'name'=>'一等奖','v'=>5),
            array('id'=>3,'name'=>'二等奖','v'=>10),
            array('id'=>4,'name'=>'三等奖','v'=>12),
            array('id'=>5,'name'=>'四等奖','v'=>22),
            array('id'=>6,'name'=>'没中奖','v'=>500)
        );

        $result = array();
        foreach ($proArr as $key => $val) {
            $arr[$key] = $val['v'];
        }
        // 概率数组的总权重
        $proSum = array_sum($arr);

        // 概率数组循环
        foreach ($arr as $k => $v) {

            // 从 1 到概率总数中任意取值
            $randNum = mt_rand(1, $proSum);
            $aa[$k] = $randNum . '+' . $v . '+' . $proSum;
            if ($randNum <= $v) {
                $result = $proArr[$k];
                // 找到符合条件的值就跳出 foreach 循环
                // dump($result);
                break;
            } else {
                $proSum = $proSum - $v;
                $bb[$k] = $randNum . '+' . $v . '+' . $proSum;
            }

        }

        dump($aa);
        dump($bb);
        // // dump($result);
        // // dump(__DIR__);
        // $path = __DIR__ . DS . 'log';
        // if(!is_dir($path)){
        //     mkdir($path);
        // }

        // $fileName = $path . DS . 'log.txt';
        // 创建文件和打开文件的函数都是 fopen
        // $cFile = fopen($fileName,'a+');
        // $a = json_encode($aa) . "\r\n";
        // $b = json_encode($bb) . "\r\n";
        // fwrite($cFile,$a);
        // fwrite($cFile,$b);
        // fclose($cFile);

        // 读文件
        // $lines = file($fileName);
        // foreach ($lines as $key => $value) {
        //     dump($value.'555555555');
        // }
        // dump($lines);

    }

```

以下代码也是意义程度上相同的代码，但是这种算法用的不多

```php

/**
我是d1-d2：0-1
我是d1-d2：1-51
我是d1-d2：51-56
我是d1-d2：56-156
我是d1-d2：156-166
我是d1-d2：166-166
我是d1-d2：166-666
我是d1-d2：666-688
我是d1-d2：688-700
我是d1-d2：700-900**/


function get_rand($proArr) 
{   
    $result = array();
    foreach ($proArr as $key => $val) { 
        $arr[$key] = $val['v']; 
    }  
    $proSum = array_sum($arr);      // 计算总权重
    $randNum = mt_rand(1, $proSum);
    $d1 = 0;
    $d2 = 0;
    for ($i=0; $i < count($arr); $i++)
    {
        $d2 += $arr[$i];
        if($i==0)
        {
            $d1 = 0;
        }
        else
        {
            $d1 += $arr[$i-1];
        }
        if($randNum >= $d1 && $randNum <= $d2)
        {
            $result = $proArr[$i];
        }
    }
    unset ($arr); 
    return $result;
}
```

开启百分百中奖模式

```php

<?php
	
    /**
    *  $prize_arr 参与抽奖人员数据
    *  id: 一般是成员ID
    *  name 姓名
    *  v   得奖概率
    ***/
    $prize_arr = array( 
        '0' => array('id'=>1,'name'=>'小王','v'=>1), 
        '1' => array('id'=>2,'name'=>'小李','v'=>5), 
        '2' => array('id'=>3,'name'=>'小张','v'=>10), 
        '3' => array('id'=>4,'name'=>'小二','v'=>12), 
        '4' => array('id'=>5,'name'=>'小菜','v'=>22), 
        '6' => array('id'=>6,'name'=>'小范','v'=>50), 
        '7' => array('id'=>7,'name'=>'小范01','v'=>50), 
        '8' => array('id'=>8,'name'=>'小范02','v'=>100), 
        '9' => array('id'=>9,'name'=>'小范03','v'=>50), 
        '10' => array('id'=>10,'name'=>'小范04','v'=>50), 
        '11' => array('id'=>11,'name'=>'小范05','v'=>50), 
        '12' => array('id'=>12,'name'=>'小范06','v'=>50), 
        '13' => array('id'=>13,'name'=>'小范07','v'=>50), 
        '14' => array('id'=>14,'name'=>'小范08','v'=>50), 
        '15' => array('id'=>15,'name'=>'小范09','v'=>100), 
        '16' => array('id'=>16,'name'=>'小范10','v'=>100), 
    );  

    foreach ($prize_arr as $key => $val) { 
        $arr[$key] = $val['v']; 
    } 
    
    $total_num = '8'; //设定得中奖人数量
    
    $temp_rest=array();
    for($i=0;$i<$total_num;$i++)
    {
        $rid = get_rand($arr,true); //根据概率获取人员ID
        $temp_rest[]= $prize_arr[$rid]; //中奖项
        unset($prize_arr[$rid]); 
        unset($arr[$rid]); 
    }

    print_r($temp_rest);//得出结果

    /****
    *   得出当前中奖人
    *   $is_status是否开启概率为100必中: 默认不开启 
    ***/
    function get_rand($proArr,$is_status = false) { 
        $result = ''; 
        if($is_status){
            $rest = get_100($proArr);  //调用获取100命中
        }else{ 
            $rest ='';
        }

        if(empty($rest) || !isset($rest)){
            //概率数组的总概率精度
            $proSum = array_sum($proArr); 
            //概率数组循环
            foreach ($proArr as $key => $proCur) { 
                $randNum = mt_rand(1, $proSum); 
                if ($randNum <= $proCur) { 
                    $result = $key; 
                    break; 
                } else { 
                    $proSum -= $proCur; 
                }   
            } 
        }else{
            $result = $rest;
        }
        unset ($proArr); 
        return $result; 
    }
    function get_100($arr_mast){
        $result = ''; 
        foreach ($arr_mast as $key => $value_mast) { 
           if($value_mast== 100){
                $result = $key; 
                break; 
           }
          
        } 
        unset ($arr_mast); 
        return $result; 
    }
```
