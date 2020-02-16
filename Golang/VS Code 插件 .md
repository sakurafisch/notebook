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
go get -u -v github.com/mdempsky/gocode
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v github.com/cweill/gotests/...
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/josharian/impl
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
go get -u -v github.com/haya14busa/goplay/cmd/goplay
go get -u -v github.com/godoctor/godoctor
go get -u -v github.com/stamblerre/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/go-delve/delve/cmd/dlv
go get -u -v github.com/stamblerre/gocode
go get -u -v golang.org/x/lint/golint
go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint
go get -u -v golang.org/x/tools/cmd/goimports
go get -u -v golang.org/x/tools/cmd/godoc
pause
exit
```

## Where to Go from here...

你可以从 [安装 Go](http://golang.org/doc/install/) 开始，或者下载 [Go App Engine SDK](http://code.google.com/appengine/downloads.html#Google_App_Engine_SDK_for_Go).

一旦安装了 Go， 应当继续阅读的内容。 它包含了参考、指南、视频等等更多资料。

了解如何组织 Go 代码并在其上工作，参阅 [这个视频](http://www.youtube.com/watch?v=XCsL89YtqCs)，或者阅读 [如何编写 Go 代码](http://golang.org/doc/code.html)。

在标准库上需要帮助的话，参考 [包手册](http://golang.org/pkg/)。语言本身的帮助，阅读 [语言规范](http://golang.org/ref/spec)是件令人愉快的事情。

进一步探索 Go 的并发模型，参阅 [Go 并发模型](http://www.youtube.com/watch?v=f6kdp27TYZs) ([幻灯片](http://talks.golang.org/2012/concurrency.slide)) 以及 [深入 Go 并发模型](https://www.youtube.com/watch?v=QDDwwePbDtw) ([幻灯片](http://talks.golang.org/2013/advconc.slide)) 并且阅读 [使用通讯共享内存](http://golang.org/doc/codewalk/sharemem/) 的代码之旅。

想要开始编写 Web 应用，参阅 [一个简单的编程环境](http://vimeo.com/53221558) ([幻灯片](http://talks.golang.org/2012/simple.slide)) 并且阅读 [编写 Web 应用](http://golang.org/doc/articles/wiki/) 的指南.

[GO 中的一等公民函数](http://golang.org/doc/codewalk/functions/) 展示了有趣的函数类型。

[Go Blog](http://blog.golang.org/) 有着众多的关于 Go 的文章信息。

[mikespook 的博客](http://www.mikespook.com/tag/golang/)有大量中文的关于 Go 的文章和翻译。

开源电子书 [Go Web 编程](https://github.com/astaxie/build-web-application-with-golang) 和 [Go入门指南](https://github.com/Unknwon/the-way-to-go_ZH_CN) 能够帮助你更加深入的了解和学习 Go 语言。

访问 [golang.org](http://golang.org/) 了解更多内容。

## 社区

- [Golang China](http://groups.google.com/group/golang-china) - China.
- [Golang JP](https://plus.google.com/communities/107075098212007233819) - Japan.
- [Golang Korea](https://www.facebook.com/groups/golangko/about/) - Korea.
- [Golang Taiwan](http://golang.tw/) - Taiwan.