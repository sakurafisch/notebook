# Kotlin

[Kotlin 官网](https://kotlinlang.org/)

[使用 Kotlin 开发 Android 应用](https://developer.android.com/kotlin)

[2015年6月，InfoQ推断Kotlin可能替代Java在Android的地位。](https://www.infoq.cn/article/2015/06/Android-JVM-JetBrains-Kotlin/)

[2017年5月，谷歌支持使用 Kotlin 开发安卓应用。](https://techcrunch.com/2017/05/17/google-makes-kotlin-a-first-class-language-for-writing-android-apps/)当时，Google在随后的主题演讲中指出，这只是一种附加语言，不能替代其现有的Java和C ++支持。

2017年11月，在第一届Kotlin专题大会KotlinConf上，Kotlin 首席设计师 Andrey Breslav 宣布Kotlin将支持iOS开发和Web开发，这意味着Kotlin 向全平台开发迈出了重要的一步。

[2019年5月，谷歌把 Kotlin 作为开发安卓应用的首选语言。](https://techcrunch.com/2019/05/07/kotlin-is-now-googles-preferred-language-for-android-app-development/)“我们宣布，我们迈出的下一个重要步骤是我们将Kotlin放在首位，” Android首席倡导者Chet Haase说。Haase说：“我们知道现在并不是每个人都在Kotlin上，但是我们相信您应该到达那里。” “可能有正当的理由使您仍然使用C ++和Java编程语言，这完全可以。这些都不会消失。”

Kotlin包括对Java本身当前不支持的[许多功能的](https://kotlinlang.org/docs/reference/comparison-to-java.html)支持。



## 变量声明

Kotlin 使用两个不同的关键字（即 `val` 和 `var`）来声明变量。

- `val` 用于值从不更改的变量。您不能为使用 `val` 声明的变量重新赋值。
- `var` 用于值可以更改的变量。

注意：Kotlin 没有关键字`new`。

## 变量类型

`Int` 是一种表示整数的类型，它是可以在 Kotlin 中表示的众多数值类型之一。与其他语言类似，您也可以使用 `Byte`、`Short`、`Long`、`Float` 和 `Double`，具体取决于您的数值数据。

Kotlin 在数字和字符串相加时，选择了作为字符串拼接处理，这一点跟 Java 一致。

## 类型推断

当您为 `languageName` 赋予初始值时，Kotlin 编译器可根据所赋值的类型来推断类型。Kotlin 是一种静态类型的语言。这意味着，类型在编译时解析且从不改变。

## Null 安全

## 动态类型(not for JVM)


```kotlin
var name: dynamic = ...
```

## 多行字符串

```kotlin
// 多行字符串
val s1 = """This is a
multi-line string."""

```

## 字符串插值

```kotlin
// 字符串插值
val apples = 3
val oranges = 5
val fruitSummary = "I have ${apples + oranges}  pieces of fruit."
```

## 类型检查

```kotlin
val test = Square(BigDecimal("5.2"), "square")
if (test is Square) {
    println("test is square.")
}
if (test !is Double) {
    println("test is not square.")
}

// 获取类型
println(test.javaClass.kotlin);
```

## 范围操作符

```kotlin
val names = arrayOf("Anna", "Alex", "Brian", "Jack") 
val count = names.count() 
for (i in 0..count - 1) {     
    println("Person ${i + 1} is called ${names[i]}") 
} 
// Person 1 is called Anna
// Person 2 is called Alex
// Person 3 is called Brian
// Person 4 is called Jack
```

## 集合

在 Dart，Kotlin 和 Swift 中，数组或者列表都是集合类型。Kotlin 更接近 Java，数组定长内容可变，列表变长内容可变(mutable)或者可不变(immutable)，默认是不可变。

## 数组

```kotlin
val shoppingList = mutableListOf("catfish", "water", "tulips", "blue paint") 
shoppingList[1] = "bottle of water"
// shoppingList[4] = "rock" // 编译错误，index 越界
shoppingList.add("rock") // OK

println(shoppingList) // 输出 [catfish, bottle of water, tulips, blue paint, rock]

val shoppingArray = arrayOf("catfish", "water", "tulips", "blue paint") 
shoppingArray[1] = "bottle of water"

println(shoppingArray) // 不会直接输出数组元素，需要另外处理
```

> Kotlin 的 集合类型除了 array ，其他都是默认 immutable 的。所以如果你需要修改集合，需要使用 mutableListOf ，mutableSetOf 或者mutableMapOf 。
> Kotlin 的 List 添加元素 API 是 add 。打印输出用 [] 注记。

### Set

```kotlin
val s = mutableSetOf<Int>()
println("s.isEmpty() is ${s.isEmpty()}")

s.add(1)
s.add(2)
s.add(1)

// 或
val s = mutableSetOf<Int>(1, 2, 1)

println(s)              // 输出只包含1, 2 的形式，如 [1, 2]
```

> Kotlin 的 Set 的添加 API 也是 add ，注记（打印）形式也是 [] ，声明时初始化是把元素放在函数的圆括号内。

### Map（字典）

```kotlin
val occupations = mutableMapOf("Malcolm" to "Captain", "Kaylee" to "Mechanic")
occupations["Jayne"] = "Public Relations"

println(occupations) // 输出 {Malcolm=Captain, Kaylee=Mechanic, Jayne=Public Relations}
```

## 控制流

 Kotlin 用增强版的 when 取代了 switch ，if 和 when 不仅可做控制语句，也可作表达式使用。

```kotlin
// When
val id = 5
when (id) {
  1 -> print("id == 1")
  2 -> print("id == 2")
else -> {
  print("id is undefined")
  }
}

// Inclusive for loop
for (index in 1..5) {}

// range match
val nb = 42 
when (nb) {     
    in 0..7, 8, 9 -> println("single digit")
    10 -> println("double digits")
    in 11..99 -> println("double digits")
    in 100..999 -> println("triple digits")
    else -> println("four or more digits")
}

// Cascade notation (..)
// N/A，可以用 Builder 模式或者 apply, also 函数 
```

## 函数

```kotlin
// 定义
fun greet(name: String, day: String): String {
    return "Hello $name, today is $day."
}
// 调用
greet("Bob", "Tuesday")
```

### 可变数量的参数

```kotlin
fun sumOf(vararg numbers: Int): Int {     
    var sum = 0
    for (number in numbers) {         
        sum += number     
    }     
    return sum
}
sumOf(42, 597, 12)
```

> 一个 Kotlin 函数只能有一个可变数量。对于顺序，没有强制约束，可以放在普通参数之前或者之后。当同时包含可变参数和普通参数时，调用时普通参数需要使用命名参数的调用语法。

```kotlin
fun sumOfNumbers(vararg numbers: Double, initialSum: Double): Double {
    var sum = initialSum
    for(number in numbers) {
        sum += number
    }
    return sum
}
sumOfNumbers(1.5, 2.5, initialSum=100.0) // Result = 104.0
```

### 函数类型

在 Kotlin 中函数是一等公民。

```kotlin
fun makeIncrementer(): (Int) -> Int {     
    val addOne = fun(number: Int): Int {         
        return 1 + number     
    }     
    return addOne 
} 
val increment = makeIncrementer() 
increment(7)

 // makeIncrementer 可以写成更短的形式
fun makeIncrementer2() = fun(number: Int) = 1 + number
```

### 命名参数

Kotlin 类似 Python，纯命名参数方式可以随意调整顺序，意味着开发者记忆 API 不需要记顺序。

```kotlin
fun area(width: Int, height: Int) = width * height
area(width = 2, height = 3)

// 混合顺序参数和命名参数的写法
area(2, height = 3)

// 全部采用命名参数时可以按任意顺序书写
area(height = 3, width = 2)
```

## Lambda 表达式、闭包

```kotlin
// Lambda 表达式标准形式
val sum: (Int, Int) -> Int = { x: Int, y: Int -> x + y }
println(sum(1, 2))

// 拖尾 lambda 表达式
// 如果函数的最后一个参数是函数，那么作为相应参数传入的 lambda 表达式可以放在圆括号之外：
val product = items.fold(1) { acc, e -> acc * e }

// 唯一参数 lambda 表达式
// 如果该 lambda 表达式是调用时唯一的参数，那么圆括号可以完全省略：
run { println("...") }
```

## 类

Kotlin 的花样特别多，分为主从两类构造器，主构造器必有 (没声明的话会默认生成无参的 public 主构造器) ，主构造器隐含初始化块等。写法上主构造器是直接跟在类头。

### 声明和定义

```kotlin
class Shape {
    var numberOfSides = 0
    fun simpleDescription() =
        "A shape with $numberOfSides sides."
}

var shape = Shape()
shape.numberOfSides = 7
var shapeDescription = shape.simpleDescription()
```

### 继承

```kotlin
open class NamedShape(val name: String) {
    var numberOfSides = 0

    open fun simpleDescription() =
        "A shape with $numberOfSides sides."
}

class Square(var sideLength: BigDecimal, name: String) :
        NamedShape(name) {
    init {
        numberOfSides = 4
    }

    fun area() = sideLength.pow(2)

    override fun simpleDescription() =
        "A square with sides of length $sideLength."
}

val test = Square(BigDecimal("5.2"), "square")
println(test.area())
println(test.simpleDescription())
```

## 映射

```kotlin
val numbers = listOf(20, 19, 7, 12)
numbers.map { 3 * it }
```

## 排序

```kotlin
listOf(1, 5, 3, 12, 2).sorted()
```

## 接口（协议）编程

```kotlin
interface Nameable {
    fun name(): String
}

fun <T: Nameable> f(x: T) {
    println("Name is " + x.name())
}

class NameX : Nameable {
    override fun name() : String {
        return "NameX";
    }
}

f(x = NameX())
```