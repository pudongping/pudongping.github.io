---
title: Laravel 中的模型事件与 Observer
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
abbrlink: c9de0987
date: 2021-08-25 20:19:08
img:
coverImg:
password:
summary:
---


# Laravel 中的模型事件与 Observer

当模型已存在，不是新建的时候，依次触发的顺序是:

saving -> updating -> updated -> saved

当模型不存在，需要新增的时候，依次触发的顺序则是

saving -> creating -> created -> saved

那么 saving,saved 和 updating,updated 到底有什么区别呢？

updating 和 updated 会在数据库中的真值修改前后触发。

saving 和 saved 则会在 Eloquent 实例的 original 数组真值更改前后触发。


```php

class UserObserver
{

    /**
     * 监听数据即将创建的事件。
     *
     * @param  User $user
     * @return void
     */
    public function creating(User $user)
    {

    }

    /**
     * 监听数据创建后的事件。
     *
     * @param  User $user
     * @return void
     */
    public function created(User $user)
    {

    }

    /**
     * 监听数据即将更新的事件。
     *
     * @param  User $user
     * @return void
     */
    public function updating(User $user)
    {

    }

    /**
     * 监听数据更新后的事件。
     *
     * @param  User $user
     * @return void
     */
    public function updated(User $user)
    {

    }

    /**
     * 监听数据即将保存的事件。
     *
     * @param  User $user
     * @return void
     */
    public function saving(User $user)
    {

    }

    /**
     * 监听数据保存后的事件。
     *
     * @param  User $user
     * @return void
     */
    public function saved(User $user)
    {

    }

    /**
     * 监听数据即将删除的事件。
     *
     * @param  User $user
     * @return void
     */
    public function deleting(User $user)
    {

    }

    /**
     * 监听数据删除后的事件。
     *
     * @param  User $user
     * @return void
     */
    public function deleted(User $user)
    {

    }

    /**
     * 监听数据即将从软删除状态恢复的事件。
     *
     * @param  User $user
     * @return void
     */
    public function restoring(User $user)
    {

    }

    /**
     * 监听数据从软删除状态恢复后的事件。
     *
     * @param  User $user
     * @return void
     */
    public function restored(User $user)
    {

    }
}

```
