# JavaScript笔记

## 概述

- ECMAScript是JavaScript的语言标准，JavaScript是ECMAScript的一种实现。
- JavaScript算是一种嵌入式语言，本身没有IO函数，需要宿主来提供。最典型的宿主是浏览器和服务器端的Node项目。
- JavaScript不是纯粹的面向对象语言，也支持其他编程范式，例如函数式编程。
- JavaScript的核心语法部分十分精简，只包含基本的语法构造和标准库，更多地依赖于宿主环境提供的API。

### 重要的编程风格

1. Javascript中区块起首的大括号紧跟在关键字后面，避免[难以察觉的错误](https://wangdoc.com/javascript/features/style.html#区块)。

2. 尽量使用严格相等运算符`===`，以避免[自动类型转换带来的不确定性](https://wangdoc.com/javascript/features/style.html#相等和严格相等)。

3. `switch... case`结构可以用面向对象结构替代。

   ```
   function doActions(action) {
       var actions = {
           'one': function () {
               return 1;
           },
           'two': function () {
               return 2;
           },
           'run': function () {
               return 'run';
           }
       };
   
       if (typeof actions[action] !== 'function') {
           throw new Error('Invalid action.');
       }
   
       return actions[action]();
   }
   ```

4. 严格模式，在产生任何实际运行效果的语句前添加`'use strict'`字符串。

## 开发环境

### 控制台与console对象

向控制台输出信息，自动在连续两个参数产生的输出间添加空格，并在每次输出的结尾添加换行符。

1. `console.log()`, `console.debug()`, `console.info()`
2. `console.warn()`
3. `console.error()`

支持格式字符串，如`%s`, `%d`或`%i`, `%f`, 对象的链接`%o`。

可以按自己的需要覆盖console的方法，如为console的输出添加时间字符串。

```
['log', 'warn', 'error'].forEach(function (method) {
    console[method] = console[method].bind(
        console,
        new Date().toISOString()
    );
});
```

`console.table`可以将符合类型的数据转换为表格显示。

```
console.table(
    {
        Alice: { name: 'Alice Kane', score: 32 },
        Bob: {name: 'Bob Kingdom', score: 44 }
    } // 或数组类型
);
```

- `console.dir()`用于对对象进行审察，格式比直接使用`console.log()`美观。
- `console.count('tag')`用于计数，输出它被调用了多少次。
- `console.time('tag')`和`console.timeEnd('tag')`用于计算操作花费的时间。
- `console.group()`, `console.groupEnd()`和`console.groupCollapsed()`用于对大量信息进行分组。
- `console.trace()`用于显示调用栈。
- `console.clear()`用于清空输出。

## let 和 const

`let`和`const`命令用于声明变量。

`let`声明的变量是可变的，`const`声明的变量是不可变的。

```js
let foo = 1;
foo = 2;

const bar = 1;
bar = 2; // 报错
```

注意，如果`const`声明的变量指向一个对象，那么该对象的属性是可变的。

```js
const foo = {
  bar: 1
};

foo.bar = 2;
```

上面代码中，变量`foo`本身是不可变的，即`foo`不能指向另一个对象。但是，对象内部的属性是可变的，这是因为这时`foo`保存的是一个指针，这个指针本身不可变，但它指向的对象本身是可变的。

## IIFE

[参考文档](https://developer.mozilla.org/zh-CN/docs/Glossary/%E7%AB%8B%E5%8D%B3%E6%89%A7%E8%A1%8C%E5%87%BD%E6%95%B0%E8%A1%A8%E8%BE%BE%E5%BC%8F)

**IIFE**（ 立即调用函数表达式）是一个在定义时就会立即执行的 [JavaScript](https://developer.mozilla.org/en-US/docs/Glossary/JavaScript) [函数](https://developer.mozilla.org/en-US/docs/Glossary/function)。

```js
(function () {
    statements
})();
```

这是一个被称为 [自执行匿名函数](https://developer.mozilla.org/en-US/docs/Glossary/Self-Executing_Anonymous_Function) 的设计模式，主要包含两部分。第一部分是包围在 [`圆括号运算符`](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Operators/Grouping) `()` 里的一个匿名函数，这个匿名函数拥有独立的词法作用域。这不仅避免了外界访问此 IIFE 中的变量，而且又不会污染全局作用域。

第二部分再一次使用 `()` 创建了一个立即执行函数表达式，JavaScript 引擎到此将直接执行函数。

## 解构赋值

[参考文档](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment)

解构赋值语法是一种 Javascript 表达式。通过解构赋值, 可以将属性/值从对象/数组中取出,赋值给其他变量。

```js
var a, b, rest;
[a, b] = [10, 20];
console.log(a); // 10
console.log(b); // 20

[a, b, ...rest] = [10, 20, 30, 40, 50];
console.log(a); // 10
console.log(b); // 20
console.log(rest); // [30, 40, 50]

({ a, b } = { a: 10, b: 20 });
console.log(a); // 10
console.log(b); // 20


// Stage 4（已完成）提案中的特性
({a, b, ...rest} = {a: 10, b: 20, c: 30, d: 40});
console.log(a); // 10
console.log(b); // 20
console.log(rest); // {c: 30, d: 40}

var x = [1, 2, 3, 4, 5];
var [y, z] = x;
console.log(y); // 1
console.log(z); // 2

var foo = ["one", "two", "three"];
var [one, two, three] = foo;
console.log(one); // "one"
console.log(two); // "two"
console.log(three); // "three"
```

解构赋值时，还可以设置默认值。

```js
let [x, y = 'b'] = ['a']; // x='a', y='b'
```

上面代码中，变量`y`解构赋值时没有取到值，所以默认值就生效了。

## Promise 对象

Promise 是 ES6 引入的封装异步操作的统一接口。它返回一个对象，包含了异步操作的信息。

Promise 本身是一个构造函数，提供了`resolve`和`reject`两个方法。一旦异步操作成功结束，就调用`resolve`方法，将 Promise 实例对象的状态改为`resolved`，一旦异步操作失败，就调用`reject`方法，将 Promise 实例的状态改成`rejected`。

```js
function timeout(duration = 0) {
  return new Promise((resolve, reject) => {
    setTimeout(resolve, duration);
  })
}
```

上面代码中，`timeout`函数返回一个 Promise 实例，在指定时间以后，将状态改为`resolved`。

```js
var p = timeout(1000).then(() => {
  return timeout(2000);
}).then(() => {
  throw new Error("hmm");
}).catch(err => {
  return Promise.all([timeout(100), timeout(200)]);
})
```

一旦 Promise 实例的状态改变以后，就可以使用`then()`方法指定下面将要执行的函数，`catch()`方法用来处理`rejected`状态的情况。

## module

ES6 意义最重大的语法变化，就是引入了模块（module）。

一个模块内部，使用`export`命令输出对外的接口。

```js
// lib/math.js
export function sum(x, y) {
  return x + y;
}
export var pi = 3.141593;
```

上面的模块输出了`sum`和`pi`两个接口。

`import`命令用于引入该模块。

```js
// app.js
import * as math from "lib/math";
alert("2π = " + math.sum(math.pi, math.pi));
```

上面代码中，`*`表示引入所有接口，也可以只引入指定的接口。

```js
// otherApp.js
import {sum, pi} from "lib/math";
alert("2π = " + sum(pi, pi));
```

## Array.prototype.filter()

[参考文档](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/filter)

### 语法

```js
var newArray = arr.filter(callback(element[, index[, array]])[, thisArg])
```

###  参数

- `callback`

  用来测试数组的每个元素的函数。返回 `true` 表示该元素通过测试，保留该元素，`false` 则不保留。它接受以下三个参数：

  `element`数组中当前正在处理的元素。`index`可选正在处理的元素在数组中的索引。`array`可选调用了 `filter` 的数组本身。

- `thisArg`可选

  执行 `callback` 时，用于 `this` 的值。

### 返回值

一个新的、由通过测试的元素组成的数组，如果没有任何数组元素通过测试，则返回空数组。

举个🌰

```js
function isBigEnough(element) {
  return element >= 10;
}
var filtered = [12, 5, 8, 130, 44].filter(isBigEnough);
// filtered is [12, 130, 44] 
```

## 常见问题

- this、作用域、优先级等综合考察：

```js
function Foo() {
  getName = function () { alert (1); };
  return this;
}
Foo.getName = function () { alert (2);};
Foo.prototype.getName = function () { alert (3);};
var getName = function () { alert (4);};
function getName() { alert (5);}
 
//请写出以下输出结果：
Foo.getName();
getName();
Foo().getName();
getName();
new Foo.getName();
new Foo().getName();
new new Foo().getName();
```

[答案](https://mp.weixin.qq.com/s/X40KEH37cRj01a_AuTzKrw)