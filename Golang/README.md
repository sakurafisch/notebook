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

## 类型转换

表达式 `T(v)` 将值 `v` 转换为类型 `T`。

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

```go
i := 42
f := float64(i)
u := uint(f)
```

## 类型推导

在定义一个变量但不指定其类型时（使用没有类型的 `var` 或 `:=` 语句）， 变量的类型由右值推导得出。

当右值定义了类型时，新变量的类型与其相同：

```go
var i int
j := i // j 也是一个 int
```

但是当右边包含了未指名类型的数字常量时，新的变量就可能是 `int` 、 `float64` 或 `complex128`。 这取决于常量的精度：

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

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

因为数组的长度是数组类型的一个部分，不同长度或不同类型的数据组成的数组都是不同的类型，因此在Go语言中很少直接使用数组（不同长度的数组因为类型不同无法直接赋值）

`array`就是数组，它的定义方式如下：

```go
var arr [n]type
```

```go
var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6
```

我们可以用`fmt.Printf`函数提供的`%T`或`%#v`谓词语法来打印数组的类型和详细信息：

```go
fmt.Printf("b: %T\n", b)  // b: [3]int
fmt.Printf("b: %#v\n", b) // b: [3]int{1, 2, 3}
```

数组可以使用另一种`:=`来声明

```go
a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组
b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
c := [...]int{4, 5, 6} // 自动根据元素个数来计算长度
```

由于长度也是数组类型的一部分，因此`[3]int`与`[4]int`是不同的类型。

当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。

### 遍历数组

```go
    for i := range a {
        fmt.Printf("a[%d]: %d\n", i, a[i])
    }
    for i, v := range b {
        fmt.Printf("b[%d]: %d\n", i, v)
    }
    for i := 0; i < len(c); i++ {
        fmt.Printf("c[%d]: %d\n", i, c[i])
    }
```

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

## 字符串

一个字符串是一个不可改变的字节序列，字符串通常是用来包含人类可读的文本数据。和数组不同的是，字符串的元素不可修改，是一个只读的字节数组。

每个字符串的长度虽然也是固定的，但是字符串的长度并不是字符串类型的一部分。

由于Go语言的源代码要求是UTF8编码，导致Go源代码中出现的字符串面值常量一般也是UTF8编码的。源代码中的文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列。因为字节序列对应的是只读的字节序列，因此字符串可以包含任意的数据，包括byte值0。我们也可以用字符串表示GBK等非UTF8编码的数据，不过这种时候将字符串看作是一个只读的二进制数组更准确，因为`for range`等语法并不能支持非UTF8编码的字符串的遍历。

Go语言字符串的底层结构在`reflect.StringHeader`中定义：

```go
type StringHeader struct {
    Data uintptr  // 指向底层字节数组
    Len  int  // 字符串的字节的长度
}
```

字符串其实是一个结构体，因此字符串的赋值操作也就是`reflect.StringHeader`结构体的复制过程，并不会涉及底层字节数组的复制。

字符串虽然不是切片，但是支持切片操作，不同位置的切片底层也访问同一块内存数据（因为字符串是只读的，相同的字符串面值常量通常是对应同一个字符串常量）

Go语言的源文件都是采用UTF8编码。因此，Go源文件中出现的字符串面值常量一般也是UTF8编码的（对于转义字符，则没有这个限制）。提到Go字符串时，我们一般都会假设字符串对应的是一个合法的UTF8编码的字符序列。可以用内置的`print`调试函数或`fmt.Print`函数直接打印，也可以用`for range`循环直接遍历UTF8解码后的Unicode码点值。

下面的“Hello, 世界”字符串中包含了中文字符，可以通过打印转型为字节类型来查看字符底层对应的数据：

```go
fmt.Printf("%#v\n", []byte("Hello, 世界"))
```

输出的结果是：

```go
[]byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0xe4, 0xb8, 0x96, 0xe7, \
0x95, 0x8c}
```

分析可以发现`0xe4, 0xb8, 0x96`对应中文“世”，`0xe7, 0x95, 0x8c`对应中文“界”。我们也可以在字符串面值中直指定UTF8编码后的值（源文件中全部是ASCII码，可以避免出现多字节的字符）。

```go
fmt.Println("\xe4\xb8\x96") // 打印: 世
fmt.Println("\xe7\x95\x8c") // 打印: 界
```

如果不想解码UTF8字符串，想直接遍历原始的字节码，可以将字符串强制转为`[]byte`字节序列后再行遍历（这里的转换一般不会产生运行时开销）：

```go
for i, c := range []byte("世界abc") {
    fmt.Println(i, c)
}
```

或者是采用传统的下标方式遍历字符串的字节数组：

```go
const s = "\xe4\x00\x00\xe7\x95\x8cabc"
for i := 0; i < len(s); i++ {
    fmt.Printf("%d %x\n", i, s[i])
}
```

Go语言除了`for range`语法对UTF8字符串提供了特殊支持外，还对字符串和`[]rune`类型的相互转换提供了特殊的支持。

```go
fmt.Printf("%#v\n", []rune("世界"))              // []int32{19990, 30028}
fmt.Printf("%#v\n", string([]rune{'世', '界'})) // 世界
```

## slice

`slice` 并不是真正意义上的动态数组，而是一个引用类型。

`slice` 的结构定义，`reflect.SliceHeader`：

```go
type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
```

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

从 Go1.2 开始 slice 支持第三个参数用以指定其容量。

