# Dart

## 关键字速览

| [abstract](https://www.dartcn.com/guides/language/language-tour#抽象类) 2 | [dynamic](https://www.dartcn.com/guides/language/language-tour#重要的概念) 2 | [implements](https://www.dartcn.com/guides/language/language-tour#隐式接口) 2 | [show](https://www.dartcn.com/guides/language/language-tour#导入库的一部分) 1 |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [as](https://www.dartcn.com/guides/language/language-tour#类型判定运算符) 2 | [else](https://www.dartcn.com/guides/language/language-tour#if-和-else) | [import](https://www.dartcn.com/guides/language/language-tour#使用库) 2 | [static](https://www.dartcn.com/guides/language/language-tour#类变量和方法) 2 |
| [assert](https://www.dartcn.com/guides/language/language-tour#assert) | [enum](https://www.dartcn.com/guides/language/language-tour#枚举类型) | [in](https://www.dartcn.com/guides/language/language-tour#for-循环) | [super](https://www.dartcn.com/guides/language/language-tour#扩展类继承) |
| [async](https://www.dartcn.com/guides/language/language-tour#异步支持) 1 | [export](https://www.dartcn.com/guides/libraries/create-library-packages) 2 | [interface](https://stackoverflow.com/questions/28595501/was-the-interface-keyword-removed-from-dart) 2 | [switch](https://www.dartcn.com/guides/language/language-tour#switch-和-case) |
| [await](https://www.dartcn.com/guides/language/language-tour#异步支持) 3 | [extends](https://www.dartcn.com/guides/language/language-tour#扩展类继承) | [is](https://www.dartcn.com/guides/language/language-tour#类型判定运算符) | [sync](https://www.dartcn.com/guides/language/language-tour#生成器) 1 |
| [break](https://www.dartcn.com/guides/language/language-tour#break-和-continue) | [external](https://stackoverflow.com/questions/24929659/what-does-external-mean-in-dart) 2 | [library](https://www.dartcn.com/guides/language/language-tour#库和可见性) 2 | [this](https://www.dartcn.com/guides/language/language-tour#构造函数) |
| [case](https://www.dartcn.com/guides/language/language-tour#switch-和-case) | [factory](https://www.dartcn.com/guides/language/language-tour#工厂构造函数) 2 | [mixin](https://www.dartcn.com/guides/language/language-tour#为类添加功能mixins) 2 | [throw](https://www.dartcn.com/guides/language/language-tour#throw) |
| [catch](https://www.dartcn.com/guides/language/language-tour#catch) | [false](https://www.dartcn.com/guides/language/language-tour#booleans) | [new](https://www.dartcn.com/guides/language/language-tour#使用构造函数) | [true](https://www.dartcn.com/guides/language/language-tour#booleans) |
| [class](https://www.dartcn.com/guides/language/language-tour#实例变量) | [final](https://www.dartcn.com/guides/language/language-tour#final-和-const) | [null](https://www.dartcn.com/guides/language/language-tour#默认值) | [try](https://www.dartcn.com/guides/language/language-tour#catch) |
| [const](https://www.dartcn.com/guides/language/language-tour#final-和-const) | [finally](https://www.dartcn.com/guides/language/language-tour#finally) | [on](https://www.dartcn.com/guides/language/language-tour#catch) 1 | [typedef](https://www.dartcn.com/guides/language/language-tour#typedefs) 2 |
| [continue](https://www.dartcn.com/guides/language/language-tour#break-和-continue) | [for](https://www.dartcn.com/guides/language/language-tour#for-循环) | [operator](https://www.dartcn.com/guides/language/language-tour#重写运算符) 2 | [var](https://www.dartcn.com/guides/language/language-tour#变量) |
| [covariant](https://www.dartcn.com/guides/language/sound-problems#the-covariant-keyword) 2 | [Function](https://www.dartcn.com/guides/language/language-tour#函数) 2 | [part](https://www.dartcn.com/guides/libraries/create-library-packages#organizing-a-library-package) 2 | [void](https://medium.com/dartlang/dart-2-legacy-of-the-void-e7afb5f44df0) |
| [default](https://www.dartcn.com/guides/language/language-tour#switch-和-case) | [get](https://www.dartcn.com/guides/language/language-tour#getters-和-setters) 2 | [rethrow](https://www.dartcn.com/guides/language/language-tour#catch) | [while](https://www.dartcn.com/guides/language/language-tour#while-和-do-while) |
| [deferred](https://www.dartcn.com/guides/language/language-tour#延迟加载库) 2 | [hide](https://www.dartcn.com/guides/language/language-tour#导入库的一部分) 1 | [return](https://www.dartcn.com/guides/language/language-tour#函数) | [with](https://www.dartcn.com/guides/language/language-tour#为类添加功能mixins) |
| [do](https://www.dartcn.com/guides/language/language-tour#while-和-do-while) | [if](https://www.dartcn.com/guides/language/language-tour#if-和-else) | [set](https://api.dartlang.org/stable/dart-core/Set-class.html) 2 | [yield](https://www.dartcn.com/guides/language/language-tour#生成器) 3 |

避免使用这些单词作为标识符。 但是，如有必要，标有上标的关键字可以用作标识符：

- 带有 **1** 上标的单词为 **上下文关键字**， 仅在特定位置具有含义。 他们在任何地方都是有效的标识符。
- 带有 **2** 上标的单词为 **内置标识符**， 为了简化将 JavaScript 代码移植到 Dart 的工作， 这些关键字在大多数地方都是有效的标识符， 但它们不能用作类或类型名称，也不能用作 import 前缀。
- 带有 **3** 上标的单词是与 Dart 1.0 发布后添加的[异步支持](https://www.dartcn.com/guides/language/language-tour#异步支持)相关的更新，作为限制类保留字。
  不能在标记为 `async` ，`async*` 或 `sync*` 的任何函数体中使用 `await` 或 `yield` 作为标识符。

关键字表中的剩余单词都是**保留字**。 不能将保留字用作标识符。

## 语言特性

可选命名参数、`..`（级联运算符）和`?.`（条件成员访问运算符）以及`??`（判空赋值运算符）

## 变量类型

Dart 在数字类型和字符串相加时，需要数字类型先调用`toString()`方法，这样做可作为字符串拼接处理。

## 动态类型

```dart
dynamic name = 'John';
```

## immutable binding 和 immutable object

```dart
// Dart 的 immutable binding 和 immutable object
final array = const [1, 2, 3];
const array2 = [1, 2, 3]; // 当类型为集合时，右边隐含 const，Effective Dart 建议不写。
```


## final 和 const 的区别

[Dart 中 static, final, const 区别](https://juejin.im/post/5d396a03f265da1b68370110)

[How does the const constructor actually work?](https://stackoverflow.com/questions/21744677/how-does-the-const-constructor-actually-work)

> Dart 中的 final 和 const 都能用来表明变量在一次赋值后就不想要被改变或者不需要改变。final 用于编译期不可计算的值，const 则用于编译期可计算的值。const 比 final 更“深度” —— 一个 final 的集合，其集合的元素不是 final 的；而一个 const 的集合，集合内的内容也是 const 的，即所谓的 “the object will be frozen and completely immutable (对象被冻结且完全不可变)”。 const 既可用在左边，也可用在右边，用在右边表明值不可变 (变量本身还可变)。当 const 用在左边，且左边的变量是集合，则隐含了右边的值同时也是 const 的。通过上下文可以得知表达式一定是常量，那么 const 可以被省略，也建议省略。例如：

```dart
const primaryColors = [
  Color("red", [255, 0, 0]),
  Color("green", [0, 255, 0]),
  Color("blue", [0, 0, 255]),
];
```

> 一个类的成员变量，如果是 const ，必须同时声明为 static ，即类(域)成员变量。这容易理解，因为如果是实例变量，按照设计是跟实例走，而 const 本身要求编译期可知，这就要求类的 const 变量同时也要是 static 变量。 Dart 2.5 之后，你还可以通过类型检查和转换，以及 collection if, spread 操作符来定义 (const) 常量。 const constructor 是一个难点，可以参考这个 stackoverflow 链接里的回答：其中有一个概念叫 "canonical instance"，也就是 a "canonicalized" instance。const constructor 对于还没切换到惰性初始化之前的 Dart 还蛮重要的。

## 多行字符串

```dart
// 多行字符串
var s1 = """This is a
multi-line string.""";

```

## 字符串插值

```dart
// 字符串插值
final apples = 3;
final oranges = 5;
final fruitSummary = "I have ${apples + oranges}  pieces of fruit.";
```

## 类型检查

```dart
final test = Square(5.2, "square");
if (test is Square) {
    print("test is square.");
}

if (test is! Square) {
    print("test is not square.");
}

// 获取类型
print(test.runtimeType);
```

## 集合

在 Dart，Kotlin 和 Swift 中，数组或者列表都是集合类型。Dart 和 Swift 的做法相同之处在于都没有像 Java 或者 Kotlin 那样把列表和数组细分开，不同之处是 Dart 字面名选择了 List ，而 Swift 字面名选择了 Array。

### 数组（包含List）

```dart
var emptyList = <int>[];
List fixedLengthList = new List(3);
List growableList = new List();

fixedLengthList[0] = 0;    // OK
growableList.add(0);      // OK

var complexList = [];     // 如果不指定类型，则类型是 List<dynamic>
complexList.add(1);       // OK
complexList.add("apple"); // OK，因为类型是 List<dynamic> ，
                          // 所以列表可以混合存放不同子类型 

print(complexList); // 输出 [1, "apple"]
```

> Dart 对于定长和变长的设计，容易理解。但个人对这种用数字参数来区别两种 List 的方式有点不适应，可能是因为 Java 的编程经验导致？
> Dart 的 List 添加元素 API 是 add ，打印输出用 [] 注记。

### Set

```dart
var s = Set<int>();
// 或
Set<int> s = Set<int>();
Set<int> s = {}; // 如果有显式的类型注释，右边可以直接用 {} 初始化
print("s.isEmpty is ${s.isEmpty}");
s.add(1);
s.add(2);
s.add(1);

// 或
Set<int> s = {1, 2, 1};

print(s);               // 输出只包含1, 2 的形式，如 {1, 2}   
```

Dart 的 Set 的添加 API 是 add ，注记（打印）形式是用 {} ，显式类型注释时可以直接用花括号初始化集合内容。

### Map (字典)

```dart
// 空 Map
var emptyMap = {}; // 不论是否有显式类型注释，dart 都可以用 {} 来初始化 Map，
                  // 如果不指定类型，则类型是 _InternalLinkedHashMap<dynamic, dynamic>

// Mutable Map
var occupations = {"Malcolm": "Captain", "Kaylee": "Mechanic"};
occupations["Jayne"] = "Public Relations";

print(occupations); // 输出 {Malcolm: Captain, Kaylee: Mechanic, Jayne: Public Relations}
```

## 控制流

级联注记 (Cascade notation)是一个小特色。这种书写方式最大的好处是：把一堆高度相关的代码逻辑“强制”聚在一处，使代码看起来更流畅，整洁。

Dart 的 switch 跟 Java 一样，需要主动 break。

```dart
// Switch case 
var id = 5;
switch (id) {
  case 1:
    print("id == 1");
    break;
  case 2:
    print("id == 2");
    break;
  default:
    print("id is undefined");
}

// Inclusive for loop
for (var i = 1; i <= 5; i++) {}

// range match
// N/A

// Cascade notation (..)
querySelector('#confirm')
  ..text = 'Confirm'
  ..classes.add('important')
  ..onClick.listen((e) =>
window.alert('Confirmed!'));
```

## 函数

```dart
// 定义
String greet(String name, String day) {
    return "Hello $name, today is $day.";
}
// 调用
greet("Bob", "Tuesday");
```

### 命名参数

 Dart 的命名参数同时也是可选参数。

```dart
// Dart 的命名参数同时也是可选参数，用 {} 包裹
int area({int width, int height}) {
    return width * height;
}

area(width: 2, height: 3)
// 命名参数可以按任意顺序书写
area(height: 3, width: 2)

// Dart 的可选顺序参数，用 [] 包裹
int area([int width, int height]) {
    return width * height;
}

area(); // 可编译，运行时错误
area(2); // 可编译，运行时错误
area(2, 3); // OK
```

### 函数类型

在 Dart 中函数是一等公民。

```dart
// 写法1
int Function(int) makeIncrementer(int a) {
    return (a) {
        return 1 + a;
    };
}

// 写法2
int Function(int) makeIncrementer2(int a) {
    int myInnerFunction(a) {
        return 1 + a;
    }

    final myFunction = myInnerFunction;
    return myFunction;
}

// 写法3：类型是 Closure: (int) => int
int Function(int) increment = (int a) {
    return 1 + a;
};

// 写法4：类型是 Closure: (dynamic) => num
var increment2 = (a) {
    return 1 + a;
};

print(makeIncrementer);
print(makeIncrementer(1));
print(makeIncrementer(1)());

print(increment);
print(increment(1));
```

写法 1、2 和写法 3 、4 并不对等，推荐 3、4 的写法。

留意结尾的分号，这里有没有分号结尾的差别类似于 C++ 中的声明和定义。写法1、2 仅声明，写法 3、4 声明的同时也定义了，所以必须加分号。

## Lambda 表达式、闭包

```dart
// 匿名函数标准形式
// ([[Type] param1[, …]]) { 
//  codeBlock; 
// }; 
var list = ['apples', 'bananas', 'oranges'];
list.forEach((item) {
  print('${list.indexOf(item)}: $item');
});
```

## 类

### 声明和定义

```dart
class Shape {
    var numberOfSides = 0;
    String simpleDescription() {
        return "A shape with $numberOfSides sides.";
    }
}

var shape = Shape();
shape.numberOfSides = 7;
var shapeDescription = shape.simpleDescription();
```

### 继承

```dart
class NamedShape {
    int numberOfSides = 0;
    String name;

    NamedShape(String name) {
        this.name = name;
    }
    
    // 或者采用语法糖写法
    NamedShape(this.name);

    String simpleDescription() {
        return "A shape with $numberOfSides sides.";
    }
}

class Square extends NamedShape {
    double sideLength;

    Square(this.sideLength, String name) : super(name) {
        numberOfSides = 4;
    }

    double area() {
        return sideLength * sideLength;
    }

    // Dart 官方建议审慎使用 @override 注解，当超类不在自身控制下时使用
    // 这个注解主要是为了超类更改了代码，子类不知情仍正常工作时可以提醒程序员
    @override
    String simpleDescription() {
        return "A square with sides of length $sideLength.";
    }
}

final test = Square(;5.2, "square");
print(test.area());
print(test.simpleDescription());
```

## 映射

```dart
final numbers = [20, 19, 7, 12];

// 写法1
numbers.map((number) => 3 * number).toList();

// 写法2：类似 Python 的语法
[for (var number in numbers) 3 * number];
```

## 排序

```dart
var list = [1, 5, 3, 12, 2];
list.sort();
```

## 接口（协议）编程

```dart
// 被声明为 abstract 因此不能被实例化
abstract class AbstractContainer {
  // 定义构造器, 字段, 方法...

  void updateChildren(); // Abstract method.
}

// Dart 移除了显式接口声明，
// 每个类都隐式的定义个了一个包含该类和该类实现的接口的所有实例成员 
// (说人话：类就是接口)。
// 一个类通过 implements 语句声明它们实现一个或者多个接口，
// 然后提供接口要求的 API，例如：

// Person. 包含了 greet() 方法的隐式接口
class Person {
    // 仅库可见，属于接口的一部分
    final _name;

    // 不属于接口的一部分，因为这是个构造器
    Person(this._name);

    // 属于接口的一部分
    String greet(String who) => 'Hello, $who. I am $_name.';
}

// 一个 Person 接口的实现
class Impostor implements Person {
    get _name => '';

    String greet(String who) => 'Hi $who. Do you know who I am?';
}

String greetBob(Person person) => person.greet('Bob');

void main() {
    print(greetBob(Person('Kathy')));
    print(greetBob(Impostor()));
}

// 如果想要像 Java 那样的接口，建议用抽象类作为接口，相当于是把关心的部分提炼出来
```

