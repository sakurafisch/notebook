# React Hook

```js
import React, { useState } from 'react';
```

*Hook* 是 React 16.8 的新增特性。它可以让你在不编写 class 的情况下使用 state 以及其他的 React 特性。

+ Hook 的本质是 JavaScript 函数

+ Hook是[向后兼容](https://reactjs.org/docs/hooks-intro.html#no-breaking-changes)的。
+ Hook在class内部不起作用。



## 使用规则

+ 只在最顶层使用 Hook
+ 只在 React 函数中调用 Hook

使用 ESLint 插件强制执行这两条规则

```sh
npm install eslint-plugin-react-hooks --save-dev
```

```json
// 你的 ESLint 配置
{
  "plugins": [
    // ...
    "react-hooks"
  ],
  "rules": {
    // ...
    "react-hooks/rules-of-hooks": "error", // 检查 Hook 的规则
    "react-hooks/exhaustive-deps": "warn" // 检查 effect 的依赖
  }
}
```



## 示例：一个简单的计数器

```js
import React, { useState } from 'react';

function Example() {
  // 声明一个叫 “count” 的 state 变量。
  const [count, setCount] = useState(0);  // 这里使用了Hook

  // Similar to componentDidMount and componentDidUpdate:
  useEffect(() => {
    // Update the document title using the browser API
    document.title = `You clicked ${count} times`;
  });

  return (
    <div>
      <p>You clicked {count} times</p>
      <button onClick={() => setCount(count + 1)}>
        Click me
      </button>
    </div>
  );
}
```

```js
const [count, setCount] = useState(0);
```

这里的 `useState` 就是一个Hook。



## useState

使用 [解构赋值](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment#Array_destructuring) 语法。

```js
const [count, setCount] = useState(0);
```

+ `useState()` 接受唯一的参数作为初始 state。该参数可以是数字、字符串或对象。
+ `useState()` 返回值为：第一个值是当前的 state，第二个值是更新 state 的函数。
+ React 在重复渲染时保留这个 state。

+ 修改 state 时，不会把新的 state 和旧的 state 进行合并。.
+ 在函数结束时，state 仍然保留。



## use Effect

可以把 `useEffect` Hook 看做 `componentDidMount`，`componentDidUpdate` 和 `componentWillUnmount` 这三个函数的组合。

+ 当 React 渲染组件时，会保存已使用的 effect，并在更新完 DOM 后执行它。
+ 每次我们重新渲染，都会生成*新的* effect，替换掉之前的。

### 无需清除的 effect 示例

```js
function Example() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    document.title = `You clicked ${count} times`;
  });
}
```



