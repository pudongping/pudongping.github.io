---
title: 一文吃透 Promise 与 async/await，异步编程也能如此简单！建议收藏！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 前端
tags:
  - 前端
  - promise
  - 异步编程
abbrlink: da88494b
date: 2025-08-12 11:39:29
img:
coverImg:
password:
summary:
---

在现代编程开发中，“异步”两个字几乎贯穿始终：你写的接口请求、定时器、事件监听、动画控制……背后都绕不开异步编程。

那到底啥是**异步**呢？说到异步，我们就需要结合**同步**来讲讲，这样就更加清晰明了。

今天这篇文章，我们来讲讲在前端编程中的异步。

![](https://upload-images.jianshu.io/upload_images/14623749-10a759c049bf064e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

简单点说：

- **同步操作**：代码一行一行执行，上一行没跑完，下一行就得等着。
- **异步操作**：某些任务（比如网络请求）可能需要等一段时间才能完成，而程序可以先不等它，继续往下执行，等它好了再回来处理。

举个栗子🌰

```js
console.log("1. 开始做饭");

setTimeout(() => {
  console.log("2. 饭做好了");
}, 3000);

console.log("3. 玩手机中...");
```

你将看到的输出内容是：

```markdown
1. 开始做饭
3. 玩手机中...
2. 饭做好了（大约 3 秒后）
```

这就体现了异步：**做饭**这件事需要 3 秒时间，不影响你“玩手机”的下一步操作。

---

## 🌱 一、为什么我们需要 Promise？

早期处理异步的方式是**回调函数（callback）**。

```js
function cook(callback) {
  setTimeout(() => {
    callback("🍳 饭做好了");
  }, 3000);
}

cook((msg) => {
  console.log(msg);
});
```

上面的代码，貌似也没啥太大的问题，但是，如果是需要**做三道菜**呢？代码可能就是这样的了：

```js
cook((msg1) => {
  console.log(msg1);
  cook((msg2) => {
    console.log(msg2);
    cook((msg3) => {
      console.log(msg3);
      // 继续嵌套……
    });
  });
});
```

是的，这就是恐怖的“**回调地狱**”：

明眼人都能够看得出来，这样的代码太难维护，太容易出 bug，于是 ES6 就推出了 Promise 来解决这个问题。

---

## 🔐 二、什么是 Promise？

一句话总结：

> Promise 是 JavaScript 中处理异步操作的一种对象，它表示一个还没完成但将来一定会完成（或失败）的操作结果。

它有三个状态：

| 状态        | 意义       |
| --------- | -------- |
| pending   | 初始状态，进行中 |
| fulfilled | 操作成功     |
| rejected  | 操作失败     |

状态一旦变化，就不可逆。

---

## ✨ 三、Promise 基本用法

我们来写一个最简单的 Promise：

```js
const p = new Promise((resolve, reject) => {
  const success = true;

  if (success) {
    resolve("成功啦！");
  } else {
    reject("失败了～");
  }
});

p.then((result) => {
  console.log("成功：", result);
}).catch((err) => {
  console.error("失败：", err);
});
```

### 🔍 解读一下：

* `resolve(value)` 表示异步任务**成功**，传出结果。
* `reject(error)` 表示异步任务**失败**。
* `.then()` 是成功的回调，`.catch()` 是失败的回调。

---

## 🔗 四、Promise 链式调用

你可以连续使用 `.then()` 来处理一系列异步任务：

```js
new Promise((resolve) => {
  setTimeout(() => {
    console.log("1️⃣ 数据加载中...");
    resolve("原始数据");
  }, 1000);
})
  .then((data) => {
    console.log("2️⃣ 第一次处理：", data);
    return data + " + 步骤一完成";
  })
  .then((newData) => {
    console.log("3️⃣ 第二次处理：", newData);
  })
  .catch((err) => {
    console.error("发生错误：", err);
  });
```

每个 `.then()` 的返回值会传给下一个 `.then()`。

---

## ⚙️ 五、实战：接口请求场景

实际开发中我们经常写这种：

```javascript
fetch("https://api.example.com/user/123")
  .then((res) => {
    if (!res.ok) throw new Error("网络请求失败！");
    return res.json(); // 转成 JSON
  })
  .then((user) => {
    console.log("用户信息：", user);
  })
  .catch((err) => {
    console.error("请求失败：", err);
  });
```

这里用的 `fetch` 就是一个返回 Promise 的 API。

---

## 🎭 六、async/await 是什么？

async/await 是 ES7 对 Promise 的一种 **语法糖**，让我们用看起来是**同步的写法**来写异步代码，这使得代码可以更清晰、可读性更高。

来，让我们来感受一下差距👇

### 👀 Promise 写法

```js
fetch("/api/data")
  .then((res) => res.json())
  .then((data) => console.log("数据：", data))
  .catch((err) => console.error("出错了", err));
```

### async/await 写法

```js
async function getData() {
  try {
    const res = await fetch("/api/data");
    const data = await res.json();
    console.log("数据：", data);
  } catch (err) {
    console.error("出错了", err);
  }
}

getData();
```

将代码做了一个简单的调整之后，是不是感觉瞬间就清爽多了？

- **async** 表示这个函数里有异步逻辑
- **await** 用来“等待”一个 Promise 的结果

---

## 📌 七、async/await 的几个重点规则

### ✅ 1. async 函数总是返回一个 Promise

```js
async function hello() {
  return "你好";
}

hello().then((res) => console.log(res)); // 输出 "你好"
```

### ✅ 2. await 只能用在 async 函数内部

```js
await fetch("/api"); // ❌ 语法错误

// 正确写法👇
async function load() {
  const res = await fetch("/api");
}
```

### ✅ 3. 错误要 try/catch 包裹住

```js
async function load() {
  try {
    const res = await fetch("/api/bad-url");
    const data = await res.json();
  } catch (err) {
    console.error("捕获错误：", err);
  }
}
```

---

## 🧪 八、多个异步任务的组合

### Promise.all

如果同时发出多个请求，并且需要等全部成功后再处理，那么我们就需要使用：

```js
const getUser = fetch("/api/user");
const getPosts = fetch("/api/posts");

const [userRes, postsRes] = await Promise.all([getUser, getPosts]);
```

### Promise.race

如果有多个异步任务同时执行，但是只需要谁先执行完毕就用谁，那么就可以使用：

```js
const res = await Promise.race([fetch("/a"), fetch("/b")]);
```

以上的代码，如果 `/a` 先执行完，那么就 `res` 就是 `/a` 的返回结果，如果是 `/b` 先返回完，那么就是 `/b` 的返回结果。总之，谁执行的快，谁先执行完成，`res` 就是谁的结果。


---

## 🏁 总结一下

Promise 是前端异步编程的核心工具，而 async/await 则可以让你的代码更加优雅，如果你能够熟练掌握它们：

* 能让你的代码更清晰；
* 能写出更优雅、可维护的业务逻辑；
* 还能在面试中脱颖而出！


