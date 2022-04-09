---
title: 支付宝花呗分期费用计算 PHP 示例代码
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Pay
tags:
  - 支付宝花呗分期
  - 支付
abbrlink: c63afd43
date: 2022-01-07 17:38:47
img:
coverImg:
password:
summary:
---

# 支付宝花呗分期费用计算 PHP 示例代码

## 支付宝-花呗分期-计费结果示例

> 当天汇率为：汇率 1 日元 = 0.05548 人民币。


商户承担手续费花呗分期费率

期数 | 费率
--- | ---
3 | 1.80%
6 | 4.5%
12 | 7.5%


用户承担手续费花呗分期费率

期数 | 费率
--- | ---
3 | 2.30%
6 | 4.50%
12 | 7.50%


计算结果示例如下：

> 以下计算结果示例以用户承担手续费花呗分期费率计算得出。计算时间为：2022 年 1 月 6 日。

期数 | 本金（人民币）| 本金（日元） | 每月应还款金额（人民币，单位：元） | 总手续费（人民币，单位：元） |  折算年化率（单利）约为
--- | --- | --- | --- | --- | --- 
3 | 5.45 | 98.00 | 1.85 | 0.13 | 13.7%
3 | 3.45 | 62.00 | 1.17 | 0.08 | 13.7%
6 | 3.45 | 62.00 | 0.59 | 0.16 | 15.3%
12 | 3.45 | 62.00 | 0.30 | 0.26 | 13.6%
3 | 1111.11 | 20027.00 | 378.89 | 25.56 | 13.7%
6 | 1111.11 | 20027.00 | 193.51 | 50.00 | 15.3%
12 | 1111.11 | 20027.00 | 99.53 | 83.33 | 13.6%
3 | 123.67 | 2229.00 | 42.16 | 2.84 | 13.7%
6 | 123.67 | 2229.00 | 21.53 | 5.57 | 15.3%
12 | 123.67 | 2229.00 | 11.07 | 9.28 | 13.6%

## 支付宝花呗分期费用计算 PHP 示例代码

