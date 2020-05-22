# Create React App

[参考文档](https://create-react-app.dev/docs/getting-started)

## 新建工程

```bash
npx create-react-app my-app --template typescript
# or
yarn create react-app my-app --template typescript
cd my-app
npm start
```

## 引入webpack

```shell
npm install --save-dev webpack webpack-cli # 引入 webpack
```

## 引入antd

```shell
yarn add antd # 引入 antd
```

## 引入mobx

```shell
npm install --save mobx mobx-react # 引入 mobx
```

要启用 ESNext 的装饰器 (可选), 参见下面。

CDN:

- https://unpkg.com/mobx/lib/mobx.umd.js
- https://cdnjs.com/libraries/mobx

```shell
npm install --save-dev @babel/plugin-proposal-decorators
npm install --save-dev @babel/plugin-proposal-class-properties
# 还需要额外的配置
```

