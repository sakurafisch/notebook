# Helidon

## Prerequisites

Helidon 需要配置 Java 11 和 Maven，详情可参考 [这里](https://helidon.io/docs/latest/#/about/03_prerequisites)

Docker 和 Kubernetes 是可选的配置。

## 社区热度

当前 `Helidon` 的社区热度不如 `Quarkus` 和 `Micronaut`。(2021年8月29日)

但是 `Helidon` 是由 `Oracle` 主导开发的。它充分使用了 Java 的语言特性。因此值得关注。

## SE 和 MP 两种风格

它目前分为两个系列：

- Helidon SE
- Helidon MP

其中 Helidon SE 的写法比较像 gin 和 koa，而 Helidon MP 的写法比较像 Spring。

SE 只是一个薄薄的封装，适合小型的开发；而 MP 是稍重的封装，功能较全也臃肿一些。

## CLI

CLI 用法可参考 [这里](https://helidon.io/docs/latest/#/about/05_cli)

### 安装

#### zsh

```zsh
curl -O https://helidon.io/cli/latest/linux/helidon
chmod +x ./helidon
sudo mv ./helidon /usr/local/bin/
```

#### powershell

```powershell
PowerShell -Command Invoke-WebRequest -Uri "https://helidon.io/cli/latest/windows/helidon.exe" -OutFile "C:\Windows\system32\helidon.exe"
```

### 基本用法

```zsh
helidon init
# Then answer the questions.
cd <myproject>
helidon dev
```