之前我们一直采用这种方式在 slice 或者 array 基础上来获取一个 slice：

```
var array [10]int
slice := array[2:4]
```

这个例子里面 slice 的容量是 8，新版本里面可以指定这个容量：

```
slice = array[2:4:7]
```

上面这个的容量就是`7-2`，即 5。这样这个产生的新的 slice 就没办法访问最后的三个元素。

如果 slice 是这样的形式`array[:i:j]`，即第一个参数为空，默认值就是0。

### 添加切片元素

内置的泛型函数`append`可以在切片的尾部追加`N`个元素：

```go
var a []int
a = append(a, 1)               // 追加1个元素
a = append(a, 1, 2, 3)         // 追加多个元素, 手写解包方式
a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
```

在容量不足的情况下，`append`的操作会导致重新分配内存，可能导致巨大的内存分配和复制数据代价。即使容量足够，依然需要用`append`函数的返回值来更新切片本身，因为新切片的长度已经发生了变化。

除了在切片的尾部追加，我们还可以在切片的开头添加元素：

```go
var a = []int{1,2,3}
a = append([]int{0}, a...)        // 在开头添加1个元素
a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片
```

在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。因此，从切片的开头添加元素的性能一般要比从尾部追加元素的性能差很多。

由于`append`函数返回新的切片，也就是它支持链式操作。我们可以将多个`append`操作组合起来，实现在切片中间插入元素：

```go
var a []int
a = append(a[:i], append([]int{x}, a[i:]...)...)     // 在第i个位置插入x
a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片
```

每个添加操作中的第二个`append`调用都会创建一个临时切片，并将`a[i:]`的内容复制到新创建的切片中，然后将临时创建的切片再追加到`a[:i]`。

可以用`copy`和`append`组合可以避免创建中间的临时切片，同样是完成添加元素的操作：

```go
a = append(a, 0)     // 切片扩展1个空间
copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置
a[i] = x             // 设置新添加的元素
```

第一句`append`用于扩展切片的长度，为要插入的元素留出空间。第二句`copy`操作将要插入位置开始之后的元素向后挪动一个位置。第三句真实地将新添加的元素赋值到对应的位置。操作语句虽然冗长了一点，但是相比前面的方法，可以减少中间创建的临时切片。

用`copy`和`append`组合也可以实现在中间位置插入多个元素(也就是插入一个切片):

```go
a = append(a, x...)       // 为x切片扩展足够的空间
copy(a[i+len(x):], a[i:]) // a[i:]向后移动len(x)个位置
copy(a[i:], x)            // 复制新添加的切片
```

稍显不足的是，在第一句扩展切片容量的时候，扩展空间部分的元素复制是没有必要的。没有专门的内置函数用于扩展切片的容量，`append`本质是用于追加元素而不是扩展容量，扩展切片容量只是`append`的一个副作用。

### 删除切片元素

根据要删除元素的位置有三种情况：从开头位置删除，从中间位置删除，从尾部删除。其中删除切片尾部的元素最快：

```go
a = []int{1, 2, 3}
a = a[:len(a)-1]   // 删除尾部1个元素
a = a[:len(a)-N]   // 删除尾部N个元素
```

删除开头的元素可以直接移动数据指针：

```go
a = []int{1, 2, 3}
a = a[1:] // 删除开头1个元素
a = a[N:] // 删除开头N个元素
```

删除开头的元素也可以不移动数据指针，但是将后面的数据向开头移动。可以用`append`原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）：

```go
a = []int{1, 2, 3}
a = append(a[:0], a[1:]...) // 删除开头1个元素
a = append(a[:0], a[N:]...) // 删除开头N个元素
```

也可以用`copy`完成删除开头的元素：

```go
a = []int{1, 2, 3}
a = a[:copy(a, a[1:])] // 删除开头1个元素
a = a[:copy(a, a[N:])] // 删除开头N个元素
```

对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用`append`或`copy`原地完成：

```go
a = []int{1, 2, 3, ...}

a = append(a[:i], a[i+1:]...) // 删除中间1个元素
a = append(a[:i], a[i+N:]...) // 删除中间N个元素

a = a[:i+copy(a[i:], a[i+1:])]  // 删除中间1个元素
a = a[:i+copy(a[i:], a[i+N:])]  // 删除中间N个元素
```

删除开头的元素和删除尾部的元素都可以认为是删除中间元素操作的特殊情况。

## map

`map`也是一种引用类型，如果两个`map`同时指向一个底层，那么一个改变，另一个也相应的改变：

`map` 也就是Python中字典的概念，它的格式为 `map[keyType]valueType`

`map` 的读取和设置也类似 `slice` 一样，通过 `key` 来操作，只是 `slice` 的 `index` 只能是 `int` 类型，而 `map` 多了很多类型。

```go
// 声明一个字典,其 key 是 string 类型，值是 int 类型,这种方式的声明需要在使用之前使用make初始化
var numbers map[string]int
// 另一种map的声明方式
numbers = make(map[string]int)
numbers["one"] = 1  //赋值
numbers["ten"] = 10 //赋值
numbers["three"] = 3

fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
// 打印出来如:第三个数字是: 3
```

使用map过程中需要注意的几点：

