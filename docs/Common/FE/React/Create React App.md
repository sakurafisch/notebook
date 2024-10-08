# Create React App

[Create React App](https://create-react-app.dev/docs/getting-started)

[yarn cli](https://classic.yarnpkg.com/zh-Hans/docs/cli/remove)

## 新建工程

```bash
yarn create react-app my-app --template typescript
cd my-app
yarn start
```

## 常用命令

```shell
yarn test
yarn build
```

## 引入webpack（不建议）

由于yarn会自动安装webpack依赖，若手动再安装一个，可能引起版本冲突。

```shell
yarn add --dev webpack webpack-cli webpack-dev-server # 引入 webpack
```

[更多配置](https://blog.envylabs.com/getting-started-with-webpack-2-ed2b86c68783)

## 引入antd

```shell
yarn add antd # 引入 antd
```

## 引入mobx

```shell
yarn add mobx mobx-react # 引入 mobx
```

要启用 ESNext 的装饰器 (可选), 参见下面。

CDN:

- https://unpkg.com/mobx/lib/mobx.umd.js
- https://cdnjs.com/libraries/mobx

```shell
yarn add --dev @babel/plugin-proposal-decorators
yarn add --dev @babel/plugin-proposal-class-properties
# 还需要额外的配置
```

## 引入Less

```shell
yarn add --dev less
```

