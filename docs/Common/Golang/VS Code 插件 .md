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

## 配置代理

[仓库链接](https://github.com/goproxyio/goproxy)

[参考文档](https://goproxy.io/zh/)

确保 Go 版本是1.13 及以上

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
# or
go env -w GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct


# 设置不走 proxy 的私有仓库，多个用逗号相隔（可选）
go env -w GOPRIVATE=*.corp.example.com
```

设置完上面几个环境变量后，您的 `go` 命令将从公共代理镜像中快速拉取您所需的依赖代码了。[私有库的支持请看这里](https://goproxy.io/zh/docs/goproxyio-private.html)。

如果网络还是有问题，试试清空 module 缓存

```bash
go clean --modcache
```

## 一键安装脚本

```shell
# 一键安装 VS Code 插件&工具
echo on
code --install-extension ms-vscode.go
cd %GOPATH%
go get -u -v golang.org/x/tools/gopls@latest
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

## Makefile

```makefile
build:
    go build -o bin/main main.go

run:
    go run main.go

compile:
    GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o /bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o /bin/main-windows-386 main.go
```

## 框架

**Beego** 面向Go编程语言的开源高性能web框架

- [https://github.com/astaxie/beego](https://link.zhihu.com/?target=https%3A//github.com/astaxie/beego)
- [https://beego.me](https://link.zhihu.com/?target=https%3A//beego.me)

**Buffalo** 使用Go语言快速构建Web应用

- [https://github.com/gobuffalo/buffalo](https://link.zhihu.com/?target=https%3A//github.com/gobuffalo/buffalo)
- [https://gobuffalo.io](https://link.zhihu.com/?target=https%3A//gobuffalo.io)

**Echo** 高性能、极简Go语言Web框架

- [https://github.com/labstack/echo](https://link.zhihu.com/?target=https%3A//github.com/labstack/echo)
- [https://echo.labstack.com](https://link.zhihu.com/?target=https%3A//echo.labstack.com)

**Gin** Go语言编写的HTTP Web框架，它以更好的性能实现了类似Martini的API，性能更好

- [https://github.com/gin-gonic/gin](https://link.zhihu.com/?target=https%3A//github.com/gin-gonic/gin)
- [https://gin-gonic.github.io/gin](https://link.zhihu.com/?target=https%3A//gin-gonic.github.io/gin)

**Iris** 全宇宙最快的Go语言Web框架，完备MVC支持，拥抱未来

- [https://github.com/kataras/iris](https://link.zhihu.com/?target=https%3A//github.com/kataras/iris)
- [https://iris-go.com](https://link.zhihu.com/?target=https%3A//iris-go.com)

**Revel** Go语言的高效、全栈Web框架

- [https://github.com/revel/revel](https://link.zhihu.com/?target=https%3A//github.com/revel/revel)
- [https://revel.github.io](https://link.zhihu.com/?target=https%3A//revel.github.io)

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
- [社区文档](https://learnku.com/go/docs)
- [Golang JP](https://plus.google.com/communities/107075098212007233819) - Japan.
- [Golang Korea](https://www.facebook.com/groups/golangko/about/) - Korea.
- [Golang Taiwan](http://golang.tw/) - Taiwan.



