---
title: PHP中常用的设计模式
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - 设计模式
abbrlink: e0806fd7
date: 2024-06-26 11:30:42
img:
coverImg:
password:
summary:
---


1. 创建型模式  
   单例模式、工厂模式、简单工厂模式、抽象工厂模式、建造者模式、原型模式
2. 结构型模式    
   适配器模式、组合模式、代理模式、外观模式、装饰器模式、桥接模式、享元模式
3. 行为型模式  
   观察者模式、迭代子模式、策略模式、命令模式、模板方法模式、责任链模式、备忘录模式、状态模式、访问者模式、中介者模式、注册模式

### 单例模式
- 什么是单例模式？  
  单例模式通俗定义，一个类只有一个实例。而且是内部自行实例化并向整个系统全局地提供这个实例。它不会创建实例副本，而是返回单例类内部存储的实例一个引用。

> 上面的代码用静态变量 实现了单例模式和创建单例的静态方法 getInstance(). 请注意以下几点：
>
> - 构造函数 __construct() 被声明为 protected 是为了防止用 new 操作符在这个类之外创建新的实例。
> - 魔术方法 __clone() 被声明为 private 是为了防止用 clone 操作符克隆出新的实例.
> - 魔术方法 __wakeup() 被声明为 private 是为了防止通过全局函数 unserialize() 反序列化这个类的实例。
> - 新的实例是用过静态方法 getInstance() 使用后期静态绑定生成的。这允许我们对 Singleton 类进行继承，并且在取得 SingletonChild 的单例时不会出现问题。    
    单例模式是非常有用的，特别是我们需要确保在整个请求的声明周期内只有一个实例存在。典型的应用场景是，当我们有一个全局的对象（比如配置类）或一个共享的资源（比如事件队列）时。
>
> 你应该非常小心地使用单例模式，因为它非常自然地引入了全局状态到你的应用中，降低了可测试性。 在大多数情况下，依赖注入可以（并且应该）代替单例类。 使用依赖注入意味着我们不会在设计应用时引入不必要的耦合，因为对象使用共享的或全局的资源，不再需要耦合具体的类。

```php
<?php
/**
 *  我们如何来写一个属于自己的单例类呢？
 *  写一个单例类是否有什么规则可寻？
 *  1.有一个静态成员变量来保存类的唯一实例
 *  2.构造函数和克隆函数必须申明为私有的（防止外部程序能通过 new 关键字进行创建对象）
 *  3.公共的静态方法可以访问类的实例作为一个引用返回
 *  4.只能实例化一次
 */


class Obj
{

    protected static $_Ins;

    public $str = null;

    private function __construct()
    {
        $this->str = mt_rand(100,999);
    }

    private function __clone(){

    }

    public static function getInstance(){
        if (!(self::$_Ins instanceof self)){
            self::$_Ins = new self();
        }
        return self::$_Ins;
    }


}

$a = Obj::getInstance();
$b = Obj::getInstance();

var_dump($a);
echo '<br/><hr/>';
var_dump($b);
echo '<br/><hr/>';
var_dump($a === $b); // true


或者采用 ci 框架比较古老的单例方法

<?php

class ClassName
{
	
	private static $instance;

	function __construct()
	{
		// 运用了 $this 因此下面调用的时候必须先实例化
		self::$instance =& $this;
	}

	public static function &get_instance()
	{
		return self::$instance;
	}

}

$aa = new ClassName;
$bb = ClassName::get_instance();
$cc = ClassName::get_instance();

echo '<pre>';
var_dump($aa);

echo '<pre>';
var_dump($bb);

echo '<pre>';
var_dump($cc);
```

### 策略模式
- 什么是策略模式？  
  定义了算法家族，分别封装起来，让它们之间可以相互替换，此模式让算法的变化不会影响到使用算法的用户。使用策略模式可以实现 Ioc、依赖倒置、控制反转
