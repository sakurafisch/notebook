# Java 笔记

## 修饰符和访问权限

访问级别修饰符确定其他类是否可以使用特定字段或调用特定方法。

### 类修饰符

- public：可以被任意 class 访问。一个程序的主类必须是公共类。

- abstract：抽象类。没有实现的方法，需要子类实现。

- final：不能被继承。

- friendly 或 没有修饰符：默认的修饰符，同一 package 的 class 可访问。

还可以用 private 或 protected 修饰 inner class。

非静态嵌套类（内部类）可以访问封闭类的其他成员，即使它们被声明为私有的也是如此。静态嵌套类无权访问封闭类的其他成员。

> Non-static nested classes (inner classes) have access to other members of the enclosing class, even if they are declared private. Static nested classes do not have access to other members of the enclosing class.

`InnerClass` 的实例只能存在于 `OuterClass` 的实例中，并且可以直接访问其封闭实例的方法和字段。要实例化内部类，必须首先实例化外部类。

> An instance of `InnerClass` can exist only within an instance of `OuterClass` and has direct access to the methods and fields of its enclosing instance. To instantiate an inner class, you must first instantiate the outer class. 

```java
// 实例化内部类的语法
OuterClass.InnerClass innerObject = outerObject.new InnerClass();
```

[Local Class](https://docs.oracle.com/javase/tutorial/java/javaOO/localclasses.html) 和 [Anonymous Class](https://docs.oracle.com/javase/tutorial/java/javaOO/anonymousclasses.html) 是两种典型的内部类。

#### Local Class 摘要
starting in Java SE 8, a local class can access local variables and parameters of the enclosing block that are final or *effectively final*. A variable or parameter whose value is never changed after it is initialized is effectively final.

从 Java SE 8 开始，如果在方法中声明了内部类，则它可以访问该方法的参数。

> Starting in Java SE 8, if you declare the local class in a method, it can access the method's parameters. 

静态方法中的内部类只能引用外部类的静态成员。

> Local classes in static methods can only refer to static members of the enclosing class.

局部类是非静态的，因为它们可以访问封闭块的实例成员。因此，它们不能包含大多数静态声明。

> Local classes are non-static because they have access to instance members of the enclosing block. Consequently, they cannot contain most kinds of static declarations.

您不能在块内声明接口；接口本质上是静态的。

> You cannot declare an interface inside a block; interfaces are inherently static.

您不能在本地类中声明静态初始化器或成员接口。

> You cannot declare static initializers or member interfaces in a local class.

本地类可以具有静态成员，前提是它们是常量变量。

> A local class can have static members provided that they are constant variables.

#### Anonymous Class 摘要

匿名类可以访问其外部类的成员。

> An anonymous class has access to the members of its enclosing class.

匿名类无法在其封闭范围内访问未声明为 final 或 effectively final 的局部变量。

> An anonymous class cannot access local variables in its enclosing scope that are not declared as `final` or effectively final.

您不能在匿名类中声明静态初始化器或成员接口。

> You cannot declare static initializers or member interfaces in an anonymous class.

匿名类可以具有静态成员，前提是它们是常量变量。

> An anonymous class can have static members provided that they are constant variables.

您不能在匿名类中声明构造函数。

> you cannot declare constructors in an anonymous class.

您可以在匿名类中声明以下内容：

- Fields    // 字段
- Extra methods (even if they do not implement any methods of the supertype)
- Instance initializers    // 实例初始化器
- Local classes    // 内部类

### 字段修饰符

- public：所有的 class 均可访问。
- protected：同一 package 的 class 可访问，且不同包内的子类可访问。
- friendly 或 没有修饰符：同一 package 的 class 可访问，且不同包内的子类不可访问。
- private：仅本类可访问。
- final：常量，不可重新赋值。
- static：静态变量。
- transient：为系统保留，暂无特别作用。
- volatile：可同时被几个线程控制和修改。

### 方法修饰符

- public：所有的 class 均可访问。
- protected：同一 package 的 class 可访问，且不同包内的子类可访问。
- friendly 或 没有修饰符：同一 package 的 class 可访问，且不同包内的子类不可访问。
- private：仅本类可访问。
- final：不可重载。
- static：静态方法。
- synchronize：在多个线程中，该修饰符用于在运行前，对方法加锁，以防止其他线程的访问，运行结束后解锁。
- native，本地修饰符。指定此方法的方法体是用其他语言在程序外部编写的。

## Lambda 表达式

[JDK 8 新特性](https://www.oracle.com/technetwork/java/javase/8-whats-new-2157071.html)

[Lambda 表达式](https://docs.oracle.com/javase/tutorial/java/javaOO/lambdaexpressions.html)

Lambda 表达式，也可称为闭包。

Lambda 允许把函数作为一个方法的参数。

```java
// Lambda 表达式语法
(parameters) -> expression
// 或者
(parameters) ->{ statements; }
```

lambda 表达式只能引用final 或 effectively final 的外层局部变量，这就是说不能在 lambda 内部修改定义在域外的局部变量，否则会编译错误。

## 方法引用

[方法引用](https://docs.oracle.com/javase/tutorial/java/javaOO/methodreferences.html)

### 引用静态方法

```java
// 常规的写法
Arrays.sort(rosterAsArray,
    (a, b) -> Person.compareByAge(a, b)
);
// 引用静态方法的写法
Arrays.sort(rosterAsArray, Person::compareByAge);
// 方法引用Person :: compareByAge在语义上与lambda表达式（a，b）-> Person.compareByAge（a，b）相同。
```

### 引用特定对象的实例方法

The following is an example of a reference to an instance method of an arbitrary object of a particular type:

```java
String[] stringArray = { "Barbara", "James", "Mary", "John",
    "Patricia", "Robert", "Michael", "Linda" };
Arrays.sort(stringArray, String::compareToIgnoreCase);
```

The equivalent lambda expression for the method reference `String::compareToIgnoreCase` would have the formal parameter list `(String a, String b)`, where `a` and `b` are arbitrary names used to better describe this example. The method reference would invoke the method `a.compareToIgnoreCase(b)`.

### 引用构造函数

You can reference a constructor in the same way as a static method by using the name `new`. The following method copies elements from one collection to another:

```java
public static <T, SOURCE extends Collection<T>, DEST extends Collection<T>>
    DEST transferElements(
        SOURCE sourceCollection,
        Supplier<DEST> collectionFactory) {
        
        DEST result = collectionFactory.get();
        for (T t : sourceCollection) {
            result.add(t);
        }
        return result;
}
```

The functional interface `Supplier` contains one method `get` that takes no arguments and returns an object. Consequently, you can invoke the method `transferElements` with a lambda expression as follows:

```java
Set<Person> rosterSetLambda =
    transferElements(roster, () -> { return new HashSet<>(); });
```

You can use a constructor reference in place of the lambda expression as follows:

```java
Set<Person> rosterSet = transferElements(roster, HashSet::new);
```

The Java compiler infers that you want to create a `HashSet` collection that contains elements of type `Person`. Alternatively, you can specify this as follows:

```java
Set<Person> rosterSet = transferElements(roster, HashSet<Person>::new);
```

## Package 摘要

[Package](https://docs.oracle.com/javase/tutorial/java/package/index.html)

程序包是一组相关类型的组合，提供访问保护和名称空间管理。

> A package is a grouping of related types providing access protection and name space management. 

Java 语言包本身以 `java` 或 `javax` 开头。

### 使用 package 的代码组织形式

```java
//in the Draggable.java file
package graphics;
public interface Draggable {
    . . .
}

//in the Graphic.java file
package graphics;
public abstract class Graphic {
    . . .
}

//in the Circle.java file
package graphics;
public class Circle extends Graphic
    implements Draggable {
    . . .
}

//in the Rectangle.java file
package graphics;
public class Rectangle extends Graphic
    implements Draggable {
    . . .
}

//in the Point.java file
package graphics;
public class Point extends Graphic
    implements Draggable {
    . . .
}

//in the Line.java file
package graphics;
public class Line extends Graphic
    implements Draggable {
    . . .
}
```

### 使用 package 的成员

为了方便起见，Java 编译器为每个源文件自动导入两个完整的包：

1. `Java.lang` 包

2. 当前包（当前文件的包）。

To use a `public` package member from outside its package, you must do one of the following:

- Refer to the member by its fully qualified name

```java
graphics.Rectangle myRect = new graphics.Rectangle();
```

- Import the package member

```java
import graphics.Rectangle;
Rectangle myRectangle = new Rectangle();
```

- Import the member's entire package

```java
import graphics.*;
Circle myCircle = new Circle();
Rectangle myRectangle = new Rectangle();
```

使用 `import` 语句，一般只能导入包中的一个成员或整个包。

但有两种特殊情况：

- 可以导入封闭类的公共嵌套类。

  例如，如果 `graphics.Rectangle` 类包含有用的嵌套类，例如 `Rectangle.DoubleWide` 和 `Rectangle.Square` ，则可以使用以下两条语句导入 Rectangle 及其嵌套类。

```java
import graphics.Rectangle;
import graphics.Rectangle.*;  // 这个 import 语句不会导入 Rectangle
```
- 使用 static import 语句导入要使用的常量和静态方法。

  例如， `JavaLang.Math` 类定义了 PI 常量和许多静态方法，包括计算正弦、余弦、切线、平方根、极大值、最小值、指数等的方法。

```java
// 在 javaLang.Math 里定义的常量和静态方法
public static final double PI 
    = 3.141592653589793;
public static double cos(double a)
{
    ...
}
```

```java
// static import
import static java.lang.Math.PI;  // 只导入常量 PI
import static java.lang.Math.*;  // 导入全部常量和静态方法
```

### 当域名特殊时，包名的约定

在某些情况下，internet 域名可能不是有效的包名称。如果域名包含连字符或其他特殊字符，如果包名称以非法用作 Java 名称开头的数字或其他字符开头，或者包名称包含保留的 Java 关键字（如 `int` ），则可能发生这种情况。在这种情况下，建议的约定是添加下划线。例如：

| Domain Name                   | Package Name Prefix           |
| ----------------------------- | ----------------------------- |
| `hyphenated-name.example.org` | `org.example.hyphenated_name` |
| `example.int`                 | `int_.example`                |
| `123name.example.com`         | `com.example._123name`        |

### 源文件组织方法

详见 [Managing Source and Class Files](https://docs.oracle.com/javase/tutorial/java/package/managingfiles.html)

## Exception 摘要

一个 exception 对象包含了错误信息，包括错误类型和出错时的程序状态。

> An *exception object*  contains information about the error, including its type and the state of the program when the error occurred. 

Creating an exception object and handing it to the runtime system is called *throwing an exception*.

The runtime system searches the call stack for a method that contains a block of code that can handle the exception. This block of code is called an *exception handler*.

An exception handler is considered appropriate if the type of the exception object thrown matches the type that can be handled by the handler.

### Exception 的种类

- *checked exception* 例如：  *java.io.FileNotFoundException*
- *error* 例如：*java.io.IOError*
- *runtime exception* 例如：*NullPointerException*

Errors and runtime exceptions are collectively known as *unchecked exceptions*.

> 译：错误和运行时异常统称为未检查异常。

代码示例

```java
// Note: This class will not compile yet.
import java.io.*;
import java.util.List;
import java.util.ArrayList;

public class ListOfNumbers {

    private List<Integer> list;
    private static final int SIZE = 10;

    public ListOfNumbers () {
        list = new ArrayList<Integer>(SIZE);
        for (int i = 0; i < SIZE; i++) {
            list.add(new Integer(i));
        }
    }

    public void writeList() {
	// The FileWriter constructor throws IOException, which must be caught.
        PrintWriter out = new PrintWriter(new FileWriter("OutFile.txt"));

        for (int i = 0; i < SIZE; i++) {
            // The get(int) method throws IndexOutOfBoundsException, which must be caught.
            out.println("Value at: " + i + " = " + list.get(i));
        }
        out.close();
    }
}
```

If you try to compile the [`ListOfNumbers`](https://docs.oracle.com/javase/tutorial/essential/exceptions/examples/ListOfNumbers.java) class, the compiler prints an error message about the exception thrown by the `FileWriter` constructor. However, it does not display an error message about the exception thrown by `get`. The reason is that the exception thrown by the constructor, `IOException`, is a checked exception, and the one thrown by the `get` method, `IndexOutOfBoundsException`, is an unchecked exception.

If a `catch` block handles more than one exception type, then the `catch` parameter is implicitly `final`. 

将清理代码放在 finally 块中始终是一个好的习惯，即使预期没有异常。

如果 JVM 在执行 try 或 catch 代码时退出，则 finally 块可能不执行。同样，如果执行 try 或 catch 代码的线程被中断或终止，那么即使整个应用程序继续运行，finally 块也可能不会执行。

## 泛型

详见

- [Java 泛型](https://docs.oracle.com/javase/tutorial/java/generics/index.html)

- [Java泛型-续](https://docs.oracle.com/javase/tutorial/extra/generics/index.html)

