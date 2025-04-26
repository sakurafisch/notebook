# Protobuf

Protobuf 是 [Protocol Buffers](https://developers.google.com/protocol-buffers) 的简称。Protobuf中最基本的数据单元是message，是类似Go语言中结构体的存在。在message中可以嵌套message或其它的基础数据类型的成员。

## 安装

1. 下载并安装[官方的protoc工具](https://github.com/google/protobuf/releases)

2. 安装代码生成插件

```go
go get github.com/golang/protobuf/protoc-gen-go
```

## 源代码示例

```protobuf
// hello.proto

syntax = "proto3";  // 采用proto3的语法

package main;

message String {
    string value = 1;
}
```

第三版的Protobuf对语言进行了提炼简化，所有成员均采用类似Go语言中的零值初始化（不再支持自定义默认值），因此消息成员也不再需要支持required特性。

message关键字定义一个新的String类型，在最终生成的Go语言代码中对应一个String结构体。String类型中只有一个字符串类型的value成员，该成员编码时用1编号代替名字。

## 生成代码

```bash
$ protoc --go_out=. hello.proto
```

其中`go_out`参数告知protoc编译器去加载对应的protoc-gen-go工具，然后通过该工具生成代码，生成代码放到当前目录。最后是一系列要处理的protobuf文件的列表。

这里只生成了一个hello.pb.go文件，其中String结构体内容如下：

```go
type String struct {
    Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
    return fileDescriptor_hello_069698f99dd8f029, []int{0}
}

func (m *String) GetValue() string {
    if m != nil {
        return m.Value
    }
    return ""
}
```

生成的结构体中还会包含一些以`XXX_`为名字前缀的成员，我们已经隐藏了这些成员。同时String类型还自动生成了一组方法，其中ProtoMessage方法表示这是一个实现了proto.Message接口的方法。此外Protobuf还为每个成员生成了一个Get方法，Get方法不仅可以处理空指针类型，而且可以和Protobuf第二版的方法保持一致（第二版的自定义默认值特性依赖这类方法）。

## 使用所生成的代码

基于新的String类型，我们可以重新实现HelloService服务：

```go
type HelloService struct{}

func (p *HelloService) Hello(request *String, reply *String) error {
    reply.Value = "hello:" + request.GetValue()
    return nil
}
```

其中Hello方法的输入参数和输出的参数均改用Protobuf定义的String类型表示。因为新的输入参数为结构体类型，因此改用指针类型作为输入参数，函数的内部代码同时也做了相应的调整。

## 定制代码生成插件

Protobuf的protoc编译器是通过插件机制实现对不同语言的支持。

Go语言的protoc-gen-go插件里面又实现了一层静态插件系统。比如protoc-gen-go内置了一个gRPC插件，用户可以通过`--go_out=plugins=grpc`参数来生成gRPC相关代码，否则只会针对message生成相关代码。

参考gRPC插件的代码，可以发现generator.RegisterPlugin函数可以用来注册插件。插件是一个generator.Plugin接口：

```go
// A Plugin provides functionality to add to the output during
// Go code generation, such as to produce RPC stubs.
type Plugin interface {
    // Name identifies the plugin.
    Name() string  // Name方法返回插件的名字，和protoc插件的名字无关
    
    // Init is called once after data structures are built but before code generation begins.
    Init(g *Generator)  // g参数中包含Proto文件的所有信息。
    
    // Generate方法用于生成主体代码。
    Generate(file *FileDescriptor)

    // GenerateImports用于生成对应的导入包代码。
    // It is called after Generate.
    GenerateImports(file *FileDescriptor)
}
```

可以设计一个netrpcPlugin插件，用于为标准库的RPC框架生成代码：

```go
import (
    "github.com/golang/protobuf/protoc-gen-go/generator"
)

type netrpcPlugin struct{ *generator.Generator }  // 内置匿名的`*generator.Generator`成员

// Name方法返回插件的名字。
func (p *netrpcPlugin) Name() string                { return "netrpc" }

// 插件从g参数对象继承了全部的公有方法
func (p *netrpcPlugin) Init(g *generator.Generator) { p.Generator = g }

func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
    if len(file.Service) > 0 {
        p.genImportCode(file)  // 调用自定义的genImportCode函数生成导入代码
    }
}

func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
    for _, svc := range file.Service {
        p.genServiceCode(svc)  // 调用自定义的genServiceCode方法生成每个服务的代码。
    }
}
```

目前，自定义的genImportCode和genServiceCode方法只是输出一行简单的注释：

```go
func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
    p.P("// TODO: import code")
}

func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
    p.P("// TODO: service code, Name = " + svc.GetName())
}
```

要使用该插件需要先通过generator.RegisterPlugin函数注册插件，可以在init函数中完成：

```go
func init() {
    generator.RegisterPlugin(new(netrpcPlugin))
}
```

因为Go语言的包只能静态导入，我们无法向已经安装的protoc-gen-go添加我们新编写的插件。我们将重新克隆protoc-gen-go对应的main函数：

```go
package main

import (
    "io/ioutil"
    "os"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/protoc-gen-go/generator"
)

func main() {
    g := generator.New()

    data, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        g.Error(err, "reading input")
    }

    if err := proto.Unmarshal(data, g.Request); err != nil {
        g.Error(err, "parsing input proto")
    }

    if len(g.Request.FileToGenerate) == 0 {
        g.Fail("no files to generate")
    }

    g.CommandLineParameters(g.Request.GetParameter())

    // Create a wrapped version of the Descriptors and EnumDescriptors that
    // point to the file that defines them.
    g.WrapTypes()

    g.SetPackageNames()
    g.BuildTypeNameMap()

    g.GenerateAllFiles()

    // Send back the results.
    data, err = proto.Marshal(g.Response)
    if err != nil {
        g.Error(err, "failed to marshal output proto")
    }
    _, err = os.Stdout.Write(data)
    if err != nil {
        g.Error(err, "failed to write output proto")
    }
}
```

为了避免对protoc-gen-go插件造成干扰，我们将我们的可执行程序命名为protoc-gen-go-netrpc，表示包含了netrpc插件。然后用以下命令重新编译hello.proto文件：

```bash
$ protoc --go-netrpc_out=plugins=netrpc:. hello.proto
```

其中`--go-netrpc_out`参数告知protoc编译器加载名为protoc-gen-go-netrpc的插件，插件中的`plugins=netrpc`指示启用内部唯一的名为netrpc的netrpcPlugin插件。在新生成的hello.pb.go文件中将包含增加的注释代码。

至此，手工定制的Protobuf代码生成插件可以工作了。