> 以下代码摘抄自 [pudongping/global-pay 跨境支付宝 PHP 支付插件包](https://github.com/pudongping/global-pay)

```php

<?php
/**
 * document link： https://opendocs.alipay.com/mini/introduce/antcreditpay-istallment
 * document link： https://opendocs.alipay.com/open/277/105952
 *
 * Created by PhpStorm
 * User: Alex
 * Date: 2021-08-29 16:03
 * E-mail: <276558492@qq.com>
 */
declare(strict_types=1);

class HbFqCost
{

    const USER_ASSUME = 0;  // 用户承担手续费
    const SELLER_ASSUME = 100;  // 商家承担手续费
    const VALIDATE_NPER = [3, 6, 12];  // 花呗分期合法的分期数

    public static $rate = [
        self::USER_ASSUME => [
            3 => 0.023,
            6 => 0.045,
            12 => 0.075
        ],
        self::SELLER_ASSUME => [
            3 => 0.018,
            6 => 0.045,
            12 => 0.075
        ],
    ];

    /**
     * 获取花呗分期计费情况
     *
     * @param float $totalAmount  本金
     * @param bool $isShowAll  是否显示每一期的还款数
     * @param bool $isSellerPercent  是否商家承担所有的手续费
     * @return array
     */
    public function fetchHbFqCost(float $totalAmount, bool $isShowAll = false, bool $isSellerPercent = false): array
    {
        $assume = $isSellerPercent ? self::SELLER_ASSUME : self::USER_ASSUME;
        $rates = self::$rate[$assume];
        $data = [];

        foreach ($rates as $nper => $rate) {
            $data[] = $this->calHbFqCost($nper, $rate, $totalAmount, $isShowAll);
        }

        return $data;
    }

    /**
     * 计算花呗分期手续费
     *
     * @param int $nper  期数
     * @param float $rate  费率
     * @param float $totalAmount  本金
     * @param bool $showAll  是否显示每一期的还款数
     * @return array
     */
    public function calHbFqCost(int $nper, float $rate, float $totalAmount, bool $showAll = false)
    {
        $totalAmountCent = bcmul((string)$totalAmount, '100', 4);  // 1. 把金额单位转化成分 cent
        // 用户每期本金
        $perAmount = floor(floatval(bcdiv($totalAmountCent, (string)$nper, 4)));  // 2. 计算每期本金（用总金额/总期数，结果以分表示，向下取整）

        // 用户每期手续费
        $buyerTotalCost = (float)bcmul($totalAmountCent, (string)$rate, 4);  //  2. 用转化为分后的金额乘以买家费率，得到以分表示的买家总费用（总手续费）
        $roundTotalCost = round($buyerTotalCost, 0, PHP_ROUND_HALF_EVEN);  // 3. 对费用进行取整（取整规则为 ROUND_HALF_EVEN ）
        $perCharge = floor(floatval(bcdiv((string)$roundTotalCost, (string)$nper, 4)));  // 4. 计算每期费用（用总费用/总期数，结果以分表示，向下取整）

        // 用户每期总费用
        $perTotalAmount = bcadd((string)$perAmount, (string)$perCharge);

        // 金额以 [元] 为单位
        $perAmountYuan = floatval(bcdiv((string)$perAmount, '100', 2));
        $perChargeYuan = floatval(bcdiv((string)$perCharge, '100', 2));
        $perTotalAmountYuan = floatval(bcdiv((string)$perTotalAmount, '100', 2));
        $buyerTotalCostYuan = round(floatval(bcdiv((string)$buyerTotalCost, '100', 4)), 2);  // 花呗分期的总手续费实行“四舍五入”的原则进行计算

        $ret = [
            'nper' => $nper,  // 期数
            'total_amount' => $totalAmount,  // 本金
            'total_charge' => $buyerTotalCostYuan,  // 总手续费
            'rate' => $rate,  // 利率
            'per_charge' => $perChargeYuan,  // 每期手续费
            'per_amount' => $perAmountYuan,  // 每期本金
            'per_total_amount' => $perTotalAmountYuan,  // 每期总费用
            'refund_list' => [],  // 还款列表
        ];

        if ($showAll) {
            $ret['refund_list'] = $this->getRefundList($ret);
        }

        return $ret;
    }

    /**
     * 获取还款的列表
     *
     * @param array $params
     * @return array
     */
    public function getRefundList(array $params)
    {
        $nper = $params['nper'];  // 期数

        $data = [];
        for ($i = 1; $i <= $nper; $i++) {
            $item = [];
            $item['nper'] = $i;  // 第几期
            $item['charge'] = $params['per_charge'];  // 当前期数所需要支付的手续费
            $item['amount'] = $params['per_amount'];  // 当前期数所需要支付的本金数
            $item['current_total_amount'] = $params['per_total_amount'];  // 当前期数所需要支付的总费用
            $data[] = $item;
        }

        $charges = array_column($data, 'charge');
        // 计算的所有手续费总和
        $chargesSum = array_reduce($charges, function ($carry, $item) {
            return bcadd((string)$carry, (string)$item, 2);
        }, '0');

        $amounts = array_column($data, 'amount');
        // 计算的所有本金总和
        $amountsSum = array_reduce($amounts, function ($carry, $item) {
            return bcadd((string)$carry, (string)$item, 2);
        }, '0');

        // 如果所需支付的总手续费大于计算后的手续费总和，那么则需要将缺少的手续费补加到第一期
        if ($params['total_charge'] > (float)$chargesSum) {
            $data[0]['charge'] = floatval(bcadd(strval($data[0]['charge']), bcsub(strval($params['total_charge']), strval($chargesSum), 2), 2));
        }
        // 如果所需要支付的本金大于计算后的本金总和，那么则需要将缺少的本金补加到第一期
        if ($params['total_amount'] > (float)$amountsSum) {
            $data[0]['amount'] = floatval(bcadd(strval($data[0]['amount']), bcsub(strval($params['total_amount']), strval($amountsSum), 2), 2));
        }
        // 第一期所需要支付的总金额
        $data[0]['current_total_amount'] = floatval(bcadd(strval($data[0]['charge']), strval($data[0]['amount']), 2));

        return $data;
    }

}

```

## 使用

```php

$istallment = new HbFqCost();

$result = $istallment->fetchHbFqCost(123.67);

var_dump($result);

```