- `map`是无序的，每次打印出来的`map`都会不一样，它不能通过`index`获取，而必须通过`key`获取
- `map`的长度是不固定的，也就是和`slice`一样，也是一种引用类型
- 内置的`len`函数同样适用于`map`，返回`map`拥有的`key`的数量
- `map`的值可以很方便的修改，通过`numbers["one"]=11`可以很容易的把key为`one`的字典值改为`11`
- `map`和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制

`map`内置有判断是否存在`key`的方式，通过`delete`删除`map`的元素：

```go
// 初始化一个字典
rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
csharpRating, ok := rating["C#"]
if ok {
	fmt.Println("C# is in the map and its rating is ", csharpRating)
} else {
	fmt.Println("We have no rating associated with C# in the map")
}

delete(rating, "C")  // 删除key为C的元素
```

## make、new操作

`make`用于内建类型（`map`、`slice` 和`channel`）的内存分配。`new`用于各种类型的内存分配。

内建函数`new`本质上说跟其它语言中的同名函数功能一样：`new(T)`分配了零值填充的`T`类型的内存空间，并且返回其地址，即一个`*T`类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型`T`的零值。有一点非常重要：

> `new`返回指针。

内建函数`make(T, args)`与`new(T)`有着不同的功能，make只能创建`slice`、`map`和`channel`，并且返回一个有初始值(非零)的`T`类型，而不是`*T`。本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。例如，一个`slice`，是一个包含指向数据（内部`array`）的指针、长度和容量的三项描述符；在这些项目被初始化之前，`slice`为`nil`。对于`slice`、`map`和`channel`来说，`make`初始化了内部的数据结构，填充适当的值。

> `make`返回初始化后的（非零）值。

