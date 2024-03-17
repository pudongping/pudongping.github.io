---
title: 一文弄懂 Swoole 进程、线程、协程到底是什么？
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - Swoole
abbrlink: 289a4dbc
date: 2024-03-17 14:10:22
img:
coverImg:
password:
summary:
---

## swoole 和 workerman 的区别
- workerman 使用纯 php 编写；swoole 采用 c/c++ 语言编写，作为 php 扩展
- workerman 多进程；swoole 协程，多进程，多线程
- swoole 性能优于 workerman
- swoole 并没有用 libevent，所以不需要安装 libevent
- swoole 并不依赖 php 的 stream/sockets/pcntl/posix/sysvmsg 等扩展

## master （主进程）、reactor （线程）、manager（管理进程）、worker （工作进程） 、 task （任务进程）

对于 Linux 系统，是无法区分**进程**和**线程**的，CPU 是单进程（因此**并行执行**只有1，因为只有 1 个 CPU，但是可用开多个进程/线程去执行，这样就达到了**并发执行**，CPU 可以通过轮训调度从多个进程/线程之间来回做切换，但是多进程/线程切换也就产生了切换之间的成本**进程/线程的数量越多，切换成本就越大，也就越浪费，因为这种切换并没有在业务代码中消耗**）

一个线程，又可用分为用户空间（用户线程）和内核空间（内核线程），内核线程就是 thread（线程），用户线程就是 co-routine（协程，golang 中叫做 goroutine），然后又可以通过一个线程去绑定一个 **协程调度器**，一个协程调度器又可以绑定多个协程，这样性能高低就完全取决于优化协程调度器，哪门语言的协程调度器做的好，那么性能就高。

