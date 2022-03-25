# MVC及其衍生

原文链接：[MVC，MVP 和 MVVM 的图示 - 阮一峰](https://www.ruanyifeng.com/blog/2015/02/mvcmvp_mvvm.html)

## MVC

![img](https://www.ruanyifeng.com/blogimg/asset/2015/bg2015020105.png)

所有通信都是单向的。

1.  View 传送指令到 Controller
2.  Controller 完成业务逻辑后，要求 Model 改变状态
3.  Model 将新的数据发送到 View，用户得到反馈

## MVP

MVP 模式将 Controller 改名为 Presenter，同时改变了通信方向。

![img](https://www.ruanyifeng.com/blogimg/asset/2015/bg2015020109.png)

1. 各部分之间的通信，都是双向的。

2. View 与 Model 不发生联系，都通过 Presenter 传递。

3. View 非常薄，不部署任何业务逻辑，称为"被动视图"（Passive View），即没有任何主动性，而 Presenter非常厚，所有逻辑都部署在那里。

## MVVM

![img](https://www.ruanyifeng.com/blogimg/asset/2015/bg2015020110.png)

唯一的区别是，它采用双向绑定（data-binding）：View的变动，自动反映在 ViewModel，反之亦然。