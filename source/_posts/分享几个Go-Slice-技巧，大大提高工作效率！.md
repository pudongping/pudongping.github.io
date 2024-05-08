---
title: 分享几个Go Slice 技巧，大大提高工作效率！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Go
tags:
  - Go
  - Golang
abbrlink: 669a361f
date: 2024-05-09 01:10:56
img:
coverImg:
password:
summary:
---

## 声明一个数组

```go

var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 让编译器统计数组字面值中元素的数目
var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 类型： [10]int 长度： 10 容量： 10 值为： [1 2 3 4 5 6 7 8 9 10]
fmt.Printf("类型： %T 长度： %v 容量： %v 值为： %v \n", arr, len(arr), cap(arr), arr)
// 比较，类型： bool 结果： true 
fmt.Printf("比较，类型： %T 结果： %v \n", arr == [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, arr == [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
// 比较，类型： bool 结果： true 
fmt.Printf("比较，类型： %T 结果： %v \n", arr == arr1, arr == arr1)

```

## 基于数组创建一个切片

```go

var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 基于数组创建一个切片
var slice = arr[:]

// 类型： [10]int 长度： 10 容量： 10 值为： [1 2 3 4 5 6 7 8 9 10]
fmt.Printf("类型： %T 长度： %v 容量： %v 值为： %v \n", arr, len(arr), cap(arr), arr)
// 类型： []int 长度： 10 容量： 10 值为： [1 2 3 4 5 6 7 8 9 10] 
fmt.Printf("类型： %T 长度： %v 容量： %v 值为： %v \n", slice, len(slice), cap(slice), slice)
// 类型不同，无法运算，编译器直接编译不过
// fmt.Printf("比较，类型： %T 结果： %v \n", arr == slice, arr == slice)

```

## 切片截取

```go

var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 左闭右开的索引区间
// slice2 的长度为 3 即（6-3=3）容量为 7 即（10-3=7）
var slice2 = slice1[3:6]
// slice3 的长度为 3 即（6-3=3）容量为 6 即（9-3=6）
slice3 := slice1[3:6:9]

// 类型： []int 长度： 10 容量： 10 值为： [1 2 3 4 5 6 7 8 9 10]
fmt.Printf("类型： %T 长度： %v 容量： %v 值为： %v \n", slice1, len(slice1), cap(slice1), slice1)
// 类型： []int 长度： 3 容量： 7 值为： [4 5 6]
fmt.Printf("类型： %T 长度： %v 容量： %v 值为： %v \n", slice2, len(slice2), cap(slice2), slice2)
// 类型： []int 长度： 3 容量： 6 值为： [4 5 6]
fmt.Printf("类型： %T 长度： %v 容量： %v 值为： %v \n", slice3, len(slice3), cap(slice3), slice3)

slice4 := slice2[:cap(slice2)]
slice5 := slice3[:cap(slice3)]

// slice4 ==> [4 5 6 7 8 9 10]
// slice2 实际引用的数组是 [4 5 6 7 8 9 10]
fmt.Printf("slice4 ==> %v \n", slice4)
// slice5 ==> [4 5 6 7 8 9]
// slice3 实际引用的数组是 [4 5 6 7 8 9]
fmt.Printf("slice5 ==> %v \n", slice5)

```

## 切片拷贝

```go

var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
slice2 := slice1[4:7]
// slice1 ==> [1 2 3 4 5 6 7 8 9 10]
fmt.Printf("slice1 ==> %v \n", slice1)
// slice2 ==> [5 6 7]
fmt.Printf("slice2 ==> %v \n", slice2)

// 修改切割后的切片值，发现原始切片也被修改了（浅拷贝，切片为引用传递）
slice2[1] = 123
// slice1 ==> [1 2 3 4 5 123 7 8 9 10]
fmt.Printf("slice1 ==> %v \n", slice1)
// slice2 ==> [5 123 7]
fmt.Printf("slice2 ==> %v \n", slice2)

// 复制一个切片时，如果不申明长度，是不会复制的
var slice3 []int
copy(slice3, slice2)
// slice1 ==> [1 2 3 4 5 123 7 8 9 10]
fmt.Printf("slice1 ==> %v \n", slice1)
// slice2 ==> [5 123 7]
fmt.Printf("slice2 ==> %v \n", slice2)
// slice3 ==> []
fmt.Printf("slice3 ==> %v \n", slice3)

slice4 := make([]int, len(slice2))
copy(slice4, slice2)
// 通过复制后，修改复制后的切片不会影响原切片（深拷贝，值传递）
slice4[1] = 567
// slice1 ==> [1 2 3 4 5 123 7 8 9 10]
fmt.Printf("slice1 ==> %v \n", slice1)
// slice2 ==> [5 123 7]
fmt.Printf("slice2 ==> %v \n", slice2)
// slice4 ==> [5 567 7] 
fmt.Printf("slice4 ==> %v \n", slice4)

```

## ArrayChunk 数组切割

```go

// ArrayChunkString 将一个数组分割成多个
// s := []string{"a1", "a2", "a3", "a4", "a5", "a6", "a7"}
// size := 2
// output: [[a1 a2] [a3 a4] [a5 a6] [a7]]
func ArrayChunkString(s []string, size int) [][]string {
	if size < 1 {
		panic("size: cannot be less than 1")
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]string
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, s[i*size:end])
		i++
	}
	return n
}

// ArrayChunkInt 将一个切片分割成多个
// i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// size := 3
// output: [[0 1 2] [3 4 5] [6 7 8] [9]]
func ArrayChunkInt(i []int, size int) [][]int {
	batches := make([][]int, 0, (len(i)+size-1)/size)

	for size < len(i) {
		i, batches = i[size:], append(batches, i[0:size:size])
	}
	batches = append(batches, i)

	return batches
}

```

## ArrayReverse 数组反转

```go

// ArrayReverseInt 数组反转
// i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// output: [9 8 7 6 5 4 3 2 1 0]
func ArrayReverseInt(i []int) []int {
	for left, right := 0, len(i)-1; left < right; left, right = left+1, right-1 {
		i[left], i[right] = i[right], i[left]
	}
	return i
}

```
