# 使用nrm快速切换npm源

nrm 是一个 NPM 源管理器，允许你快速地在如下 NPM 源间切换。

## Install

```
sudo npm install -g nrm
```

## 使用

### 列出可用的源

```bash
$nrm ls
  npm ---- https://registry.npmjs.org/
  cnpm --- http://r.cnpmjs.org/
  taobao - http://registry.npm.taobao.org/
  eu ----- http://registry.npmjs.eu/
  au ----- http://registry.npmjs.org.au/
  sl ----- http://npm.strongloop.com/
  nj ----- https://registry.nodejitsu.com/
  pt ----- http://registry.npmjs.pt/
```

### 切换

```bash
$nrm use taobao
   Registry has been set to: http://registry.npm.taobao.org/
```

### 增加源

```bash
$nrm add <registry> <url> [home]
```

### 删除源

```bash
$nrm del <registry>
```

### 测试速度

```bash
$nrm test
```