![image.png](https://upload-images.jianshu.io/upload_images/14623749-e8dbb161c619f9c2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-877e8c01041475e3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### master 进程
主进程内有多个 reactor 线程，reactor 线程基于 epoll/kqueue 进行网络事件轮训。收到数据后转发到 worker 进程去处理。

### reactor 线程
reactor 线程个数默认和 cpu 核心数一致

### manager 进程
对所有 worker 进程进行管理，worker 进程生命周期结束或者发生异常时自动回收，并创建新的 worker 进程

### worker 进程
对收到的数据进行处理，包括协议解析和响应请求。如果 swoole 没有设置 worker_num ，底层会启动与 cpu 数量一致的 worker 进程

### task 进程
task 进程和 worker 进程是同级别的，可以将 worker 进程中的一些服务投递给 task 进程进行处理，来达到分担 worker 进程工作量的目的

### coroutine 协程
协程可以理解为纯用户态的线程，其通过协作而不是抢占来进行切换，相对于进程或者线程，协程所有的操作都可以在用户态完成，创建和切换的消耗更低，Swoole 可以为每一个请求创建对应的协程，根据 IO 的状态来合理的调度协程。

在 Swoole 4.x 中，协程（Coroutine）取代了异步回调，成为 Swoole 官方推荐的编程方式。Swoole 协程解决了异步回调编程困难的问题，使用协程可以以传统同步编程的方法编写代码，底层自动切换为异步 IO，既保证了编程的简单性，又可借助异步 IO，提升系统的并发能力。

> 注：Swoole 4.x 之前的版本也支持协程，不过 4.x 版本对协程内核进行了重构，功能更加强大，提供了完整的协程+通道特性，带来全新的 CSP 编程模型，后续介绍和示例都是基于 Swoole 4.x 版本。

![进程/线程结构图](https://upload-images.jianshu.io/upload_images/14623749-78ba06649d969718.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![swoole 架构](https://upload-images.jianshu.io/upload_images/14623749-aa2ab771fa24e858.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![swoole 运行流程图](https://upload-images.jianshu.io/upload_images/14623749-e90a521635fbfee1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## 进程和线程
- 进程  
  进程是一个具有一定独立功能的程序在一个数据集上的一次动态执行的过程，是操作系统进行资源分配和调度的一个独立单位，是应用程序运行的载体。
- 线程  
  线程是程序执行中一个单一的顺序控制流程，是程序执行流的最小单元，是处理器调度和分派的基本单位。一个进程可以有一个或多个线程。

### 进程和线程的区别
- 线程是程序执行的最小单位，而进程是操作系统分配资源的最小单位。
- 一个进程由一个或多个线程组成，线程是一个进程中代码的不同执行路线。
- 进程之间相互独立，但同一进程下的各个线程之间共享程序的内存空间（包括代码段，数据集，堆等）及一些进程级的资源（如打开文件和信号等），某进程内的线程在其他进程不可见。
- 调度和切换：线程上下文切换比进程上下文切换要快得多。
- 线程天生的共享内存空间，线程间的通信更简单，避免了进程 IPC（进程间的通信） 引入新的复杂度。
- 进程开销大，线程开销小。

### php 实现多进程

pcntl 是 php 官方的多进程扩展，只能在 Linux 环境下使用，编译 php 的时候需要开启 --enable-pcntl   
pcntl_fork 在当前进程当前位置产生分支（子进程），fork 是创建了一个子进程，父进程和子进程都从 fork 得位置开始向下继续执行，不同的是父进程执行过程中，得到的 fork 返回值为子进程号，而子进程得到得是 0。

```php

<?php

$pid = pcntl_fork();

if ($pid > 0) {
    echo '我是父进程' . PHP_EOL; 
} elseif ($pid == 0) {
    echo '我是子进程' . PHP_EOL;
}
```

多进程例子：会先输出 `111` 然后输出 `222`

```php

<?php

if ($pid = pcntl_fork() == 0) {
    sleep(1);
    echo '222';
}

if ($pid == 0) {
    if (pcntl_fork() == 0) {
        echo '111';
    }
}

// 两段代码并行执行，并不会因为第一个判断中有了 sleep 就停止往下执行代码

```

### 进程间的通信
在各个进程中，内存空间都是不一致的，各个变量都是在不同的内存空间。进程间空间独立，数据不能共享。

例子如下：

```php

<?php

$str = 'alex';
$pid = pcntl_fork();

if ($pid > 0) {
    $str .= '===111';
    echo $str . PHP_EOL;  // 会输出 alex===111
} else {
    echo $str . PHP_EOL;  // 会输出 alex
}

```

进程间的通信可以采用多种方式，比如：管道通信、消息队列通信、进程信号通信、套接字通信、第三方通信（使用文件操作，或者mysql，或者redis等方法也可以实现通信）、共享内存通信（映射一段能被其他进程所访问的内存，这段共享内存由一个进程创建，但多个进程都可以访问。共享内存是最快的 IPC <进程间通信> 方式，它是针对其他进程间通信方式运行效率低而专门设计的。`swoole` 就是采用共享内存通信）

## 线程和协程

提起并发编程，最常见的就是 **多线程编程** 。线程是操作系统能够进行调度的最小单位，共享同一进程的数据和资源，并行地处理多个任务。（同一进程下多线程之间的内存是共享的）。

多线程存在两个问题，在线程数量过多时，问题被放大的尤为明显。

- 线程的上下文切换造成的开销。
- 线程之间对资源的竞争问题（线程安全）。

协程也是一种异步方案。在代码 IO 阻塞时，当前协程让出 CPU 执行权，让其它协程执行，待 IO 执行完毕，阻塞的协程继续执行。虽然代码是异步执行，但写代码看起来像是同步的。支持协程的编程语言实现了协程的调度器，提供了 channel 机制进行协程间通信 （CSP模型中消息传递的实现）。基于 CSP 模型的协程方案，实现了无共享内存无锁的并发，可以匹配异步回调的性能。

**Swoole 的协程在底层实现上是单线程的**，因此同一时间只有一个协程在工作，协程的执行是串行的，这与线程不同，多个线程会被操作系统调度到多个 CPU 并行执行。

一个协程正在运行时，其他协程会停止工作。当前协程执行阻塞 IO 操作时会挂起，底层调度器会进入事件循环。当有 IO 完成事件时，底层调度器恢复事件对应的协程的执行。

在 Swoole 中对 CPU 多核的利用，仍然依赖于 Swoole 引擎的多进程机制。



### 协程的特性

协程极高的执行效率。因为子程序切换不是线程切换，而是由程序自身控制，因此，没有线程切换的开销。  
不需要多进程的锁机制，因为只有一个线程，也不存在同时写变量冲突。（swoole 是开了一个线程，多个协程的机制）

### 线程和协程的适用场景

协程（协同程序），同一时间只能执行某个协程。开辟多个协程开销不大，协程适合对某任务进行分时处理。**IO密集型**   
线程，同一时间可以同时执行多个线程。开辟多条线程开销很大。线程适合多任务同时处理。 **CPU密集型**
