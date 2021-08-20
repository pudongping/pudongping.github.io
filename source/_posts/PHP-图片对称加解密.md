---
title: PHP 图片对称加解密
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - 图片加解密
abbrlink: 8331d4d3
date: 2021-08-20 20:31:52
img:
coverImg:
password:
summary:
---

# PHP 图片加解密

> 可以将人员身份证图片通过修改字节加密，并且可将身份证信息也写入图片中

```php

<?php

/***
 *                    .::::.
 *                  .::::::::.
 *                 :::::::::::  FUCK YOU
 *             ..:::::::::::'
 *           '::::::::::::'
 *             .::::::::::
 *        '::::::::::::::..
 *             ..::::::::::::.
 *           ``::::::::::::::::
 *            ::::``:::::::::'        .:::.
 *           ::::'   ':::::'       .::::::::.
 *         .::::'      ::::     .:::::::'::::.
 *        .:::'       :::::  .:::::::::' ':::::.
 *       .::'        :::::.:::::::::'      ':::::.
 *      .::'         ::::::::::::::'         ``::::.
 *  ...:::           ::::::::::::'              ``::.
 * ```` ':.          ':::::::::'                  ::::..
 *                    '.:::::'                    ':'````..
 */

class Encrypt 
{
	
	/**
	 * 图片对称加密
	 *
	 * @param [string] $filePath 图片路径
	 * @return void
	 */
	public function enc($filePath)
	{
		// 文档中建议：为移植性考虑，强烈建议在用 fopen() 打开文件时总是使用 'b' 标记。
		$fileId = fopen($filePath, 'rb+');

		// 取出文件大小的字节数 （29124）
		$fileSize = fileSize($filePath);

		// 读取文件，返回所读取的字符串 （读出来的为二进制序列）
		$img = fread($fileId, $fileSize);

		// 使用“无符号字符”，从二进制字符串对数据进行解包
		// （pack、unpack用法）https://segmentfault.com/a/1190000008305573
		$imgUnpack = unpack('C*', $img); // $fileSize 长度的一维数组 [ 1=>255, 2=>216, 3=>255, ……, 29124=>217 ]
		
		// 关闭一个已打开的文件指针		
		fclose($fileId);

		$tempArr = [];
		// 自定义加密规则
		for ($i = 1; $i <= $fileSize; $i++) { 
			$value = 0;
			if ($i % 3 == 0) {
				$value = 2;
			} elseif ($i % 5 == 0) {
				$value = 4;
			} elseif ($i % 7 == 0) {
				$value = 6;
			}
			$byte = $imgUnpack[$i];	// 图片原始字节
			$byte = $byte + $value; // 经过加密规则之后的字节
			// 打包成二进制字符串
			$tempArr[] = pack('C*', $byte);
		}

		$img = implode('', $tempArr);	// 将解包之后的一维数组装换成字符串
		file_put_contents($filePath, $img); // 重写图片
	}


	/**
	 * 图片对称解密
	 *
	 * @param [string] $filePath	图片路径
	 * @return void
	 */
	public function dec($filePath)
	{
		$fileId = fopen($filePath, 'rb+');
		$fileSize = filesize($filePath);
		$img = fread($fileId, $fileSize);
		$imgUnpack = unpack('C*', $img);
		fclose($fileId);

		$tempArr = [];
		// 开始解密
		for ($i = 1; $i <= $fileSize; $i++) { 
			$value = 0;
			if ($i % 3 == 0) {
				$value = 2;
			} elseif ($i % 5 == 0) {
				$value = 4;
			} elseif ($i % 7 == 0) {
				$value = 6;
			}
			$byte = $imgUnpack[$i];
			$byte = $byte - $value;
			$tempArr[] = pack('C*', $byte);
		}
		$img = implode('', $tempArr);
		file_put_contents($filePath, $img);
	}


	/**
	 * 图片追加信息
	 *
	 * @param [string] $filePath	图片路径
	 * @param [array] $cardmsg	需要添加的信息数组
	 * @param [array] $separate	分隔数组（类似于做一个加密分隔 key）
	 * @return void
	 */
	public function encmsg($filePath, $cardmsg, $separate)
	{
		// 文档中建议：为移植性考虑，强烈建议在用 fopen() 打开文件时总是使用 'b' 标记。
		$fileId = fopen($filePath, 'rb+');
		// 取出文件大小的字节数 （29124）
		$fileSize = fileSize($filePath);
		// 读取文件，返回所读取的字符串 （读出来的为二进制序列）
		$img = fread($fileId, $fileSize);
		// 使用“无符号字符”，从二进制字符串对数据进行解包
		// （pack、unpack用法）https://segmentfault.com/a/1190000008305573
		$imgUnpack = unpack('C*', $img); // $fileSize 长度的一维数组 [ 1=>255, 2=>216, 3=>255, ……, 29124=>217 ]
		// 关闭一个已打开的文件指针		
		fclose($fileId);

		// 处理身份信息
		$cardmsgJson = json_encode($cardmsg, JSON_UNESCAPED_UNICODE);
		$cardmsgUnpack = unpack('C*', $cardmsgJson);

		// 合并图片字节、自定义分隔数组（类似手动加 key 值）、身份信息字节
		$mergeArr = array_merge($imgUnpack, $separate, $cardmsgUnpack);

		$pack = [];
		foreach ($mergeArr as $k => $v) {
			$pack[] = pack('C*', $v);
		}
		$packStr = join('', $pack);
		file_put_contents($filePath, $packStr); // 重写图片
	}


	/**
	 * 获取追加进图片的信息
	 *
	 * @param [string] $filePath	图片路径
	 * @param [array] $separate	定义的分隔数组（分隔 key）
	 * @return [string] 追加进的图片信息
	 */
	public function decmsg ($filePath, $separate) 
	{
		// 文档中建议：为移植性考虑，强烈建议在用 fopen() 打开文件时总是使用 'b' 标记。
		$fileId = fopen($filePath, 'rb+');
		// 取出文件大小的字节数 (29192)
		$fileSize = fileSize($filePath);
		// 读取文件，返回所读取的字符串 （读出来的为二进制序列）
		$img = fread($fileId, $fileSize);

		// 使用“无符号字符”，从二进制字符串对数据进行解包
		$imgUnpack = unpack('C*', $img); // $fileSize 长度的一维数组 [ 1=>255, 2=>216, 3=>255, ……, 29192=>217 ]
		// 关闭一个已打开的文件指针		
		fclose($fileId);

		$imgUnpackStr = join(',',$imgUnpack); // 将一维数组转换为字符串
		$separateStr = implode(',', $separate); // 将一维数组转换为字符串
		$imgAndCardmsgArr = explode($separateStr, $imgUnpackStr); // 以自定义分隔符分隔出图片字节和身份信息字节
		
		$cardmsgArr = explode(',', $imgAndCardmsgArr[1]); // 取出身份信息字节
		unset($cardmsgArr[0]); // 去除身份信息字节首位空白 （字符串转数组时所留）
		$cardmsg = '';
		foreach ($cardmsgArr as $k => $v) {
			$cardmsg .= pack('C*', $v);	// 打包成二进制文件字符串
		}
 
		return json_decode($cardmsg, true);
	}



}


$encrypt = new Encrypt();

$path = './001.jpg';

$separate = [255, 0, 255, 0, 255, 0, 255, 206, 210, 202, 199, 183, 214, 184, 244]; // 15字节
$cardmsg = ['name' => '张三', 'gender' => '男', 'idcard' => 12345678910]; // 53字节






```
