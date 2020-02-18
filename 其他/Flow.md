# Flow

## 概述

- Flow is a static type checker for JavaScript code. 
- Flow checks your code for errors through **static type annotations**.

## 安装

```shell
yarn add --dev @babel/core @babel/cli @babel/preset-flow # Setup Compiler Babel
```

.babelrc

```json
{
  "presets": ["@babel/preset-flow"]
}
```

package.json

```json
{
  "name": "my-project",
  "main": "lib/index.js",
  "scripts": {
    "build": "babel src/ -d lib/",
    "prepublish": "yarn run build"
  }
}
```

```shell
yarn add --dev flow-bin  # Setup Flow
yarn run flow init  # 初始化
yarn run flow  # 启动
```

## 新建工程

```shell
flow init    # Initialize Project
flow status  # Run the Flow Background Process
```

