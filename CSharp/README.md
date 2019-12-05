# C#笔记

## 概述

C# 编程语言允许你构建多种类型的应用程序，例如：  

- 用于捕获、分析和处理数据的业务应用程序
- 可从 web 浏览器访问的动态 web 应用程序
- 2D 和 3D 游戏
- 金融和科研应用程序
- 基于云的应用程序
- 移动应用程序

## Hello World

```c#
using System;
namespace MyNewApp
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
        }
    }
}
```

## 新手常错

新程序员常犯的错误：

- `Console` 中的 `C` 或 `WriteLine` 中的 `W` 或 `L` 使用小写字母而非大写字母
- 在 `Console` 和 `WriteLine` 间使用逗号而不是句点
- 忘记使用双引号，或使用单引号将短语 `Hello World!` 括起来
- 在命令末尾忘记使用分号

## Console.WriteLine

### 将 char 类型文本写入控制台

如果只希望将单个字母数字字符打印到屏幕，可以用单引号将一个字母数字字符括起来，以此创建 char 类型文本。  

```c#
Console.WriteLine('b');
```

### 将 int 类型文本写入控制台

如果要将数字整数（无小数位）值打印到输出，则可以使用 int 类型文本。   

```c#
Console.WriteLine(123);
```

### 将十进制文本写入控制台

如果我们想要打印包含小数点后的值的数字，可以使用“十进制文本”。  

若要创建十进制文本，请在数字后面追加字母 `m`。 在此上下文中，`m` 称为“文本后缀”。 文本后缀告知编译器你希望使用十进制类型的值。  

```c#
Console.WriteLine(12.3m);
```

### 将 bool 类型文本写入控制台

如果我们想要打印表示 `true` 或 `false` 的值，则可以使用“bool 类型文本”。

```c#
Console.WriteLine(true);
Console.WriteLine(false);
```

## 声明变量

若要创建新变量，必须首先声明变量的数据类型，然后为其指定名称。

```c#
string firstName;  // 声明变量
char userOption;
int gameScore;
decimal particlesPerMillion;
bool processedCustomer;
```

```c#
string[] fraudulentOrderIDs = new string[3];  // 声明数组
fraudulentOrderIDs[0] = "A123";  // 为数组赋值
fraudulentOrderIDs[1] = "B456";
fraudulentOrderIDs[2] = "C789";
```

```c#
string[] fraudulentOrderIDs = { "A123", "B456", "C789" };  // 声明数组并初始化
```

 若要确定数组的大小，可使用 `Length` 属性

### 变量名称规则和约定

- 变量名可包含字母数字字符和下划线字符。 
- 变量名必须以字母或下划线开头，不能以数字开头。 开发者将下划线用于特殊目的，因此现在请勿使用。
- 变量名不能是 C# 关键字。 
- 变量名区分大小写。
- 变量名应使用骆驼式命名法。
- 变量名不应包含变量的数据类型。 

### 自动推断类型

官方名称：隐式类型本地变量

`var` 关键字指示编译器根据变量的初始化值来推断变量的数据类型。

```c#
var message = "Hello world!";
```

由于变量 `message` 会被立即设置为字符串值 `"Hello World!"`，因此 C# 编译器了解其意向并将每个消息实例视为 `string` 类型实例。

## 字符串处理

### 常用内置方法

- 使用 `ToUpper()` 或 `ToLower()` 帮助程序方法，确保字符串全部大写或全部小写。

- 使用 `Trim()` 帮助程序方法，删除前导空格或尾随空格。