![img](https://github.com/astaxie/build-web-application-with-golang/raw/master/zh/images/2.2.makenew.png?raw=true)

## 零值

关于“零值”，所指并非是空值，而是一种“变量未填充前”的默认值，通常为0。 此处罗列 部分类型 的 “零值”:

```go
int     0
int8    0
int32   0
int64   0
uint    0x0
rune    0 //rune的实际类型是 int32
byte    0x0 // byte的实际类型是 uint8
float32 0 //长度为 4 byte
float64 0 //长度为 8 byte
bool    false
string  ""
```

## if

Go里面`if`条件判断语句中不需要括号

Go的`if`条件判断语句里面还允许声明一个变量，这个变量的作用域只能在该条件逻辑块内。

## goto

用`goto`跳转到必须在当前函数内定义的标签。标签名是大小写敏感的。

## for

```go
for expression1; expression2; expression3 {
	//...
}
```

其中`expression1`和`expression3`是变量声明或者函数调用返回值之类的，`expression2`是用来条件判断，`expression1`在循环开始之前调用，`expression3`在每轮循环结束之时调用。

有些时候需要进行多个赋值操作，由于Go里面没有`,`操作符，那么可以使用平行赋值`i, j = i+1, j-1`

忽略`expression1`和`expression3 `就是`while`的功能

`break`操作是跳出当前循环，`continue`是跳过本次循环。

`for`配合`range`可以用于读取`slice`和`map`的数据：

```go
for k,v:=range map {
	fmt.Println("map's key:",k)
	fmt.Println("map's val:",v)
}
```

由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用`_`来丢弃不需要的返回值 例如：

```go
for _, v := range map{
	fmt.Println("map's val:", v)
}
```

## switch

```
switch sExpr {
case expr1:
	some instructions
case expr2:
	some other instructions
case expr3:
	some other instructions
default:
	other code
}
```

`sExpr`和`expr1`、`expr2`、`expr3`的类型必须一致。

表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；而如果`switch`没有表达式，它会匹配`true`。

Go里面`switch`默认相当于每个`case`最后带有`break`，匹配成功后不会自动向下执行其他case，而是跳出整个`switch`, 但是可以使用`fallthrough`强制执行后面的case代码。

## func

```go
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	//这里是处理逻辑代码
	//返回多个值
	return value1, value2
}
```

最好命名返回值，因为不命名返回值，虽代码更简洁，但是生成的文档可读性差。

```go
func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A+B
	Multiplied = A*B
	return
}
```

### 变参

函数可以有不定数量的参数。为了做到这点，首先需要定义函数使其接受变参：

```go
func myfunc(arg ...int) {}
```

`arg ...int`告诉Go这个函数接受不定数量的参数。这些参数的类型全部是`int`。在函数体中，变量`arg`是一个`int`的`slice`：

```go
for _, n := range arg {
	fmt.Printf("And the number is: %d\n", n)
}
```

可变数量的参数必须是最后出现的参数，可变数量的参数其实是一个切片类型的参数。

### 传值与传指针

当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上。

Go语言中`channel`，`slice`，`map`这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变`slice`的长度，则仍需要取地址传递指针）

### 函数作为值、类型

在Go中函数也是一种变量，我们可以通过`type`来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型

```go
type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
```

可以把这个类型的函数当做值来传递。

## defer

当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。

- 延迟调用的参数会立刻生成

## Panic和Recover

Go没有像Java那样的异常机制，它不能抛出异常，而是使用了`panic`和`recover`机制。

应当把它作为最后的手段来使用，也就是说，你的代码中应当没有，或者很少有`panic`的东西。

- ### Panic

> 是一个内建函数，可以中断原有的控制流程，进入一个`panic`状态中。当函数`F`调用`panic`，函数F的执行被中断，但是`F`中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方，`F`的行为就像调用了`panic`。这一过程继续向上，直到发生`panic`的`goroutine`中所有调用的函数返回，此时程序退出。`panic`可以直接调用`panic`产生。也可以由运行时错误产生，例如访问越界的数组。

- ### Recover

  - `recover` 只在 `defer` 语句中有效。
  - 必须要和有异常的栈帧只隔一个栈帧，`recover`函数才能正常捕获异常。换言之，`recover`函数捕获的是祖父一级调用函数栈帧的异常（刚好可以跨越一层`defer`函数）！

> 是一个内建的函数，可以让进入`panic`状态的`goroutine`恢复过来。`recover`仅在延迟函数中有效。在正常的执行过程中，调用`recover`会返回`nil`，并且没有其它任何效果。如果当前的`goroutine`陷入`panic`状态，调用`recover`可以捕获到`panic`的输入值，并且恢复正常的执行。

## `main`函数和`init`函数

Go里面有两个保留的函数：`init`函数（能够应用于所有的`package`）和`main`函数（只能应用于`package main`）。

这两个函数在定义时不能有任何的参数和返回值。

虽然一个`package`里面可以写任意多个`init`函数，但这无论是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个`package`中每个文件只写一个`init`函数。

Go程序会自动调用`init()`和`main()`，所以你不需要在任何地方调用这两个函数。

每个`package`中的`init`函数都是可选的，但`package main`就必须包含一个`main`函数。

程序的初始化和执行都起始于`main`包。如果`main`包还导入了其它的包，那么就会在编译时将它们依次导入。有时一个包会被多个包同时导入，那么它只会被导入一次。

当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行`init`函数（如果有的话），依次类推。等所有被导入的包都加载完毕了，就会开始对`main`包中的包级常量和变量进行初始化，然后执行`main`包中的`init`函数（如果存在的话），最后执行`main`函数。下图详细地解释了整个执行过程：

![img](https://github.com/astaxie/build-web-application-with-golang/raw/master/zh/images/2.3.init.png?raw=true)

在 `main.main` 函数执行之前所有代码都运行在同一个 goroutine，也就是程序的主系统线程中。因此，如果某个 `init` 函数内部用 go 关键字启动了新的 goroutine 的话，新的 goroutine 只有在进入 `main.main` 函数之后才可能被执行到。

## import

我们在写Go代码的时候经常用到import这个命令用来导入包文件，而我们经常看到的方式参考如下：

```go
import(
    "fmt"
)
```

上面这个fmt是Go语言的标准库，其实是去`GOROOT`环境变量指定目录下去加载该模块。

Go的import还支持如下两种方式来加载自己写的模块：

- 相对路径
```go
import “./model” //当前文件同一目录的model目录，但是不建议这种方式来import
```
- 绝对路径
```go
import “shorturl/model” //加载gopath/src/shorturl/model模块
```
上面展示了一些import常用的几种方式，但是还有一些特殊的import：

- 点操作

这个包导入之后，调用这个包的函数时，可以省略前缀的包名

```
import(
    . "fmt"
)
```

- 别名操作

把包命名成另一个名字

调用包函数时前缀变成自定义前缀，即 `f.Println("hello world")`

```go
import(
    f "fmt"
)
```

- _ 操作

_ 操作引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。

```go
import (
    "database/sql"
    _ "github.com/ziutek/mymysql/godrv"
)
```

## Struct

```go
type person struct {
	name string
	age int
}
```

### 匿名字段（嵌入字段）

Go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。

当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。

- 匿名字段能够实现字段的继承。
- 最外层的优先访问。
- 自定义类型、内置类型都可以作为匿名字段，而且可以在相应的字段上面进行函数操作（如append）。
- 所有继承来的方法的接收者参数依然是那个匿名成员本身，而不是当前的变量。

```go
package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human  // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func main() {
	// 我们初始化一个学生
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
}
```

Student访问属性age和name的时候，就像访问自己所有用的字段一样。

student还能访问Human这个字段作为字段名：

```go
mark.Human = Human{"Marcus", 55, 220}
mark.Human.age -= 1
```

所有的内置类型和自定义类型都是可以作为匿名字段，而不仅仅是struct字段

```go
package main

import "fmt"

type Skills []string

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human  // 匿名字段，struct
	Skills // 匿名字段，自定义的类型string slice
	int    // 内置类型作为匿名字段
	speciality string
}

func main() {
	// 初始化学生Jane
	jane := Student{Human:Human{"Jane", 35, 100}, speciality:"Biology"}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}
```

## method

method的语法如下：

```go
func (r ReceiverType) funcName(parameters) (results)
```

**Receiver以值传递不会改变原对象，以指针传递会改变原对象。**

- 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
- method里面可以访问接收者的字段
- 调用method通过`.`访问，就像struct里面访问字段一样
- 每种类型对应的方法必须和类型的定义在同一个包中
- 对于给定的类型，每个方法的名字必须是唯一的，方法不支持重载。

method 可以定义在任何内置类型、struct等各种类型上面。

method 也是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。

method 可以重写，类似于匿名字段。

### 方法表达式

方法表达式的特性可以将方法还原为普通类型的函数：

```go
// 不依赖具体的文件对象
// func CloseFile(f *File) error
var CloseFile = (*File).Close

// 不依赖具体的文件对象
// func ReadFile(f *File, offset int64, data []byte) int
var ReadFile = (*File).Read

// 文件处理
f, _ := OpenFile("foo.dat")
ReadFile(f, 0, data)
CloseFile(f)
```

## Interface

interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现。

```go
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}
```

如果我们定义一个interface的变量，那么这个变量里面可以存实现这个interface的任意类型的对象。因为m能够持有这三种类型的对象。

### 空 interface

任意的类型都实现了空interface(即 interface{})，也就是包含0个method的interface。

空interface对于描述起不到任何的作用(因为它不包含任何的method），但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。它有点类似于C语言的void*类型。

一个函数把interface{}作为参数，那么它可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。

### interface 函数参数

interface的变量可以持有任意实现该interface类型的对象，这给我们编写函数(包括method)提供了一些额外的思考：可以通过定义interface参数，让函数接受各种类型的参数。

举个🌰：fmt.Println是我们常用的一个函数，但是你是否注意到它可以接受任意类型的数据。打开fmt的源码文件，可以看到这样一个定义:

```go
type Stringer interface {
	 String() string
}
// Stringer 接口只有一个 String 方法，因此只要实现 String 方法即可实现 Stringer 接口。
```

也就是说，任何实现了String方法的类型都能作为参数被fmt.Println调用。

注：实现了error接口的对象（即实现了Error() string的对象），使用fmt输出时，会调用Error()方法，因此不必再定义String()方法了。

### interface变量存储的类型

我们知道interface的变量里面可以存储任意类型的数值(该类型实现了interface)。那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：

- Comma-ok断言，可以直接判断是否是该类型的变量。

```go
 value, ok = element.(T)
```

这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。

如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。

- switch测试

`element.(type)`语法不能在switch外的任何逻辑里面使用，如果要在switch外面判断一个类型就使用`comma-ok`。

### 嵌入interface

功能类似于 Struct 的匿名字段：如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method。

举个🌰

```go
// 在源码包 container/heap 的一个定义
type Interface interface {
	sort.Interface  // 嵌入字段sort.Interface
	Push(x interface{})  // a Push method to push elements into the heap
	Pop() interface{}  // a Pop elements that pops elements from the heap
}
```

sort.Interface其实就是嵌入字段，把sort.Interface的所有method给隐式的包含进来了。也就是下面三个方法：

```go
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less returns whether the element with index i should sort
	// before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```

再举一个🌰

io包下面的 io.ReadWriter ，它包含了io包下面的Reader和Writer两个interface：

```go
// io.ReadWriter
type ReadWriter interface {
	Reader
	Writer
}
```

## 反射

反射机制不仅包括要能在运行时对程序自身信息进行检测，还要求程序能进一步根据这些信息改变程序状态或结构。我们一般用到的包是reflect包。参考 [laws of reflection](http://golang.org/doc/articles/laws_of_reflection.html)

> **元编程**（英语：Metaprogramming），又译**超编程**，是指某类[计算机程序](https://zh.wikipedia.org/wiki/计算机程序)的编写，这类计算机程序编写或者操纵其它程序（或者自身）作为它们的数据，或者在[运行时](https://zh.wikipedia.org/wiki/运行时)完成部分本应在[编译时](https://zh.wikipedia.org/wiki/编译时)完成的工作。多数情况下，与手工编写全部代码相比，程序员可以获得更高的工作效率，或者给与程序更大的灵活度去处理新的情形而无需重新编译。
>
> 编写元程序的语言称之为[元语言](https://zh.wikipedia.org/wiki/元語言)。被操纵的程序的语言称之为“[目标语言](https://zh.wikipedia.org/w/index.php?title=目标语言&action=edit&redlink=1)”。一门编程语言同时也是自身的元语言的能力称之为“[反射](https://zh.wikipedia.org/wiki/反射_(计算机科学))”或者“自反”。
>
> 反射是促进元编程的一种很有价值的语言特性。把编程语言自身作为一级数据类型（如[LISP](https://zh.wikipedia.org/wiki/LISP)、[Forth](https://zh.wikipedia.org/wiki/Forth)或[Rebol](https://zh.wikipedia.org/wiki/Rebol)）也很有用。支持[泛型编程](https://zh.wikipedia.org/wiki/泛型编程)的语言也使用元编程能力。

使用reflect一般分成三步:

要去反射某个类型的值(这些值都实现了空interface)，首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)。这两种获取方式如下:

```go
t := reflect.TypeOf(i)    //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
v := reflect.ValueOf(i)   //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
```

转化为reflect对象之后我们就可以进行一些操作了，也就是将reflect对象转化成相应的值，例如:

```go
tag := t.Elem().Field(0).Tag  //获取定义在struct里面的标签
name := v.Elem().Field(0).String()  //获取存储在第一个字段里面的值
```

获取反射值能返回相应的类型和数值

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
```

最后，**反射的字段必须是可修改的**。反射的字段必须是可读写的意思是，如果下面这样写，那么会发生错误

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
v.SetFloat(7.1)
```

如果要修改相应的值，必须这样写

```go
var x float64 = 3.4
p := reflect.ValueOf(&x)
v := p.Elem()
v.SetFloat(7.1)
```

## goroutine

```go
go hello(a, b, c)
```

多个goroutine运行在同一个进程里面，共享内存数据，不过设计上我们要遵循：不要通过共享来通信，而要通过通信来共享。

> runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。

> 默认情况下，在Go 1.5将标识并发系统线程个数的runtime.GOMAXPROCS的初始值由1改为了运行环境的CPU核数。

但在Go 1.5以前调度器仅使用单线程，也就是说只实现了并发。想要发挥多核处理器的并行，需要在我们的程序中显式调用 runtime.GOMAXPROCS(n) 告诉调度器同时使用多个线程。GOMAXPROCS 设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。如果n < 1，不会改变当前设置。

在Go语言中，同一个Goroutine线程内部，顺序一致性内存模型是得到保证的。但是不同的Goroutine之间，并不满足顺序一致性内存模型，需要通过明确定义的同步事件来作为同步的参考。如果两个事件不可排序，那么就说这两个事件是并发的。为了最大化并行，Go语言的编译器和处理器在不影响上述规定的前提下可能会对执行语句重新排序（CPU也会对一些指令进行乱序执行）。

## channels

channel 默认是无缓冲的。

channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。这些值只能是特定的类型：channel类型。

无缓冲的 channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。

> 无缓冲的 channel 是同步的；两端中任意一端的 channel 将等到另一端准备好为止。

定义一个channel时，也需要定义发送到channel的值的类型。必须使用make 创建channel：

```
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```

channel通过操作符`<-`来接收和发送数据

```
ch <- v    // 发送v到channel ch.
v := <-ch  // 从ch中接收数据，并赋值给v
```

在无缓存的Channel上的每一次发送操作都有与其对应的接收操作相配对，发送和接收操作通常发生在不同的Goroutine上（在同一个Goroutine上执行2个操作很容易导致死锁）。**无缓存的Channel上的发送操作总在对应的接收操作完成前发生.**

**对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。**

### Buffered Channels

可以指定 channel 的缓冲大小

```go
ch := make(chan type, value)
```

当 value = 0 时，channel 是无缓冲阻塞读写的，当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。

> 缓冲 channel 是异步的，除非 channel 已满，否则发送或接收消息将不会等待。

对于带缓冲的Channel，**对于Channel的第`K`个接收完成操作发生在第`K+C`个发送操作完成之前，其中`C`是Channel的缓存大小。**

## Range和Close

发送者可以 `close` 一个 channel 来表示再没有值会被发送了。接收者可以通过赋值语句的第二参数来测试 channel 是否被关闭：

```gp
v, ok := <-ch
```

循环 `for i := range c` 会不断从 channel 接收值，直到它被关闭。

可以通过range，像操作slice或者map一样操作缓存类型的channel：

```go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
```

`for i := range c`能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显式的关闭channel，生产者通过内置函数`close`关闭channel。关闭channel之后就无法再发送任何数据了，在消费方可以通过语法`v, ok := <-ch`测试channel是否被关闭。如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。

> 记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic

> channel不像文件之类的，不需要经常去关闭，只有当确实没有任何发送数据了，或者想显式的结束range循环之类的

## select

`select` 语句使得一个 goroutine 在多个通讯操作上等待。

`select` 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个。

`select` 选择准备就绪的第一个channel并从中接收（或发送给它）。如果准备好一个以上的channels，则它将随机选择要接收的channel。如果没有一个channel准备就绪，该语句将阻塞直到一个可用。

`select` 语句通常用于实现超时：

```go
select {
case msg1 := <- c1:
  fmt.Println("Message 1", msg1)
case msg2 := <- c2:
  fmt.Println("Message 2", msg2)
case <- time.After(time.Second):
  fmt.Println("timeout")
}
```

`time.After` 创建一个频道，并在给定的持续时间后发送当前时间。（我们对时间不感兴趣，所以我们没有将其存储在变量中。）

我们还可以指定一个 `default` 情况：

当 `select` 中的其他条件分支都没有准备好的时候，`default` 分支会被执行。

```go
select {
case msg1 := <- c1:
  fmt.Println("Message 1", msg1)
case msg2 := <- c2:
  fmt.Println("Message 2", msg2)
case <- time.After(time.Second):
  fmt.Println("timeout")
default:
  fmt.Println("nothing ready")
}
```

## runtime goroutine

runtime包中有几个处理goroutine的函数：

- Goexit

  退出当前执行的goroutine，但是defer函数还会继续调用

- Gosched

  让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。

- NumCPU

  返回 CPU 核数量

- NumGoroutine

  返回正在执行和排队的任务总数

- GOMAXPROCS

  用来设置可以并行计算的CPU核数的最大值，并返回之前的值。
  

> Go的运行时还包含了其自己的调度器，这个调度器使用了一些技术手段，可以在n个操作系统线程上多工调度m个Goroutine。Go调度器的工作和内核的调度是相似的，但是这个调度器只关注单独的Go程序中的Goroutine。Goroutine采用的是半抢占式的协作调度，只有在当前Goroutine发生阻塞时才会导致调度；同时发生在用户态，调度器会根据具体函数只保存必要的寄存器，切换的代价要比系统线程低得多。运行时有一个`runtime.GOMAXPROCS`变量，用于控制当前运行正常非阻塞Goroutine的系统线程数目。
>

## 指针随时可能会变

不要假设变量在内存中的位置是固定不变的，指针随时可能会变。

**不能随意将指针保持到数值变量中，Go语言的地址也不能随意保存到不在GC控制的环境中，因此使用CGO时不能在C语言中长期持有Go语言对象的地址，因为：**

Go语言支持递归调用。Go语言函数的递归调用深度逻辑上没有限制，函数调用的栈是不会出现溢出错误的，因为Go语言运行时会根据需要动态地调整函数栈的大小。每个goroutine刚启动时只会分配很小的栈（4或8KB，具体依赖实现），根据需要动态调整栈的大小，栈最大可以达到GB级（依赖具体实现，在目前的实现中，32位体系结构为250MB,64位体系结构为1GB）。在Go1.4以前，Go的动态栈采用的是分段式的动态栈，通俗地说就是采用一个链表来实现动态栈，每个链表的节点内存位置不会发生变化。但是链表实现的动态栈对某些导致跨越链表不同节点的热点调用的性能影响较大，因为相邻的链表节点它们在内存位置一般不是相邻的，这会增加CPU高速缓存命中失败的几率。为了解决热点调用的CPU缓存命中率问题，Go1.4之后改用连续的动态栈实现，也就是采用一个类似动态数组的结构来表示栈。不过连续动态栈也带来了新的问题：当连续栈动态增长时，需要将之前的数据移动到新的内存空间，这会导致之前栈中全部变量的地址发生变化。**虽然Go语言运行时会自动更新引用了地址变化的栈变量的指针，但指针不再是固定不变的了。

## 原子操作

所谓的原子操作就是并发编程中“最小的且不可并行化”的操作。

通常，如果多个并发体对同一个共享资源进行的操作是原子的话，那么同一时刻最多只能有一个并发体对该资源进行操作。

一般情况下，原子操作都是通过“互斥”访问来保证的，通常由特殊的CPU指令提供保护。

写在前面：互斥锁的代价比普通整数的原子读写高很多，在性能敏感的地方可以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能。

### 互斥锁 `sync.Mutex` 

```go
import (
    "sync"
)

// 定义结构体并且立即用于变量初始化
var total struct {
    sync.Mutex
    value int
}

func worker(wg *sync.WaitGroup) {
    defer wg.Done()

    for i := 0; i <= 100; i++ {
        total.Lock()  // 进入临界区前加锁
        total.value += i  // 临界区，修改共享变量 total.value
        total.Unlock()  // 退出临界区后解锁
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go worker(&wg)
    go worker(&wg)
    wg.Wait()

    fmt.Println(total.value)
}
```

### `sync/atomic`包

`atomic.Value`原子对象提供了`Load`和`Store`两个原子方法，分别用于加载和保存数据，返回值和参数都是`interface{}`类型。

`atomic.AddUint64`函数调用保证了`total`的读取、更新和保存是一个原子操作，因此在多线程中访问也是安全的。

```go
import (
    "sync"
    "sync/atomic"
)

var total uint64

func worker(wg *sync.WaitGroup) {
    defer wg.Done()

    var i uint64
    for i = 0; i <= 100; i++ {
        atomic.AddUint64(&total, i)  // 官方支持的原子操作
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go worker(&wg)
    go worker(&wg)
    wg.Wait()
}
```

### 标准库中`sync.Once`的实现

```go
type Once struct {
    m    Mutex
    done uint32
}

func (o *Once) Do(f func()) {
    if atomic.LoadUint32(&o.done) == 1 {
        return
    }

    o.m.Lock()
    defer o.m.Unlock()

    if o.done == 0 {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}
```

### 基于`sync.Once`重新实现单件模式

```go
var (
    instance *singleton
    once     sync.Once
)

func Instance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```

## 控制并发数

在Go语言自带的godoc程序实现中有一个`vfs`的包对应虚拟的文件系统，在`vfs`包下面有一个`gatefs`的子包，`gatefs`子包的目的就是为了控制访问该虚拟文件系统的最大并发数。

`gatefs`包的应用很简单：

```go
import (
    "golang.org/x/tools/godoc/vfs"
    "golang.org/x/tools/godoc/vfs/gatefs"
)

func main() {
    fs := gatefs.New(vfs.OS("/path"), make(chan bool, 8))
    // ...
}
```

其中`vfs.OS("/path")`基于本地文件系统构造一个虚拟的文件系统，然后`gatefs.New`基于现有的虚拟文件系统构造一个并发受控的虚拟文件系统。

并发数控制的原理，就是通过带缓存管道的发送和接收规则来实现最大并发阻塞。不过`gatefs`对此做一个抽象类型`gate`，增加了`enter`和`leave`方法分别对应并发代码的进入和离开。当超出并发数目限制的时候，`enter`方法会阻塞直到并发数降下来为止。

```go
type gate chan bool

func (g gate) enter() { g <- true }
func (g gate) leave() { <-g }
```

`gatefs`包装的新的虚拟文件系统就是将需要控制并发的方法增加了`enter`和`leave`调用而已：

```go
type gatefs struct {
    fs vfs.FileSystem
    gate
}

func (fs gatefs) Lstat(p string) (os.FileInfo, error) {
    fs.enter()
    defer fs.leave()
    return fs.fs.Lstat(p)
}
```

我们不仅可以控制最大的并发数目，而且可以通过带缓存Channel的使用量和最大容量比例来判断程序运行的并发率。当管道为空的时候可以认为是空闲状态，当管道满了时任务是繁忙状态。

## 通知 goroutine 结束

### 关闭管道以广播退出指令

通过`close`来关闭`cancel`管道向多个Goroutine广播退出的指令。

当每个Goroutine收到退出指令退出时一般会进行一定的清理工作，但是退出的清理工作并不能保证被完成，因为`main`线程并没有等待各个工作Goroutine退出工作完成的机制。不过我们可以结合使用`sync.WaitGroup`来等待这些工作的完成。

```go
func worker(wg *sync.WaitGroup, cannel chan bool) {
    defer wg.Done()

    for {
        select {
        default:
            fmt.Println("hello")
        case <-cannel:
            return
        }
    }
}

func main() {
    cancel := make(chan bool)

    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker(&wg, cancel)
    }

    time.Sleep(time.Second)
    close(cancel)
    wg.Wait()
}
```

### Context 包

在Go1.7发布时，标准库增加了一个`context`包，用来简化对于处理单个请求的多个Goroutine之间与请求域的数据、超时和退出等操作。

我们可以用`context`包来重新实现前面的线程安全退出或超时的控制:

```go
func worker(ctx context.Context, wg *sync.WaitGroup) error {
    defer wg.Done()

    for {
        select {
        default:
            fmt.Println("hello")
        case <-ctx.Done():
            return ctx.Err()
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker(ctx, &wg)
    }

    time.Sleep(time.Second)
    cancel()

    wg.Wait()
}
```

当并发体超时或`main`主动停止工作者Goroutine时，每个工作者都可以安全退出。

## 错误处理

那些将运行失败看作是预期结果的函数，它们会返回一个额外的返回值，通常是最后一个来传递错误信息。如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，通常被命名为ok。比如，当从一个`map`查询一个结果时，可以通过额外的布尔值判断是否成功：

```go
if v, ok := m["key"]; ok {
    return v
}
```

但是导致失败的原因通常不止一种，很多时候用户希望了解更多的错误信息。如果只是用简单的布尔类型的状态值将不能满足这个要求。在C语言中，默认采用一个整数类型的`errno`来表达错误，这样就可以根据需要定义多种错误类型。在Go语言中，`syscall.Errno`就是对应C语言中`errno`类型的错误。在`syscall`包中的接口，如果有返回错误的话，底层也是`syscall.Errno`错误类型。

比如我们通过`syscall`包的接口来修改文件的模式时，如果遇到错误我们可以通过将`err`强制断言为`syscall.Errno`错误类型来处理：

```go
err := syscall.Chmod(":invalid path:", 0666)
if err != nil {
    log.Fatal(err.(syscall.Errno))
}
```

### 用 `recover` 把 `panic` 转化为输出错误信息

错误被认为是一种可以预期的结果；而异常则是一种非预期的结果，发生异常可能表示程序中存在BUG或发生了其它不可控的问题。

但是对于那些提供类似Web服务的框架而言；它们经常需要接入第三方的中间件。因为第三方的中间件是否存在BUG是否会抛出异常，Web框架本身是不能确定的。为了提高系统的稳定性，Web框架一般会通过`recover`来防御性地捕获所有处理流程中可能产生的异常，然后将异常转为普通的错误返回。

以JSON解析器为例，说明recover的使用场景。考虑到JSON解析器的复杂性，即使某个语言解析器目前工作正常，也无法肯定它没有漏洞。因此，当某个异常出现时，我们不会选择让解析器崩溃，而是会将panic异常当作普通的解析错误，并附加额外信息提醒用户报告此错误。

```go
func ParseJSON(input string) (s *Syntax, err error) {
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("JSON: internal error: %v", p)
        }
    }()
    // ...parser...
}
```

标准库中的`json`包，在内部递归解析JSON数据的时候如果遇到错误，会通过抛出异常的方式来快速跳出深度嵌套的函数调用，然后由最外一级的接口通过`recover`捕获`panic`，然后返回相应的错误信息。

Go语言库的实现习惯: 即使在包内部使用了`panic`，但是在导出函数时会被转化为明确的错误值。

## Go访问C内存创建大于2GB的内存

因为Go语言实现的限制，我们无法在Go语言中创建大于2GB内存的切片（具体请参考makeslice实现代码）。不过借助cgo技术，我们可以在C语言环境创建大于2GB的内存，然后转为Go语言的切片使用：

```go
package main

/*
#include <stdlib.h>

void* makeslice(size_t memsize) {
    return malloc(memsize);
}
*/
import "C"
import "unsafe"

func makeByteSlize(n int) []byte {
    p := C.makeslice(C.size_t(n))
    return ((*[1 << 31]byte)(p))[0:n:n]
}

func freeByteSlice(p []byte) {
    C.free(unsafe.Pointer(&p[0]))
}

func main() {
    s := makeByteSlize(1<<32+1)
    s[len(s)-1] = 255
    print(s[len(s)-1])
    freeByteSlice(s)
}
```

例子中我们通过 `makeByteSlize` 来创建大于4G内存大小的切片，从而绕过了Go语言实现的限制。而 `freeByteSlice` 辅助函数则用于释放从C语言函数创建的切片。

因为C语言内存空间是稳定的，基于C语言内存构造的切片也是绝对稳定的，不会因为Go语言栈的变化而被移动。

## 图片

参考 [文档](http://golang.org/pkg/image/#Image)

[Package image](http://golang.org/pkg/image/#Image) 定义了 `Image` 接口：

```go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle  // 此处的 Rectangle 为 image.Rectangle
    At(x, y int) color.Color
}
```

*注意*：`Bounds` 方法的 `Rectangle` 返回值实际上是一个 [`image.Rectangle`](http://golang.org/pkg/image/#Rectangle)， 其定义在 `image` 包中。

`color.Color` 和 `color.Model` 也是接口，但是通常因为直接使用预定义的实现 `image.RGBA` 和 `image.RGBAModel` 而被忽视了。这些接口和类型由[image/color 包](http://golang.org/pkg/image/color/)定义。