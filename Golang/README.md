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

## Panic和Recover

Go没有像Java那样的异常机制，它不能抛出异常，而是使用了`panic`和`recover`机制。

应当把它作为最后的手段来使用，也就是说，你的代码中应当没有，或者很少有`panic`的东西。

- ### Panic

> 是一个内建函数，可以中断原有的控制流程，进入一个`panic`状态中。当函数`F`调用`panic`，函数F的执行被中断，但是`F`中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方，`F`的行为就像调用了`panic`。这一过程继续向上，直到发生`panic`的`goroutine`中所有调用的函数返回，此时程序退出。`panic`可以直接调用`panic`产生。也可以由运行时错误产生，例如访问越界的数组。

- ### Recover

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

method 可以定义在任何内置类型、struct等各种类型上面。

method 也是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。

method 可以重写，类似于匿名字段。

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

### Buffered Channels

可以指定 channel 的缓冲大小

```go
ch := make(chan type, value)
```

当 value = 0 时，channel 是无缓冲阻塞读写的，当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。

> 缓冲 channel 是异步的，除非 channel 已满，否则发送或接收消息将不会等待。

## Range和Close

可以通过range，像操作slice或者map一样操作缓存类型的channel

```
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

 select选择准备就绪的第一个channel并从中接收（或发送给它）。如果准备好一个以上的channels，则它将随机选择要接收的channel。如果没有一个channel准备就绪，该语句将阻塞直到一个可用。

select语句通常用于实现超时：

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