# Dart

## 关键字速览

| [abstract](https://www.dartcn.com/guides/language/language-tour#抽象类) 2 | [dynamic](https://www.dartcn.com/guides/language/language-tour#重要的概念) 2 | [implements](https://www.dartcn.com/guides/language/language-tour#隐式接口) 2 | [show](https://www.dartcn.com/guides/language/language-tour#导入库的一部分) 1 |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [as](https://www.dartcn.com/guides/language/language-tour#类型判定运算符) 2 | [else](https://www.dartcn.com/guides/language/language-tour#if-和-else) | [import](https://www.dartcn.com/guides/language/language-tour#使用库) 2 | [static](https://www.dartcn.com/guides/language/language-tour#类变量和方法) 2 |
| [assert](https://www.dartcn.com/guides/language/language-tour#assert) | [enum](https://www.dartcn.com/guides/language/language-tour#枚举类型) | [in](https://www.dartcn.com/guides/language/language-tour#for-循环) | [super](https://www.dartcn.com/guides/language/language-tour#扩展类继承) |
| [async](https://www.dartcn.com/guides/language/language-tour#异步支持) 1 | [export](https://www.dartcn.com/guides/libraries/create-library-packages) 2 | [interface](https://stackoverflow.com/questions/28595501/was-the-interface-keyword-removed-from-dart) 2 | [switch](https://www.dartcn.com/guides/language/language-tour#switch-和-case) |
| [await](https://www.dartcn.com/guides/language/language-tour#异步支持) 3 | [extends](https://www.dartcn.com/guides/language/language-tour#扩展类继承) | [is](https://www.dartcn.com/guides/language/language-tour#类型判定运算符) | [sync](https://www.dartcn.com/guides/language/language-tour#生成器) 1 |
| [break](https://www.dartcn.com/guides/language/language-tour#break-和-continue) | [external](https://stackoverflow.com/questions/24929659/what-does-external-mean-in-dart) 2 | [library](https://www.dartcn.com/guides/language/language-tour#库和可见性) 2 | [this](https://www.dartcn.com/guides/language/language-tour#构造函数) |
| [case](https://www.dartcn.com/guides/language/language-tour#switch-和-case) | [factory](https://www.dartcn.com/guides/language/language-tour#工厂构造函数) 2 | [mixin](https://www.dartcn.com/guides/language/language-tour#为类添加功能mixins) 2 | [throw](https://www.dartcn.com/guides/language/language-tour#throw) |
| [catch](https://www.dartcn.com/guides/language/language-tour#catch) | [false](https://www.dartcn.com/guides/language/language-tour#booleans) | [new](https://www.dartcn.com/guides/language/language-tour#使用构造函数) | [true](https://www.dartcn.com/guides/language/language-tour#booleans) |
| [class](https://www.dartcn.com/guides/language/language-tour#实例变量) | [final](https://www.dartcn.com/guides/language/language-tour#final-和-const) | [null](https://www.dartcn.com/guides/language/language-tour#默认值) | [try](https://www.dartcn.com/guides/language/language-tour#catch) |
| [const](https://www.dartcn.com/guides/language/language-tour#final-和-const) | [finally](https://www.dartcn.com/guides/language/language-tour#finally) | [on](https://www.dartcn.com/guides/language/language-tour#catch) 1 | [typedef](https://www.dartcn.com/guides/language/language-tour#typedefs) 2 |
| [continue](https://www.dartcn.com/guides/language/language-tour#break-和-continue) | [for](https://www.dartcn.com/guides/language/language-tour#for-循环) | [operator](https://www.dartcn.com/guides/language/language-tour#重写运算符) 2 | [var](https://www.dartcn.com/guides/language/language-tour#变量) |
| [covariant](https://www.dartcn.com/guides/language/sound-problems#the-covariant-keyword) 2 | [Function](https://www.dartcn.com/guides/language/language-tour#函数) 2 | [part](https://www.dartcn.com/guides/libraries/create-library-packages#organizing-a-library-package) 2 | [void](https://medium.com/dartlang/dart-2-legacy-of-the-void-e7afb5f44df0) |
| [default](https://www.dartcn.com/guides/language/language-tour#switch-和-case) | [get](https://www.dartcn.com/guides/language/language-tour#getters-和-setters) 2 | [rethrow](https://www.dartcn.com/guides/language/language-tour#catch) | [while](https://www.dartcn.com/guides/language/language-tour#while-和-do-while) |
| [deferred](https://www.dartcn.com/guides/language/language-tour#延迟加载库) 2 | [hide](https://www.dartcn.com/guides/language/language-tour#导入库的一部分) 1 | [return](https://www.dartcn.com/guides/language/language-tour#函数) | [with](https://www.dartcn.com/guides/language/language-tour#为类添加功能mixins) |
| [do](https://www.dartcn.com/guides/language/language-tour#while-和-do-while) | [if](https://www.dartcn.com/guides/language/language-tour#if-和-else) | [set](https://api.dartlang.org/stable/dart-core/Set-class.html) 2 | [yield](https://www.dartcn.com/guides/language/language-tour#生成器) 3 |

避免使用这些单词作为标识符。 但是，如有必要，标有上标的关键字可以用作标识符：

- 带有 **1** 上标的单词为 **上下文关键字**， 仅在特定位置具有含义。 他们在任何地方都是有效的标识符。
- 带有 **2** 上标的单词为 **内置标识符**， 为了简化将 JavaScript 代码移植到 Dart 的工作， 这些关键字在大多数地方都是有效的标识符， 但它们不能用作类或类型名称，也不能用作 import 前缀。
- 带有 **3** 上标的单词是与 Dart 1.0 发布后添加的[异步支持](https://www.dartcn.com/guides/language/language-tour#异步支持)相关的更新，作为限制类保留字。
  不能在标记为 `async` ，`async*` 或 `sync*` 的任何函数体中使用 `await` 或 `yield` 作为标识符。

关键字表中的剩余单词都是**保留字**。 不能将保留字用作标识符。