- [Contains(String)](https://docs.microsoft.com/zh-cn/dotnet/api/system.string.contains?view=netframework-4.8)返回一个值，该值指示指定的子串是否出现在此字符串中。

### 字符转义序列和逐字字符串

1. 转义反斜杠```\```的用法和C++一样。
2. 在文本字符串的前面使用 `@` 指令，可以创建逐字字符串

逐字字符串文本将保留所有空格和字符，而无需转义反斜杠。

示例：

```c#
Console.WriteLine(@"   c:\source\repos   
      (this is where your code goes)");
```

输出：

```
   c:\source\repos   
      (this is where your code goes)
```

> 该字符串跨越两行，并且此 C# 指令生成的空格保留在输出中。

3. 使用 `\u` 转义序列在文本字符串中添加编码字符，然后使用四个字符的代码表示 Unicode (UTF-16) 中的某些字符。

示例：

```c#
// Kon'nichiwa World
Console.WriteLine("\u3053\u3093\u306B\u3061\u306F World!");。
```

### 字符串串联

使用字符串连接运算符 `+`将两个字符串连接在一起

```c#
string firstName = "Bob";
string greeting = "Hello";
Console.WriteLine(greeting + " " + firstName + "!");	// Hello Bob!
```

### 字符串内插

> 字符串内插通过使用“模板”和一个/多个内插表达式将多个值合并为单个文本字符串。 内插表达式是一个变量，由一个左大括号和一个右大括号符号 `{ }` 括起来。 当文本字符串以 `$` 字符为前缀时，该字符串将变为模板。

```c#
string firstName = "Bob";
string greeting = "Hello";
Console.WriteLine($"{greeting} {firstName}!");	// Hello Bob!
```

### 合并逐字文本和字符串内插

假设你需要在模板中使用逐字文本。 可以同时使用逐字文本前缀符号 `@` 和字符串内插 `$` 符号。

```c#
string projectName = "First-Project";
Console.WriteLine($@"C:\Output\{projectName}\Data");	// C:\Output\First-Project\Data
```

## 数字运算

- 可以对数字执行数学等加法运算。
- 字符串串联和加法均使用加号。 这称为“重载运算符”，编译器根据其运算的数据类型推断合理的使用。
- 如果 C# 编译器意识到开发者试图连接数字的字符串表示形式来实现演示，则它会隐式将 `int` 转换为 `string`。
- 使用括号定义运算顺序，以显式指示编译器我们要在执行其他运算之前执行特定运算。

```c#
string firstName = "Bob";
int widgetsSold = 7;
Console.WriteLine(firstName + " sold " + widgetsSold + 7 + " widgets.");
// 输出:
// Bob sold 77 widgets.
```

```c#
string firstName = "Bob";
int widgetsSold = 7;
Console.WriteLine(firstName + " sold " + (widgetsSold + 7) + " widgets.");
// 输出:
// Bob sold 14 widgets.
```

### 数学运算符

- ```%``` 是取余运算符
- `+` 是加法运算符
- `-` 是减法运算符
- `*` 是乘法运算符
- `/` 是除法运算符

**除法的一些细节**

1. 要查看除法是否正确有效，我们需要使用支持小数点后的小数位的数据类型，如 `decimal`。

```c#
decimal decimalQuotient = 7.0m / 5;
Console.WriteLine("Decimal quotient: " + decimalQuotient);
// 输出:
// decimal quotient: 1.4
```

2. 为了使其正常运作，商（赋值运算符的左边）必须是类型 `decimal` 以及被除数或除数必须为类型 `decimal`（或两者都为该类型）。

3. 如果需要将类型 `int` 的两个变量相除，但不希望结果被截断，必须执行从 `int` 到 `decimal` 的数据类型强制转换

```c#
int first = 7;
int second = 5;
decimal quotient = (decimal)first / (decimal)second;
Console.WriteLine(quotient);
// 输出:
// 1.4
```

## 从 .NET 类库调用方法

> 数据类型也是 .NET 类库的一部分。

> 类被组织成不同的命名空间，以防发生命名冲突。

- 若要调用 .NET 类库中类的方法，请采用 `ClassName.MethodName()` 格式。
- 调用无状态方法时，无需先创建其类的新实例。
- 在调用有状态方法时，需要创建类的实例，并访问对象的方法。

示例：生成随机数字并将其打印到控制台来模拟掷骰子

```c#
// 如果多次运行代码，控制台输出中将显示介于 1 到 6 之间的数字。
Random dice = new Random();
int roll = dice.Next(1, 7);
Console.WriteLine(roll);
```

### 有状态方法与无状态方法

> 在计算中，“状态”用于描述特定时刻下执行环境的状况。 代码逐行执行时，值存储在变量中。 在执行过程中的任何时候，应用程序的当前状态为存储在内存中的所有值的集合。

- 某些方法不依赖于应用程序的当前状态，即可正常工作。例如，`Console.WriteLine()` 方法不依赖于内存中存储的任何值。

- 但另一些方法必须有权访问应用程序的状态，才能正常工作。有状态（实例）方法在字段中跟踪方法的状态，这些字段是在类上定义的变量。 对于存储状态的这些字段，类的每个新实例都将获取其自己的副本。

>  换言之，有状态方法是依赖于由以前的已执行代码行存储在内存中的值构建的。 或者，有状态方法通过更新值或将新值存储在内存中来修改应用程序的状态。 它们也称为实例方法。

- 单个类可支持有状态方法和无状态方法。 但是，需要调用有状态方法时，必须首先创建类的实例，这样方法才能访问状态。

### 为什么Next() 方法是有状态的？

为了创建随机错觉，`Next()` 方法的开发者决定将日期和时间具体捕获到毫秒的小数部分，并用它来设置一种算法，使其每次都会生成不同数字。在 `dice` 对象的生存期内捕获和维护的状态是种子值。 后续每次对 `Next()` 方法的调用都将重新运行该算法，但需确保种子发生更改，这样才不会返回相同的值。

## if-else语句

```c#
Random random = new Random();
int daysUntilExpiration = random.Next(12);
int discountPercentage = 0;
if (daysUntilExpiration == 0) {
    Console.WriteLine("Your subscription has expired.");
}
else if (daysUntilExpiration == 1) {
    Console.WriteLine("Your subscription expires within a day!");
    discountPercentage = 20;
}
else if (daysUntilExpiration <= 5) {
    Console.WriteLine($"Your subscription expires in {daysUntilExpiration} days.");
    discountPercentage = 10;
}
else if (daysUntilExpiration <= 10) {
    Console.WriteLine("Your subscription will expire soon.  Renew now!");
}

if (discountPercentage > 0) {
    Console.WriteLine($"Renew now and save {discountPercentage}%.");
}
```

## foreach语句

foreach语句的局限性：不能重新分配 `name` 的值，因为它是 `foreach` 迭代的内部实现的一部分。

```c#
string[] names = { "Bob", "Conrad", "Grant" };
foreach (string name in names) {
    Console.WriteLine(name);
}
```

```c#
int[] inventory = { 200, 450, 700, 175, 250 };
int sum = 0;
int bin = 0;
foreach (int items in inventory) {
    sum += items;
    bin++;
    Console.WriteLine($"Bin {bin} = {items} items (Running tally: {sum})");
}
Console.WriteLine($"We have {sum} items in inventory.");
```

```c#
string[] orderIDs = { "B123", "C234", "A345", "C15", "B177", "G3003", "C235", "B179" };
foreach (string orderID in orderIDs) {
    if (orderID.StartsWith("B")) {
        Console.WriteLine(orderID);
    }
}
```

## switch case 语句

```c#
int employeeLevel = 200;
string employeeName = "John Smith";

string title = "";

switch (employeeLevel) {
    case 100:
        title = "Junior Associate";
        break;
    case 200:
        title = "Senior Associate";
        break;
    case 300:
        title = "Manager";
        break;
    case 400:
        title = "Senior Manager";
        break;
    default:
        title = "Associate";
        break;
}
```

```c#
// SKU = Stock Keeping Unit
string sku = "01-MN-L";

string[] product = sku.Split('-');

string type = "";
string color = "";
string size = "";

switch (product[0])
{
    case "01":
        type = "Sweat shirt";
        break;
    case "02":
        type = "T-Shirt";
        break;
    case "03":
        type = "Sweat pants";
        break;
    default:
        type = "Other";
        break;
}

switch (product[1])
{
    case "BL":
        color = "Black";
        break;
    case "MN":
        color = "Maroon";
        break;
    default:
        color = "White";
        break;
}

switch (product[2])
{
    case "S":
        size = "Small";
        break;
    case "M":
        size = "Medium";
        break;
    case "L":
        size = "Large";
        break;
    default:
        size = "One Size Fits All";
        break;
}

Console.WriteLine($"Product: {size} {color} {type}");
```

## for语句

```c#
for (int i = 0; i < 10; i++) {
    Console.WriteLine(i);
}
```

```c#
string[] names = { "Alex", "Eddie", "David", "Michael" };
for (int i = 0; i < names.Length; i++)
    if (names[i] == "David") names[i] = "Sammy";

foreach (var name in names) Console.WriteLine(name); 
```

```c#
for (int i = 1; i < 101; i++) {
    if ((i % 3 == 0) && (i % 5 == 0))
        Console.WriteLine($"{i} - FizzBuzz");
    else if (i % 3 == 0)
        Console.WriteLine($"{i} - Fizz");
    else if (i % 5 == 0)
        Console.WriteLine($"{i} - Buzz");
    else
        Console.WriteLine($"{i}");
}
```

## do-while语句

`do-while` 语句至少循环访问一次代码块，并且可能会根据布尔表达式继续循环访问。 布尔表达式的计算结果通常取决于在代码块内生成或检索到的某个值。

```c#
// 不断生成 1 到 10 之间的随机数，直到生成数字 7。
Random random = new Random();
int current = 0;

do {
    current = random.Next(1, 11);
    Console.WriteLine(current);
} while (current != 7);
```

## while语句

`while` 语句首先计算布尔表达式，只要布尔表达式的计算结果为 `true`，就会继续循环访问代码块。

```c#
Random random = new Random();
int current = random.Next(1, 11);

while (current >= 3)
{
    Console.WriteLine(current);
    current = random.Next(1, 11);
}
Console.WriteLine($"Last number: {current}");
// 输出: 
// 9
// 7
// 5
// Last number: 1
```

