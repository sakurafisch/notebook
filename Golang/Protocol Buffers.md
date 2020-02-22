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