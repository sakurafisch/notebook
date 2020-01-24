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

## 一键安装脚本
```shell
# 一键安装 VS Code 插件&工具
echo on
code --install-extension ms-vscode.go
cd %GOPATH%
go get -u github.com/mdempsky/gocode
go get -u github.com/uudashr/gopkgs/cmd/gopkgs
go get -u github.com/ramya-rao-a/go-outline
go get -u github.com/acroca/go-symbols
go get -u github.com/cweill/gotests/...
go get -u github.com/fatih/gomodifytags
go get -u github.com/josharian/impl
go get -u github.com/davidrjenni/reftools/cmd/fillstruct
go get -u github.com/haya14busa/goplay/cmd/goplay
go get -u github.com/godoctor/godoctor
go get -u github.com/stamblerre/gocode
go get -u github.com/rogpeppe/godef
go get -u github.com/sqs/goreturns
go get -u golang.org/x/tools/cmd/guru
go get -u golang.org/x/tools/cmd/gorename
go get -u github.com/go-delve/delve/cmd/dlv
go get -u github.com/stamblerre/gocode
go get -u golang.org/x/lint/golint
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
pause
exit
```
