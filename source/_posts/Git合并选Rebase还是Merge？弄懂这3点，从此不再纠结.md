---
title: Git合并选Rebase还是Merge？弄懂这3点，从此不再纠结
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Git
tags:
  - Git
abbrlink: a479708b
date: 2025-09-01 10:32:55
img:
coverImg:
password:
summary:
---

大家好！今天我们来聊聊 Git 中两个非常重要但又容易混淆的概念：`git rebase` 和 `git merge`。

在日常团队协作开发中，我们经常需要将不同分支的代码进行合并。Git 提供了两种主要的合并方式：`git merge` 和 `git rebase`。虽然它们都能实现分支合并的目的，但使用方式和最终效果却大不相同。

那么，到底应该用 `rebase` 还是 `merge` 呢？它们有什么区别？什么时候用哪个更合适？

## 什么是 Git Merge？

### 概念解释

`git merge` 是 Git 中最常用的分支合并命令。当你使用 merge 命令时，Git 会创建一个新的提交（merge commit），保留了两个分支的所有历史记录。

### 基本用法

```bash
# 切换到目标分支（通常是 main 或 master）
git checkout main

# 合并 feature 分支
git merge feature-branch
```

### 实际场景演示

假设我们有这样一个场景：你在 `feature-login` 分支上开发登录功能，同时主分支 `main` 也有其他同事提交了新代码。

```bash
# 初始状态
# main:     A---B---C
# feature:      \---D---E

# 执行 merge 后
# main:     A---B---C---F (merge commit)
#               \       /
# feature:       \-D---E
```

代码示例：

```bash
# 1. 创建并切换到功能分支
git checkout -b feature-login

# 2. 在功能分支上进行开发
echo "登录功能实现" > login.js
git add login.js
git commit -m "添加登录功能"

# 3. 切换回主分支
git checkout main

# 4. 模拟其他人的提交
echo "修复了一个 bug" > bugfix.js
git add bugfix.js
git commit -m "修复用户头像显示问题"

# 5. 合并功能分支
git merge feature-login
```

## 什么是 Git Rebase？

### 概念解释

`git rebase` 的英文直译是「变基」，它可以**将一个分支的提交「移植」到另一个分支上**，使得提交历史呈现为一条直线，更加清晰整洁。

Rebase 会将当前分支的提交「复制」到目标分支的最新提交之后，然后放弃原来的提交。这样看起来就像是直接从目标分支的最新提交开始开发的，相当于将需要合并的分支上的提交“重放”到了合并的目标分支上。

### 基本用法

```bash
# 方式一：在功能分支上执行
git checkout feature-branch
git rebase main

# 方式二：直接指定分支
git rebase main feature-branch
```

### 实际场景演示

还是用刚才登录功能的例子：

```bash
# rebase 前
# main:     A---B---C
# feature:      \---D---E

# 执行 rebase 后
# main:     A---B---C
# feature:            \---D'---E'
```

代码示例：

```bash
# 1. 在功能分支上进行 rebase
git checkout feature-login
git rebase main

# 2. 如果有冲突，解决后继续
git add .
git rebase --continue

# 3. 将变基后的分支合并到主分支（这时会是快进合并）
git checkout main
git merge feature-login  # 这将是一个 fast-forward merge
```

## 两者的核心区别

### 1. 提交历史的差异

**Merge 的特点：**
- 保留完整的提交历史
- 会产生合并提交
- 分支结构清晰可见
- 历史记录可能比较复杂

**Rebase 的特点：**
- 创造线性的提交历史
- 不会产生额外的合并提交
- 看起来更加整洁
- 会改变提交的 SHA 值

### 2. 可视化对比

让我们用一个更直观的例子来看看：

```bash
# 使用 Merge 后的历史
*   a1b2c3d (HEAD -> main) Merge branch 'feature-login'
|\  
| * d4e5f6g 添加密码加密功能
| * h7i8j9k 实现用户登录逻辑
* | k1l2m3n 优化首页加载速度
* | n4o5p6q 修复导航栏样式问题
|/  
* q7r8s9t 初始提交

# 使用 Rebase 后的历史
* f9g8h7i (HEAD -> main) 添加密码加密功能
* e6f7g8h 实现用户登录逻辑
* d3e4f5g 优化首页加载速度
* c2d3e4f 修复导航栏样式问题
* a1b2c3d 初始提交
```