```php
<?php
/**
 * 策略模式
 * 需求：同样一份数据需要导出不同的格式
 */

class DataModel {
    private $writer;
    public function __construct($writer){
        $this->writer = $writer;
    }

    public function export(){
        $data = [
            [1, 'This is first'],
            [2, 'This is second'],
            [3, 'This is third']
        ];
        $this->writer->write($data);
    }

}

abstract class Writer {
    protected $_file;
    public function __construct($file){
        $this->_file = $file;
    }

    // 强制要求子类定义 write 方法
    abstract function write($data);
}

/**
 * 以 CSV 格式写入数据
 */
class CsvWriter extends Writer {

    public function write($data) {
        $fp = fopen($this->_file, 'w');
        foreach ($data as $row) {
            fputcsv($fp, $row);
        }
        fclose($fp);
    }

}

/**
 * 以 Html 格式写入数据
 */
class HtmlWriter extends Writer {

    public function write($data) {
        $fp = fopen($this->_file, 'w');
        fwrite($fp, '<table>');
        foreach ($data as $row) {
            fwrite($fp, '<tr><td>' . implode('</td><td>', $row) . '</td></tr>\n');
        }
        fwrite($fp, '</table>');
        fclose($fp);
    }

}

$model = new DataModel(new CsvWriter('D:\test.csv'));
$model->export();
```

```php
/**
 *
 * 需求：假如一个电商网站系统，针对男性女性用户要各自跳转到不同的商品类目，并且所有广告位展示不同的广告
 */
 
 // UserStrategy.php 
 <?php

namespace IMooc;

interface UserStrategy {

    function showAd();

    function showCategory();

}

// FemaleUserStrategy.php
<?php

namespace IMooc;

class FemaleUserStrategy implements UserStrategy {

    function showAd () {
        echo '2014新款女装';
    }

    function showCategory () {
        echo '女装';
    }

}

// MaleUserStrategy.php
<?php
namespace IMooc;

class MaleUserStrategy implements UserStrategy {

    function showAd () {
        echo '2014男装秀';
    }

    function showCategory () {
        echo '男装';
    }

}

// 调用
class Page {

    protected $strategy;

    function index () {
        echo 'AD:';
        $this->strategy->showAd();

        echo '<br/>';

        echo 'Category:';
        $this->strategy->showCategory();
    }

    function setStrategy (\IMooc\UserStrategy $strategy) {
        $this->strategy = $strategy;
    }


}

$page = new Page;

if (isset($_GET['female'])) {
    $strategy = new \IMooc\FemaleUserStrategy();
} else {
    $strategy = new \IMooc\MaleUserStrategy();
}

$page->setStrategy($strategy);
$page->index();
 
```

### 观察者模式
- 什么是观察者模式？
  定义了一种一对多的依赖关系，让多个观察者对象同时监听某一个主题对象。这个主题对象在状态发生变化时，会通知所有观察者对象，使它们能够自动更新自己。

应用场景：一个事件发生后，要执行一连串更新操作。传统的编程方式，就是在事件的代码之后直接加入处理逻辑。当更新的逻辑增多之后，代码会变得难以维护。这种方式是耦合的，侵入式的，增加新的逻辑需要修改事件主体的代码。
```php
// EventGenerator.php 中
<?php

namespace IMooc;

abstract class EventGenerator {

    private $observers = array();

    // 添加观察者
    function addObserver (Observer $observer) {
        $this->observers[] = $observer;
    }

    // 调用每一个观察者的 update 方法
    function notify () {
        foreach ($this->observers as $observer) {
            $observer->update();
        }
    }

}
```

Observer.php 中
```php
<?php

namespace IMooc;

interface Observer {

    function update ($envent_info = null);


}
```

