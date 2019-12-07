# VS Code 插件

```shell
code --install-extension ms-vscode.go
```

第一次用 VS Code 打开 Go 文件时，它会提示缺少分析工具，应该点击按钮来安装它们。



在 VS Code 中调试 Go 代码建议使用 Delve。可以通过以下 go get 命令安装：

```go
go get -u github.com/go-delve/delve/cmd/dlv
```



对默认的语法检查进行增强可以用 [GolangCI-Lint](https://github.com/golangci/golangci-lint)。可通过以下方式安装：

```shell
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
```

