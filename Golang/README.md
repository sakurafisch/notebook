# Golang 笔记

## 在本地启动文档

通过运行 `godoc -http :8000`，可以在本地启动文档。

> 访问 [localhost:8000/pkg](localhost:8000/pkg)，将看到系统上安装的所有包。  
> 
> 浏览 http://localhost:8000/pkg/testing/ 将会看到一些有用的内容。

## CLI

```bash
go build -v  # 编译代码
go clean  # 清除编译文件
go fmt  # 格式化代码
go get  # 动态获取远程代码包
go install  # 安装某个包
go test  # 读取 *_test.go ，生成并运行测试用的可执行文件
```

## 关键字速览

```go
break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var
```

## 内置基础类型

布尔值的类型为`bool`，值是`true`或`false`，默认为`false`。

整数类型有无符号和带符号两种。Go同时支持`int`和`uint`，这两种类型的长度相同，但具体长度取决于不同编译器的实现。Go里面也有直接定义好位数的类型：`rune`, `int8`, `int16`, `int32`, `int64`和`byte`, `uint8`, `uint16`, `uint32`, `uint64`。其中`rune`是`int32`的别称，`byte`是`uint8`的别称。

浮点数的类型有 `float32` 和 `float64` 两种（没有`float`类型），默认是 `float64` 。

复数类型。它的默认类型是 `complex128`（64位实数+64位虚数）。如果需要小一些的，也有 `complex64` (32位实数+32位虚数)。复数的形式为 `RE + IMi` ，其中 `RE` 是实数部分， `IM `是虚数部分，而最后的 `i` 是虚数单位。

字符串类型 `string` ，采用 `UTF-8` 字符集编码。字符串是用一对双引号（`""`）或反引号（``` ```）括起来定义，它的类型是 `string` 。字符串是不可变的，但可进行切片操作：

```go
s := "hello"
s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
fmt.Printf("%s\n", s)
```

错误类型 `error` ，专门用来处理错误信息。Go的`package`里面还专门有一个包`errors`来处理错误。

## 变量声明

```go
const constantName = value  // 定义常量
const Pi float32 = 3.1415926  // //如果需要，也可以明确指定常量的类型：
```

常量可以指定相当多的小数位数， 若指定给float32自动缩短为32bit，指定给float64自动缩短为64bit。

```go
var variableName type  // 定义一个名称为“variableName”，类型为"type"的变量
var vname1, vname2, vname3 type  // 定义三个类型都是“type”的变量
var variableName type = value  // 初始化“variableName”的变量为“value”值，类型是“type”
var vname1, vname2, vname3 type= v1, v2, v3  // 定义三个类型都是"type"的变量,并分别初始化
```

简短声明语法为 `varName := value` ，只能用在函数内部。

> Variables declared without a corresponding initialization are *zero-valued*. For example, the zero value for an `int` is `0`.

若变量在声明时没有赋初值，它的初值将为 零值（zero-value），比如：

```go
var a int
fmt.Println(a)  // 0
```

`_`（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。

已声明但未使用的变量会在编译阶段报错。

## iota枚举

Go里面有一个关键字`iota`，这个关键字用来声明`enum`的时候采用，它默认开始值是0，const中每增加一行加1：

```go
package main

import (
	"fmt"
)

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
}
```

## array

写在前面：Go 语言中数组、字符串和切片三者是密切相关的数据结构。这三种数据类型，在底层原始数据有着相同的内存结构，在上层，因为语法的限制而有着不同的行为表现。

`array`就是数组，它的定义方式如下：

```go
var arr [n]type
```

数组可以使用另一种`:=`来声明

```go
a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组
b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
c := [...]int{4, 5, 6} // 自动根据元素个数来计算长度
```

由于长度也是数组类型的一部分，因此`[3]int`与`[4]int`是不同的类型。

当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。

### 数组指针

```go
var a = [...]int{1, 2, 3} // a 是一个数组
var b = &a                // b 是指向数组的指针

fmt.Println(a[0], a[1])   // 打印数组的前2个元素
fmt.Println(b[0], b[1])   // 通过数组指针访问数组元素的方式和数组类似
```

### 多维数组

```go
// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

// 上面的声明可以简化，直接忽略内部的类型
easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
```

## slice

`slice`并不是真正意义上的动态数组，而是一个引用类型。

`slice`总是指向一个底层`array`，`slice`的声明也可以像`array`一样，只是不需要长度。

> `slice`和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用`...`自动计算长度，而声明`slice`时，方括号内没有任何字符。

```go
var fslice []int  // 和声明array一样，只是少了长度
slice := []byte {'a', 'b', 'c', 'd'}  // 声明一个slice，并初始化数据
```

`slice`可以从一个数组或一个已经存在的`slice`中再次声明。`slice`通过`array[i:j]`来获取，其中`i`是数组的开始位置，`j`是结束位置，但不包含`array[j]`，它的长度是`j-i`。

slice有一些简便的操作：

- `slice`的默认开始位置是0，`ar[:n]`等价于`ar[0:n]`
- `slice`的第二个序列默认是数组的长度，`ar[n:]`等价于`ar[n:len(ar)]`
- 如果从一个数组里面直接获取`slice`，可以这样`ar[:]`，因为默认第一个序列是0，第二个是数组的长度，即等价于`ar[0:len(ar)]`

`slice`是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值。

从概念上面来说`slice`像一个结构体，这个结构体包含了三个元素：

- 一个指针，指向数组中`slice`指定的开始位置
- 长度，即`slice`的长度
- 最大长度，也就是`slice`开始位置到数组的最后位置的长度

```go
Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
Slice_a := Array_a[2:5]
```

上面代码的真正存储结构如下图所示

[![img](https://github.com/astaxie/build-web-application-with-golang/raw/master/zh/images/2.2.slice2.png?raw=true)](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/images/2.2.slice2.png?raw=true)

`slice`有几个有用的内置函数：

- `len` 获取`slice`的长度
- `cap` 获取`slice`的最大容量
- `append` 向`slice`里面追加一个或者多个元素，然后返回一个和`slice`一样类型的`slice`
- `copy` 函数`copy`从源`slice`的`src`中复制元素到目标`dst`，并且返回复制的元素的个数

> 注：`append`函数会改变`slice`所引用的数组的内容，从而影响到引用同一数组的其它`slice`。 但当`slice`中没有剩余空间（即`(cap-len) == 0`）时，此时将动态分配新的数组空间。返回的`slice`数组指针将指向这个空间，而原数组的内容将保持不变；其它引用此数组的`slice`则不受影响。