调用观察者
```php
class Event extends \IMooc\EventGenerator {

    function trigger () {
        $this->notify();
    }


}

// 观察者1
class Observer1 implements \IMooc\Observer {

    function update ($event_info = null) {
        echo "逻辑1<br /> \n";
    }

}

// 观察者2
class Observer2 implements \IMooc\Observer {

    function update ($event_info = null) {
        echo "逻辑2<br /> \n";
    }

}

$event = new Event;
$event->addObserver(new Observer1);
$event->addObserver(new Observer2);
$event->trigger();
```

### 工厂模式
- 什么是工厂模式？
  工厂方法或者类生成对象，而不是在代码中直接 new，好处在于改动一个类的名称或者参数时，只需要改动一个地方
```php
// Factory.php
<?php
namespace IMooc;

class Factory {

    public static function createDatabase () {
        // 在这里 实例化需要用到的对象
        return new Database();
    }

}
```

外部调用时
```php
<?php
var_dump(IMooc\Factory::createDatabase());
```

### 注册树模式
- 什么是注册树模式？
  解决全局共享和交换对象

```php
// Register.php
<?php
namespace IMooc;

class Register {

    protected static $objects;

    /**
     * 注册实例
     * @param [type] $alias  实例别名
     * @param [type] $object 实例对象
     */
    static function set ($alias, $object) {
        self::$objects[$alias] = $object;
    }

    // 获取对象实例
    static function get ($name) {
        return self::$objects[$name];
    }

    // 删除对象实例
    static function _unset ($alias) {
        unset(self::$objects[$alias]);
    }

}
```
可以放在工厂文件中之后再在其他地方调用
```php
// Factory.php 中设置
Register::set('db', Database::getInstance());

// 其它文件中调用
$db = \IMooc\Register::get('db');
```

### 适配器模式
- 什么是适配器模式？
1. 适配器模式，可以将截然不同的函数接口封装成统一的API；
2. 实际应用举例，PHP的数据库操作有mysql，mysqli，pdo 3种，可以用适配器模式统一成一致。类似的场景还有 cache 适配器，将 memcache，redis，file，apc等不同的缓存函数，统一成一致。
```php
// Database.php 文件中
<?php
namespace IMooc;

interface IDatabase {

    function connect ($host, $user, $passwd, $dbname);

    function query($sql);

    function close();

}

class Database {

    protected static $db;

    // 获取对象实例
    public static function getInstance(){

        if (!(self::$db instanceof self)){
            self::$db = new self();
        }
        return self::$db;
    }

    private function __construct(){}

    private function __clone(){}

    public static function where($where){
        return $this;
    }

    public static function order($order){
        return $this;
    }


    public static function limit($limit){
        return $this;
    }


}
```
IMooc\Database\PDO.php 文件中
```php
<?php
namespace IMooc\Database;

use IMooc\IDatabase;

class PDO implements IDatabase {

    protected $conn;

    function connect ($host, $user, $passwd, $dbname) {
        $this->conn = new \PDO("mysql:host=$host;dbname=$dbname", $user, $passwd);
    }

    function query ($sql) {
        return $this->conn->query($sql);
    }

    function close () {
        unset($this->conn);
    }

}
```
调用
```php
$db = new IMooc\Database\PDO();

$db->connect('127.0.0.1', 'root', '123456', 'test');
$res = $db->query('show databases');
$db->close();
```

### 数据对象映射模式
- 什么是数据对象映射模式？
1. 数据对象映射模式，是将对象和数据存储映射起来，对一个对象的操作会映射为对数据存储的操作。