## 什么时候使用 Merge？

### 适用场景

1. **功能分支合并**：当你完成一个完整的功能开发时
2. **团队协作**：多人协作时保留清晰的分支结构
3. **重要的里程碑**：需要明确标记合并点的时候
4. **开源项目**：需要保留贡献者的提交历史

### 实际示例

```bash
# 场景：完成用户管理功能的开发
git checkout main
git pull origin main  # 确保主分支是最新的
git merge feature-user-management
git push origin main

# 查看合并历史
git log --graph --oneline
```

## 什么时候使用 Rebase？

### 适用场景

1. **清理提交历史**：在推送到远程仓库前整理提交
2. **同步主分支更新**：将主分支的新变更同步到功能分支
3. **个人开发**：在个人分支上工作时
4. **线性历史偏好**：团队偏好简洁的线性历史

### 实际示例

```bash
# 场景一：同步主分支最新变更
git checkout feature-dashboard
git rebase main  # 将主分支的新提交应用到当前分支

# 场景二：交互式 rebase 清理提交历史
git rebase -i HEAD~3  # 整理最近 3 次提交
```

### 交互式 Rebase 示例

```bash
# 执行交互式 rebase
git rebase -i HEAD~3

# 编辑器会显示类似内容：
pick a1b2c3d 添加用户登录接口
pick d4e5f6g 修复登录bug
pick g7h8i9j 添加登录日志

# 你可以进行以下操作：
# pick: 保留这个提交
# reword: 保留提交但修改提交信息
# edit: 保留提交但暂停以便修改
# squash: 将这个提交合并到前一个提交
# drop: 删除这个提交
```

## 冲突处理

### Merge 冲突处理

```bash
# 当出现合并冲突时
git merge feature-branch

# Git 会提示冲突，编辑冲突文件
# 解决冲突后
git add .
git commit  # Git 会自动生成合并提交信息
```

### Rebase 冲突处理

```bash
# 当出现 rebase 冲突时
git rebase main

# 解决冲突后
git add .
git rebase --continue

# 如果想放弃 rebase
git rebase --abort
```

### 解决冲突的差异

- Merge 在一次合并提交中解决所有冲突
- Rebase 可能会在**每个被复制的提交处**都需要解决冲突

## 最佳实践建议

### 1. 团队约定

推荐的工作流：
1. 功能分支开发期间：使用 rebase 同步主分支更新
2. 功能开发完成后：使用 merge 合并到主分支
3. 推送前：使用交互式 rebase 清理提交历史

### 2. 安全原则

Rebase 的黄金法则:

**永远不要对已经推送到远程仓库的提交执行 rebase!**
这是因为 rebase 会重写提交历史，如果其他人已经基于这些提交进行开发，会导致严重的协作问题。

```bash
# 永远不要对已经推送到远程的提交执行 rebase！
# ❌ 错误做法（如果已经推送到远程）
git rebase main

# ✅ 正确做法
git pull --rebase origin main 
# 或者
git merge main
```

### 3. 实用命令组合

```bash
# 常用的 rebase 命令
git pull --rebase    # 拉取时使用 rebase 而不是 merge
git config pull.rebase true  # 设置 pull 默认使用 rebase

# 查看分支图
git log --graph --pretty=oneline --abbrev-commit

# 撤销上次 merge（如果还没推送）
git reset --hard HEAD~1
```

## 总结

选择 `rebase` 还是 `merge` 并没有绝对的对错，关键是要根据具体场景和团队约定来决定：

- **需要保留完整历史和分支结构**：选择 `merge`
- **希望保持线性、整洁的提交历史**：选择 `rebase`
- **多人协作的公共分支**：谨慎使用 `rebase`
- **个人开发的功能分支**：`rebase` 是很好的选择

记住，工具本身没有好坏，关键是要理解它们的特点，在合适的场景下使用合适的工具。在团队协作中，**保持一致的使用规范**比选择哪种方式更重要。与团队成员协商确定适合项目的合并策略，才能让版本历史既美观又实用。

希望本文能帮助你更好地理解和使用 Git 的合并工具。

有什么问题欢迎在评论区讨论，我们一起进步！ 🚀
