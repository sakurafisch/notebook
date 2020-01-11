# Java 基本 IO 笔记

参考 [Basic I/O](https://docs.oracle.com/javase/tutorial/essential/io/index.html)

## I/O Streams

### Byte Streams

`Byte Streams` 只能用于最基本的 I/O。

CopyBytes.java 示范性地使用了 `FileInputStream` 和 `FileOutputStream`，该程序使用 byte streams 复制 xanadu.txt，一次复制一个字节。

```java
// CopyBytes.java
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

public class CopyBytes {
    public static void main(String[] args) throws IOException {

        FileInputStream fileInputStream = null;
        FileOutputStream fileOutputStream = null;

        try {
            fileInputStream = new FileInputStream("xanadu.txt");
            fileOutputStream = new FileOutputStream("outagain.txt");
            int copy;

            while ((copy = fileInputStream.read()) != -1) {
                fileOutputStream.write(copy);
            }
        } finally {
            if (fileInputStream != null) {
                fileInputStream.close();
            }
            if (fileOutputStream != null) {
                fileOutputStream.close();
            }
        }
    }
}
```

CopyBytes 将其大部分时间花费在一个简单的循环中，该循环读取输入流并写入输出流，一次写入一个字节。如下图所示：

![Simple byte stream input and output.](https://docs.oracle.com/javase/tutorial/figures/essential/byteStream.gif)

CopyBytes.java 看起来像一个普通的程序，但实际上它代表了一种低级别的 I/O，应该避免。由于 xanadu.txt 包含字符数据，最好的方法是使用 `character streams`。

### Character Streams

Java 平台使用 Unicode 约定存储字符值。`Character stream` I/O 自动将此内部格式转换为本地字符集。

使用 `stream classes` 完成的输入和输出自动转换为本地字符集。一种用 `character streams` 代替 `byte streams` 的程序，能自动适应本地字符集，并准备好进行国际化，而无需程序员额外的努力。

```java
// CopyCharacters.java
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;

public class CopyCharacters {
    public static void main(String[] args) throws IOException {

        FileReader inputStream = null;
        FileWriter outputStream = null;

        try {
            inputStream = new FileReader("xanadu.txt");
            outputStream = new FileWriter("characteroutput.txt");

            int copy;
            while ((copy = inputStream.read()) != -1) {
                outputStream.write(copy);
            }
        } finally {
            if (inputStream != null) {
                inputStream.close();
            }
            if (outputStream != null) {
                outputStream.close();
            }
        }
    }
}
```

`CopyCharacters.java` 与 `CopyBytes.java` 非常相似。最重要的区别是 `CopyCharacters.java` 使用 `FileReader` 和 `FileWriter` 进行输入和输出，而不是 `FileInputStream` 和 `FileOutputStream`。

CopyBytes.java 和 CopyCharacters.java 都使用一个 int 变量来读写。However, in CopyCharacters.java, the int variable holds a character value in its last 16 bits; in CopyBytes.java, the int variable holds a byte value in its last 8 bits.

### 面向行的 I / O

修改 CopyCharacters.java 示例以使用面向行的 I/O。为此，我们必须使用两个以前从未见过的类，[BufferedReader](https://docs.oracle.com/javase/8/docs/api/java/io/BufferedReader.html) 和 [PrintWriter](https://docs.oracle.com/javase/8/docs/api/java/io/PrintWriter.html)。

```java
// CopyLines.java
import java.io.FileReader;
import java.io.FileWriter;
import java.io.BufferedReader;
import java.io.PrintWriter;
import java.io.IOException;

public class CopyLines {
    public static void main(String[] args) throws IOException {

        BufferedReader inputStream = null;
        PrintWriter outputStream = null;

        try {
            inputStream = new BufferedReader(new FileReader("xanadu.txt"));
            outputStream = new PrintWriter(new FileWriter("characteroutput.txt"));

            String line;
            while ((line = inputStream.readLine()) != null) {
                outputStream.println(line);
            }
        } finally {
            if (inputStream != null) {
                inputStream.close();
            }
            if (outputStream != null) {
                outputStream.close();
            }
        }
    }
}
```

调用 `readLine` 返回一行文本。CopyLines.java 使用 `println` 输出每一行，`println` 附加当前操作系统的行结束符，这可能与输入文件中使用的行结束符不同。

还有许多种方法来构造文本输入和输出，参见 [Scanning and Formatting](https://docs.oracle.com/javase/tutorial/essential/io/scanfor.html) 。

### Buffered Streams

到目前为止，大多数示例都使用无缓冲 I/O，这意味着每个读或写请求都由底层操作系统直接处理。这会降低程序的效率，因为每一个这样的请求都会触发磁盘访问、网络活动等的操作。

为了减少这种开销，Java 平台实现了缓冲 I/O 流。缓冲输入流从称为缓冲区的内存区域读取数据；仅当缓冲区为空时才调用本机输入 API。类似地，缓冲输出流将数据写入缓冲区，并且仅当缓冲区已满时才调用本机输出 API。

> To reduce this kind of overhead, the Java platform implements buffered I/O streams. Buffered input streams read data from a memory area known as a buffer; the native input API is called only when the buffer is empty. Similarly, buffered output streams write data to a buffer, and the native output API is called only when the buffer is full.

无缓冲流对象被传递给缓冲流类的构造函数。修改 CopyCharacters.java 示例中的构造函数调用以使用缓冲 I/O：

```java
inputStream = new BufferedReader(new FileReader("xanadu.txt"));
outputStream = new BufferedWriter(new FileWriter("characteroutput.txt"));
```

有四个缓冲流类用于包装非缓冲流：[`BufferedInputStream`](https://docs.oracle.com/javase/8/docs/api/java/io/BufferedInputStream.html) 和 [`BufferedOutputStream`](https://docs.oracle.com/javase/8/docs/api/java/io/BufferedOutputStream.html) 创建缓冲字节流，而 [`BufferedReader`](https://docs.oracle.com/javase/8/docs/api/java/io/BufferedReader.html) 和 [`BufferedWriter`](https://docs.oracle.com/javase/8/docs/api/java/io/BufferedWriter.html) 创建缓冲字符流。

> `BufferedInputStream` and `BufferedOutputStream` create buffered byte streams, while `BufferedReader` and `BufferedWriter` create buffered character streams.

#### Flushing Buffered Streams

在关键点写出缓冲区通常是有意义的，而不必等待它被填满。这就是所谓的 flushing the buffer。

一些缓冲输出类支持 autoflush，由可选构造函数参数指定。启用自动刷新时，某些 key events 会导致刷新缓冲区。例如，每次调用 `println` 或 `format` 时，autoflush PrintWriter 对象都会刷新缓冲区。

有关这些方法的详细信息，请参见 [Formatting](https://docs.oracle.com/javase/tutorial/essential/io/formatting.html) 。

要手动地 flush a stream，请调用其 `flush` 方法。`flush` 方法对任何输出流都有效，但除非流被缓冲，否则它不会有任何效果。

### Scanning and Formatting

编程 I/O 通常涉及到转换成人们喜欢使用的格式整洁的数据。为了完成这些工作，Java 平台提供了两个 api。

1. [Scanner](https://docs.oracle.com/javase/8/docs/api/java/util/Scanner.html)：scanner API 将输入分解为与数据位相关联的单个 tokens。

> The scanner API breaks input into individual tokens associated with bits of data.

2. [Formatter](https://docs.oracle.com/javase/8/docs/api/java/util/Formatter.html)：formatting API 将数据组装成格式良好、可读的形式。

#### Scannig

Scanner 类型的对象可用于将格式化的输入分解为 tokens，并根据其数据类型转换为单个 tokens。

##### 将输入分为 Tokens

默认情况下，scanner 使用空白分隔标记。

（空白分隔标记包括空格、制表符和行结束符。参见 [Character.isWhitespace](https://docs.oracle.com/javase/8/docs/api/java/lang/Character.html#isWhitespace-char-) 。）

```java
// ScanXan.java
// 读取 xanadu.txt 中的单个单词并打印出来，每行一个。
import java.io.*;
import java.util.Scanner;

public class ScanXan {
    public static void main(String[] args) throws IOException {

        Scanner scanner = null;

        try {
            scanner = new Scanner(new BufferedReader(new FileReader("xanadu.txt")));

            while (scanner.hasNext()) {
                System.out.println(scanner.next());
            }
        } finally {
            if (scanner != null) {
                scanner.close();
            }
        }
    }
}
// 示例输出：
// In
// Xanadu
// did
// Kubla
// Khan
// A
// stately
// pleasure-dome
// ...
```

要使用其他标记分隔符，请调用 `useDelimiter()` ，并指定一个正则表达式。

```java
// 例如，假设您希望 tokens 分隔符是逗号，并且可以选择后面跟空格。您会调用:
scanner.useDelimiter(",\\s*");
```

##### Translating Individual Tokens

```java
// ScanSum.java
// 读取双精度值列表并将其加起来。
// 输入文件 usnumbers.txt 的内容如下：
// 8.5
// 32,767
// 3.14159
// 1,000,000.1
import java.io.FileReader;
import java.io.BufferedReader;
import java.io.IOException;
import java.util.Scanner;
import java.util.Locale;

public class ScanSum {
    public static void main(String[] args) throws IOException {

        Scanner scanner = null;
        double sum = 0;

        try {
            scanner = new Scanner(new BufferedReader(new FileReader("usnumbers.txt")));
            scanner.useLocale(Locale.US);

            while (scanner.hasNext()) {
                if (scanner.hasNextDouble()) {
                    sum += scanner.nextDouble();
                } else {
                    scanner.next();
                }   
            }
        } finally {
            scanner.close();
        }

        System.out.println(sum);
    }
}
```

#### Formatting

Stream objects that implement formatting are instances of either [PrintWriter](https://docs.oracle.com/javase/8/docs/api/java/io/PrintWriter.html), a character stream class, or [PrintStream](https://docs.oracle.com/javase/8/docs/api/java/io/PrintStream.html), a byte stream class.

您可能唯一需要的 `PrintStream` 对象是 `System.out` 和 `System.err` 。当需要创建格式化的输出流时，请实例化 `PrintWriter` 而不是 `PrintStream` 。

`PrintStream` 和 `PrintWriter` 的实例实现了一组用于 byte 和 character 输出的标准 write 方法。此外，`PrintStream` 和 `PrintWriter` 都实现了将内部数据转换为格式化输出的相同方法集。Two levels of formatting are provided：

- `print` 和 `println` 以标准方式格式化各个值。
- `format` 根据 `format string`  格式化数值。有许多选项用于设置精度。

##### The `print` and `println` Methods

```java
// Root.java
// 调用 print 或 println 在使用适当的 toString 方法转换值后输出单个值。
public class Root {
    public static void main(String[] args) {
        int i = 2;
        double r = Math.sqrt(i);
        
        System.out.print("The square root of ");
        System.out.print(i);
        System.out.print(" is ");
        System.out.print(r);
        System.out.println(".");

        i = 5;
        r = Math.sqrt(i);
        System.out.println("The square root of " + i + " is " + r + ".");
    }
}
// 输出如下：
// The square root of 2 is 1.4142135623730951.
// The square root of 5 is 2.23606797749979.
```

`i` 和 `r` 变量被格式化了两次：第一次的代码在被重载的 `print` 中；第二次的代码由 Java 编译器自动生成，它也使用 `toString` 。

> The `i` and `r` variables are formatted twice: the first time using code in an overload of `print`, the second time by conversion code automatically generated by the Java compiler, which also utilizes `toString`. You can format any value this way, but you don't have much control over the results.

##### The `format` Method

详见 [format string syntax](https://docs.oracle.com/javase/8/docs/api/java/util/Formatter.html#syntax)

`format` 方法基于 `format string` 格式化多个参数。`format string` 由嵌入了 `format specifiers` 的 `static text` 组成；除 `format specifiers` 外，格式字符串的输出保持不变。

```java
// Root2.java
public class Root2 {
    public static void main(String[] args) {
        int i = 2;
        double r = Math.sqrt(i);
        
        System.out.format("The square root of %d is %f.%n", i, r);  // 调用一次 format 方法格式化两个值
    }
}
// 输出如下：
// The square root of 2 is 1.414214.
```

除 `%%` 和 `％n` 外，所有格式说明符都必须与参数匹配。如果没有，则抛出异常。

在 Java 编程语言中，`\ n` 转义始终生成换行符（`\ u000A`）。除非特别需要换行符，否则不要使用 `\ n`。要为本地平台获取正确的行分隔符，请使用 `％n`。

除转换外，格式说明符还可以包含几个其他元素，以进一步自定义格式输出。

```java
// Format.java
// 这是一个示例 Format，它使用每种可能的元素。
public class Format {
    public static void main(String[] args) {
        System.out.format("%f, %1$+020.10f %n", Math.PI);
    }
}
// 输出如下：
// 3.141593, +00000003.1415926536
```

附加元素都是可选的。下图显示了较长的说明符如何分解为元素。

![Elements of a format specifier](https://docs.oracle.com/javase/tutorial/figures/essential/io-spec.gif)

元素必须按照显示的顺序出现。从右侧开始，可选元素为：

- Precision：对于浮点值，这是格式化值的数学精度。对于s和其他常规转换，这是格式化值的最大宽度。如果需要，该值将被右截断。
- Width：格式化值的最小宽度；如果需要，将填充该值。默认情况下，该值用空格左填充。
- Flags：指定其他格式选项。
- Argument Index：允许显式匹配指定的参数。

### I/O from the Command Line

Java 平台提供了两种在命令行环境下交互的方案：

1. Standard Streams
2. Console

#### Standard Streams

Standard Streams 从键盘读取输入并将输出写入显示器。它们还支持文件和程序之间的I / O，但是该功能由命令行解释器控制而不是程序。

Java平台支持三种 Standard Streams：

1. Standard Input，通过 System.in 访问。
2. Standard Output，通过 System.out 访问。
3. Standard Error，通过System.err访问。

这些对象是自动定义的，不需要打开。

`Standard Streams` 是 `byte streams` 而不是 `character streams`。`System.out` 和 `System.err` 被定义为 `PrintStream` 对象。尽管从技术上讲它是 `byte streams` ，但 `PrintStream` 还是利用内部 `character stream` 对象来模拟 `character streams` 的许多特性。

相比之下，`System.in` 是没有 `character stream` 功能的 `byte streams` 。要将标准输入用作 `character stream` ，请将 `System.in` 包装在 `InputStreamReader` 中。

```java
InputStreamReader cin = new InputStreamReader(System.in);
```

#### The Console

`Console` 对于安全密码输入特别有用。`Console` 对象还通过其 `reader` 和 `writer` 方法提供真正的 `character stream` 的输入和输出流。

在程序使用 `Console` 之前，它必须尝试通过调用 `System.console()` 来检索 `Console` 对象。如果 `Console` 对象可用，此方法将返回它。如果 `System.console` 返回 `NULL`，则不允许控制台操作，原因可能是操作系统不支持这些操作，或者是程序是在非交互环境中启动的。

`Console` 对象通过其 `readPassword` 方法支持安全密码输入。此方法通过两种方式帮助确保密码输入的安全。首先，它禁止 echoing，因此密码在用户屏幕上不可见。第二，`readPassword` 返回一个字符数组，而不是一个字符串，因此密码可以被 overwritten ，一旦不再需要它，就从内存中删除它。

```java
// Password.java
// 这是用于更改用户密码的原型程序。
// 它演示了几种 Console 方法。
import java.io.Console;
import java.util.Arrays;
import java.io.IOException;

public class Password {
    
    public static void main (String args[]) throws IOException {
        // 尝试检索控制台对象。如果对象不可用，则中止。
        Console console = System.console();
        if (console == null) {
            System.err.println("No console.");
            System.exit(1);
        }
        // 调用 Console.readLine 提示并读取用户的登录名。
        String login = console.readLine("Enter your login: ");
        char [] oldPassword = console.readPassword("Enter your old password: ");
        // 调用 verify 以确认用户有权更改密码。（在本例中，verify 是一个始终返回 true 的伪方法。）
        if (verify(login, oldPassword)) {
            boolean noMatch;
            // 重复以下步骤，直到用户输入同一密码两次。
            do {
                // 调用 Console.read password 两次以提示并读取新密码。
                char [] newPassword1 = console.readPassword("Enter your new password: ");
                char [] newPassword2 = console.readPassword("Enter new password again: ");
                noMatch = ! Arrays.equals(newPassword1, newPassword2);
                if (noMatch) {
                    console.format("Passwords don't match. Try again.%n");
                } else {
                    // 如果用户两次输入相同的密码，调用 change 来更改它。（同样，更改是一种虚拟方法。）
                    change(login, newPassword1);
                    console.format("Password for %s changed.%n", login);
                }
                // 用空白覆盖两个密码。
                Arrays.fill(newPassword1, ' ');
                Arrays.fill(newPassword2, ' ');
            } while (noMatch);
        }

        // 用空格覆盖旧密码。
        Arrays.fill(oldPassword, ' ');
    }
    
    // Dummy change method.
    static boolean verify(String login, char[] password) {
        // This method always returns
        // true in this example.
        // Modify this method to verify
        // password according to your rules.
        return true;
    }

    // Dummy change method.
    static void change(String login, char[] password) {
        // Modify this method to change
        // password according to your rules.
    }
}
```

### Data Streams

`Data streams` 支持原始数据类型和 String 类型的二进制 I/O。所有 `Data streams` 实现 `DataInput` 接口或 `DataOutput` 接口。

DataStreams.java 示例通过写出一组数据记录，然后再次读取它们来演示数据流。每条记录由与发票上的项目相关的三个值组成，如下表所示：

| Order in record | Data type | Data description | Output Method                  | Input Method                 | Sample Value     |
| --------------- | --------- | ---------------- | ------------------------------ | ---------------------------- | ---------------- |
| 1               | `double`  | Item price       | `DataOutputStream.writeDouble` | `DataInputStream.readDouble` | `19.99`          |
| 2               | `int`     | Unit count       | `DataOutputStream.writeInt`    | `DataInputStream.readInt`    | `12`             |
| 3               | `String`  | Item description | `DataOutputStream.writeUTF`    | `DataInputStream.readUTF`    | `"Java T-Shirt"` |

```java
// DataStreams.java


import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.BufferedInputStream;
import java.io.BufferedOutputStream;
import java.io.IOException;
import java.io.EOFException;

public class DataStreams {
    // 定义一些常量，其中包含数据文件的名称和将要写入其中的数据。
    static final String dataFile = "invoicedata";

    static final double[] prices = { 19.99, 9.99, 15.99, 3.99, 4.99 };
    static final int[] units = { 12, 8, 13, 29, 50 };
    static final String[] descs = { "Java T-shirt",
            "Java Mug",
            "Duke Juggling Dolls",
            "Java Pin",
            "Java Key Chain" };

    public static void main(String[] args) throws IOException {

 
        DataOutputStream out = null;
        
        try {
            // DataStreams 打开一个输出流。 
            // 由于只能将 DataOutputStream 创建为现有 byte stream 对象的包装，
            // 因此 DataStreams 提供了一个缓冲的文件输出 byte stream。
            out = new DataOutputStream(new
                    BufferedOutputStream(new FileOutputStream(dataFile)));

            // DataStreams 将写出记录并关闭输出流。
            for (int i = 0; i < prices.length; i ++) {
                out.writeDouble(prices[i]);
                out.writeInt(units[i]);
                out.writeUTF(descs[i]);  // writeUTF 方法以 UTF-8 的修改形式写出字符串值。 这是一种宽度可变的字符编码，对于常见的西方字符仅需要一个字节。
            }
        } finally {
            out.close();
        }

        DataInputStream in = null;
        double total = 0.0;
        try {
            // 现在，DataStreams 再次读回数据。
            // 首先，它必须提供 input stream 和用以保存输入数据变量。
            // 与 DataOutputStream 一样，DataInputStream 必须构造为 byte stream 的包装器。
            in = new DataInputStream(new
                    BufferedInputStream(new FileInputStream(dataFile)));

            double price;
            int unit;
            String desc;

            // 现在，DataStreams 可以读取流中的每条记录，报告其遇到的数据。
            // 注意，DataStream 通过捕获 EOFException 而不是测试无效的返回值来检测文件结束条件。
            // DataInput 方法的所有实现都使用 EOFException 而不是返回值。
            try {
                while (true) {
                    price = in.readDouble();
                    unit = in.readInt();
                    desc = in.readUTF();
                    System.out.format("You ordered %d units of %s at $%.2f%n",
                            unit, desc, price);
                    total += unit * price;
                }
            } catch (EOFException e) { }
            System.out.format("For a TOTAL of: $%.2f%n", total);
        }
        finally {
            in.close();
        }
    }
}
```

DataStreams 使用一种非常糟糕的编程技术：它使用浮点数来表示货币值。 通常，浮点数对精确的数值不利。 对于小数部分而言，这特别糟糕，因为通常的值（例如 0.1）没有二进制表示形式。

用于货币值的正确类型是 `java.math.BigDecimal`。不幸的是，`BigDecimal` 是一种对象类型，因此它不适用于数据流。但是，`BigDecimal` 可以与对象流一起使用，下文将介绍。

### Object Streams

正如 `data streams` 支持原始数据类型的 I / O 一样，`object streams` 支持对象的 I / O。

大多数（但不是全部）标准类支持其对象的序列化，是否支持对象的序列化取决于是否实现了标记接口 [`Serializable`](https://docs.oracle.com/javase/8/docs/api/java/io/Serializable.html) 。

object stream 类是 [`ObjectInputStream`](https://docs.oracle.com/javase/8/docs/api/java/io/ObjectInputStream.html) 和 [`ObjectOutputStream`](https://docs.oracle.com/javase/8/docs/api/java/io/ObjectOutputStream.html)。这些类实现了 [`ObjectInput`](https://docs.oracle.com/javase/8/docs/api/java/io/ObjectInput.html) 接口和 [`ObjectOutput`](https://docs.oracle.com/javase/8/docs/api/java/io/ObjectOutput.html) 接口，它们是 DataInput 和 DataOutput 的子接口。 这意味着 [Data Streams](https://docs.oracle.com/javase/tutorial/essential/io/datastreams.html) 中涵盖的所有原始数据类型 I / O 方法也都在 ObjectStreams 中实现。 因此，object stream 可以包含原始值和对象值的混合。 

ObjectStreams.java 创建与 DataStreams.java 相同的应用程序，但有两个变化。

- 首先，价格现在是  [`BigDecimal`](https://docs.oracle.com/javase/8/docs/api/java/math/BigDecimal.html) 对象，以更好地表示分数值。

- 其次，将 [`Calendar`](https://docs.oracle.com/javase/8/docs/api/java/util/Calendar.html) 对象写入数据文件，指示发票日期。

```java
// ObjectStreams.java
import java.io.*;
import java.math.BigDecimal;
import java.util.Calendar;

public class ObjectStreams {
    static final String dataFile = "invoicedata";

    static final BigDecimal[] prices = { 
        new BigDecimal("19.99"), 
        new BigDecimal("9.99"),
        new BigDecimal("15.99"),
        new BigDecimal("3.99"),
        new BigDecimal("4.99") };
    static final int[] units = { 12, 8, 13, 29, 50 };
    static final String[] descs = { "Java T-shirt",
            "Java Mug",
            "Duke Juggling Dolls",
            "Java Pin",
            "Java Key Chain" };

    // 如果 readObject() 没有返回预期的类型对象，则尝试将其强制转换为正确的类型可能会引发 ClassNotFoundException。
    //在这个简单的示例中，这不可能发生，因此我们不会尝试捕获异常。
    //取而代之的是，我们通过在主方法的 throws 子句中添加 ClassNotFoundException 来通知编译器我们已经知道该问题。
    public static void main(String[] args) 
        throws IOException, ClassNotFoundException {

 
        ObjectOutputStream out = null;
        try {
            out = new ObjectOutputStream(new
                    BufferedOutputStream(new FileOutputStream(dataFile)));

            out.writeObject(Calendar.getInstance());
            for (int i = 0; i < prices.length; i ++) {
                out.writeObject(prices[i]);
                out.writeInt(units[i]);
                out.writeUTF(descs[i]);
            }
        } finally {
            out.close();
        }

        ObjectInputStream in = null;
        try {
            in = new ObjectInputStream(new
                    BufferedInputStream(new FileInputStream(dataFile)));

            Calendar date = null;
            BigDecimal price;
            int unit;
            String desc;
            BigDecimal total = new BigDecimal(0);

            date = (Calendar) in.readObject();

            System.out.format ("On %tA, %<tB %<te, %<tY:%n", date);

            try {
                while (true) {
                    price = (BigDecimal) in.readObject();
                    unit = in.readInt();
                    desc = in.readUTF();
                    System.out.format("You ordered %d units of %s at $%.2f%n",
                            unit, desc, price);
                    total = total.add(price.multiply(new BigDecimal(unit)));
                }
            } catch (EOFException e) {}
            System.out.format("For a TOTAL of: $%.2f%n", total);
        } finally {
            in.close();
        }
    }
}
```

#### Output and Input of Complex Objects

`writeObject` 和 `readObject` 方法易于使用，但它们包含一些非常复杂的对象管理逻辑，但是许多对象包含对其他对象的引用。 如果 `readObject` 是流中的对象，则它必须能够还原原始对象所引用的所有对象。 这些其他对象可能有自己的引用，依此类推。 在这种情况下，`writeObject` 遍历对象引用的整个 Web 并将该 Web 中的所有对象写入流中。 因此，一次调用 `writeObject` 可以导致将大量对象写入流中。

在下图中演示了这一点，其中调用 `writeObject` 写入一个名为 `a` 的对象。 该对象包含对对象 `b` 和 `c` 的引用，而 `b` 包含对 `d` 和 `e` 的引用。 调用 `writeobject(a)` 不仅会写入 `a`，而且还会写入重建 `a` 所需的所有对象，因此也将写入此 Web 中的其他四个对象。 当 `readObject` 读回 `a` 时，其他四个对象也将被读回，并且所有原始对象引用都将保留。

![I/O of multiple referred-to objects](https://docs.oracle.com/javase/tutorial/figures/essential/io-trav.gif)

流只能包含对象的副本，尽管它可以包含对对象的任意数量的引用。 因此，如果您将一个对象显式地写入流两次，则实际上只写了两次引用。 例如，如果以下代码将对象 ob 两次写入流：

```java
Object ob = new Object();
out.writeObject(ob);
out.writeObject(ob);
```

每个 writeObject 必须由一个 readObject 进行匹配，因此读回流的代码将如下所示：

```java
Object ob1 = in.readObject();
Object ob2 = in.readObject();
```

这将导致两个变量 ob1 和 ob2 对单个对象的引用。

但是，如果将单个对象写入两个不同的流，则该对象将被有效地复制，读回两个流将看到两个不同的对象。

## File I/O

`java.nio.file` 软件包及其相关软件包 `java.nio.file.attribute` 为文件 I / O 和访问默认文件系统提供了全面的支持。

### The Path Class

Java SE 7 发行版中引入的 `Path` 类是 `java.nio.file` 包的主要入口点之一。 

Path 类是文件系统中路径的程序表示。 Path 对象包含用于构造路径的文件名和目录列表，并用于检查，定位和操作文件。

> A Path object contains the file name and directory list used to construct the path, and is used to examine, locate, and manipulate files.