```php
<?php
namespace IMooc;

class User {

    public $id;
    public $name;
    public $phone;
    public $regtime;

    protected $db;

    function __construct ($id) {
        $this->db = new \IMooc\Database\MySQLi();
        $this->db->connect('127.0.0.1', 'root', '123456', 'test');
        $res = $this->db->query('select * from user limit 1');
        $data = $res->fetch_assoc();

        $this->id = $id;
        $this->name = $data['name'];
        $this->phone = $data['phone'];
        $this->regtime = $data['regtime'];

    }

    // 运用析构函数的特性实现更新操作
    function __destruct () {
        $this->db->query("update user set name = '{$this->name}',
            phone = '{$this->phone}',regtime = '$this->regtime'
              where id = '{$this->id}' limit 1 ");
    }

}
```
调用
```php
$user = new \IMooc\User(1);


$user->phone = 18502728040;
$user->name = '张三';
$user->regtime = time();
```

### 原型模式
- 什么是原型模式？
1. 与工厂模式作用类似，都是用来创建对象。
2. 与工厂模式的实现不同，原型模式是先创建好一个原型对象，然后通过 clone 原型对象来创建新的对象。这样就免去了类创建时重复的初始化操作。
3. 原型模式适用于大对象的创建。创建一个大对象需要很大的开销，如果每次 new 就会消耗很大，原型模式仅需内存拷贝即可。

传统写法
```php
// 实例化画布对象
$canvas1 = new IMooc\Canvas();
// 初始化画布操作
$canvas1->init();
// 对画布1进行绘制
$canvas1->rect(3,6,4,12);
$canvas1->draw();


// 实例化画布对象
$canvas2 = new IMooc\Canvas();
// 初始化画布操作
$canvas2->init();
// 对画布2进行绘制
$canvas1->rect(1,6,4,18);
$canvas2->draw();
```

使用原型模式
```php
// 实例化画布对象
$prototype = new IMooc\Canvas();
// 初始化画布操作
$prototype->init();


$canvas1 = clone $prototype;
// 对画布1进行绘制
$canvas1->rect(3,6,4,12);
$canvas1->draw();

$canvas2 = clone $prototype;
// 对画布2进行绘制
$canvas2->rect(3,6,4,12);
$canvas2->draw();
```

### 装饰器模式
1. 装饰器模式（Decorator），可以动态地添加修改类的功能。
2. 一个类提供了一项功能，如果要在修改并添加额外的功能，传统的编程模式，需要写一个子类继承它，并重新实现类的方法。
3. 使用装饰器模式，仅需要在运行时添加一个装饰器对象即可实现，可以实现最大的灵活性。

IMooc\Canvas.php

```php

<?php
namespace IMooc;

class Canvas
{
    public $data;

    /**
     * 保存装饰器
     *
     * @var array
     */
    protected $decorators = array();

    /**
     * 初始化画布
     *
     * @param int $width
     * @param int $height
     */
    function init($width = 20, $height = 10)
    {
        $data = array();
        for($i = 0; $i < $height; $i++)
        {
            for($j = 0; $j < $width; $j++)
            {
                $data[$i][$j] = '*';
            }
        }
        $this->data = $data;
    }

    /**
     * 添加装饰器
     *
     * @param DrawDecorator $decorator
     */
    function addDecorator(DrawDecorator $decorator)
    {
        $this->decorators[] = $decorator;
    }

    /**
     * 画画前的调用方法
     */
    function beforeDraw()
    {
        foreach($this->decorators as $decorator)
        {
            $decorator->beforeDraw();
        }
    }

    /**
     * 画画后的调用方法
     */
    function afterDraw()
    {
        // 需要进行反转，beforeDraw 方法是先进，afterDraw 方法是先出  （先进先出，后进后出）
        $decorators = array_reverse($this->decorators);
        foreach($decorators as $decorator)
        {
            $decorator->afterDraw();
        }
    }

    /**
     * 画画
     */
    function draw()
    {
        $this->beforeDraw();
        foreach($this->data as $line)
        {
            foreach($line as $char)
            {
                echo $char;
            }
            echo "<br />\n";
        }
        $this->afterDraw();
    }

    function rect($a1, $a2, $b1, $b2)
    {
        foreach($this->data as $k1 => $line)
        {
            if ($k1 < $a1 or $k1 > $a2) continue;
            foreach($line as $k2 => $char)
            {
                if ($k2 < $b1 or $k2 > $b2) continue;
                $this->data[$k1][$k2] = '&nbsp;';
            }
        }
    }
}



```


IMooc\DrawDecorator.php

```php

<?php
namespace IMooc;

interface DrawDecorator
{
    function beforeDraw();
    function afterDraw();
}

```

IMooc\ColorDrawDecorator.php

```php

<?php
namespace IMooc;

class ColorDrawDecorator implements DrawDecorator
{
    protected $color;
    function __construct($color = 'red')
    {
        $this->color = $color;
    }
    function beforeDraw()
    {
        echo "<div style='color: {$this->color};'>";
    }
    function afterDraw()
    {
        echo "</div>";
    }
}

```


IMooc\SizeDrawDecorator.php

```php

<?php
namespace IMooc;

class SizeDrawDecorator implements DrawDecorator
{
    protected $size;
    function __construct($size = '14px')
    {
        $this->size = $size;
    }

    function beforeDraw()
    {
        echo "<div style='font-size: {$this->size};'>";
    }

    function afterDraw()
    {
        echo "</div>";
    }
}

```


调用

```php

$canvas = new \IMooc\Canvas();
$canvas->init();
$canvas->addDecorator(new \IMooc\ColorDrawDecorator('green'));
// $canvas->addDecorator(new \IMooc\SizeDrawDecorator('10px'));
$canvas->rect(3, 6, 4, 12);
$canvas->draw();

```


### 迭代器模式

1. 迭代器模式，在不需要了解内部实现的前提下，遍历一个聚合对象的内部元素。
2. 相比于传统的编程模式，迭代器模式可以隐藏遍历元素的所需的操作。


IMooc\AllUser.php

```php

<?php
namespace IMooc;

class AllUser implements \Iterator
{
    protected $ids;
    protected $data = array();
    protected $index;  // 迭代器的当前位置

    function __construct()
    {
        $db = Factory::getDatabase();
        $result = $db->query("select id from user");
        $this->ids = $result->fetch_all(MYSQLI_ASSOC);
    }

    // 当前元素
    function current()
    {
        $id = $this->ids[$this->index]['id'];
        // 获取指定 id 的用户信息
        return Factory::getUser($id);
    }

    // 下一个元素
    function next()
    {
        $this->index ++;
    }

    // 是否还有下一个元素
    function valid()
    {
        return $this->index < count($this->ids);
    }

    // 重置迭代器
    function rewind()
    {
        $this->index = 0;
    }

    // 获取当前的位置
    function key()
    {
        return $this->index;
    }

}

```

调用

```php

// 所有的用户信息
$users = new \IMooc\AllUser();

foreach ($users as $user) {
    var_dump($user);
}

```


### 代理模式
1. 在客户端与实体之间建立一个代理对象 （proxy），客户端对实体进行操作全部委派给代理对象，隐藏实体的具体实现细节。
2. proxy 还可以与业务代码分离，部署到另外的服务器。业务代码中通过 RPC 来委派任务。

IMooc\Proxy.php

```php

<?php
namespace IMooc;

class Proxy implements IUserProxy
{
    // 从数据库用于 读取 操作
    function getUserName($id)
    {
        $db = Factory::getDatabase('slave');
        $db->query("select name from user where id =$id limit 1");
    }

    // 主数据库用于 更新 操作
    function setUserName($id, $name)
    {
        $db = Factory::getDatabase('master');
        $db->query("update user set name = $name where id =$id limit 1");
    }
}

```


IMooc\IUserProxy.php

```php

<?php
namespace IMooc;

interface IUserProxy
{
    function getUserName($id);
    function setUserName($id, $name);
}

```


调用

```php

$proxy = new \IMooc\Proxy();
$proxy->getUserName($id);
$proxy->setUserName($id, $name);

```