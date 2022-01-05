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

[Interface Path](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Path.html)

Java SE 7 发行版中引入的 `Path` 类是 `java.nio.file` 包的主要入口点之一。 

Path 类是文件系统中路径的程序表示。 Path 对象包含用于构造路径的文件名和目录列表，并用于检查，定位和操作文件。

> A Path object contains the file name and directory list used to construct the path, and is used to examine, locate, and manipulate files.

### Path 操作

> Path Operations

#### 创建一个 Path

> Creating a Path

Path实例包含用于指定文件或目录位置的信息。

使用Paths（请注意复数）帮助器类中的以下get方法之一轻松创建Path对象：

```java
Path p1 = Paths.get("/tmp/foo");
Path p2 = Paths.get(args[0]);
Path p3 = Paths.get(URI.create("file:///Users/joe/FileTest.java"));
```

Paths.get方法是以下代码的简写：

```java
Path p4 = FileSystems.getDefault().getPath("/users/sally");
```

以下示例假设您的主目录为 / u / joe 来创建 /u/joe/logs/foo.log ，如果在Windows上，则为C：\ joe \ logs \ foo.log。

#### 检索有关路径的信息

> Retrieving Information about a Path

```java
Path p5 = Paths.get(System.getProperty("user.home"),"logs", "foo.log");
```

您可以将 Path 视为将这些名称元素存储为序列。 目录结构中的最高元素将位于索引 0。目录结构中的最低元素将位于索引 [n-1]，其中 n 是路径中名称元素的数量。方法可用于使用这些索引检索单个元素或路径的子序列。

![Sample directory structure](https://docs.oracle.com/javase/tutorial/figures/essential/io-dirStructure.gif)

以下代码段定义了一个 Path 实例，然后调用几种方法来获取有关该路径的信息：

```java
// None of these methods requires that the file corresponding
// to the Path exists.
// Microsoft Windows syntax
Path path = Paths.get("C:\\home\\joe\\foo");

// Solaris syntax
Path path = Paths.get("/home/joe/foo");

System.out.format("toString: %s%n", path.toString());
System.out.format("getFileName: %s%n", path.getFileName());
System.out.format("getName(0): %s%n", path.getName(0));
System.out.format("getNameCount: %d%n", path.getNameCount());
System.out.format("subpath(0,2): %s%n", path.subpath(0,2));
System.out.format("getParent: %s%n", path.getParent());
System.out.format("getRoot: %s%n", path.getRoot());
```
这是 Windows 和 Solaris OS 的输出：

| Method Invoked | Returns in the Solaris OS | Returns in Microsoft Windows | Comment                                                      |
| -------------- | ------------------------- | ---------------------------- | ------------------------------------------------------------ |
| `toString`     | `/home/joe/foo`           | `C:\home\joe\foo`            | Returns the string representation of the `Path`. If the path was created using `Filesystems.getDefault().getPath(String)` or `Paths.get` (the latter is a convenience method for `getPath`), the method performs minor syntactic cleanup. For example, in a UNIX operating system, it will correct the input string `//home/joe/foo` to `/home/joe/foo`. |
| `getFileName`  | `foo`                     | `foo`                        | Returns the file name or the last element of the sequence of name elements. |
| `getName(0)`   | `home`                    | `home`                       | Returns the path element corresponding to the specified index. The 0th element is the path element closest to the root. |
| `getNameCount` | `3`                       | `3`                          | Returns the number of elements in the path.                  |
| `subpath(0,2)` | `home/joe`                | `home\joe`                   | Returns the subsequence of the `Path` (not including a root element) as specified by the beginning and ending indexes. |
| `getParent`    | `/home/joe`               | `\home\joe`                  | Returns the path of the parent directory.                    |
| `getRoot`      | `/`                       | `C:\`                        | Returns the root of the path.                                |

前面的示例显示了绝对路径的输出。在以下示例中，指定了相对路径：

```java
// Solaris syntax
Path path = Paths.get("sally/bar");
// or
// Microsoft Windows syntax
Path path = Paths.get("sally\\bar");
```

这是 Windows 和 Solaris OS 的输出：

| Method Invoked | Returns in the Solaris OS | Returns in Microsoft Windows |
| -------------- | ------------------------- | ---------------------------- |
| `toString`     | `sally/bar`               | `sally\bar`                  |
| `getFileName`  | `bar`                     | `bar`                        |
| `getName(0)`   | `sally`                   | `sally`                      |
| `getNameCount` | `2`                       | `2`                          |
| `subpath(0,1)` | `sally`                   | `sally`                      |
| `getParent`    | `sally`                   | `sally`                      |
| `getRoot`      | `null`                    | `null`                       |

#### 从路径中删除冗余

> Removing Redundancies From a Path

以下是包括冗余的示例：

```java
/home/./joe/foo
/home/sally/../joe/foo
```

规范化方法将删除所有冗余元素，其中包括任何  `.`  和  `..` 。前面的两个示例均规范化为 `/ home / joe / foo` 。

在第二个示例中，如果 `sally` 是符号链接（symbolic link），则删除 `sally/..` 可能会导致路径不再找到目标文件。

要在确保结果找到正确文件的同时清理路径，可以使用接下来介绍的 `toRealPath` 方法。

#### 转换路径

> Converting a Path

可以使用三种方法来转换路径。

1. `toUri` ：将路径转换为可以从浏览器打开的字符串。

2. `toAbsolutePath` ：将路径转换为绝对路径。

3.  `toRealPath`： 方法返回现有文件的真实路径。
    
    `toRealPath`  方法一次执行多项操作：
    
    - 如果将 `true` 传递给此方法，并且文件系统支持符号链接，则此方法将解析路径中的所有符号链接。
    - 返回绝对路径。
    - 删除冗余元素。
    
    如果文件不存在或无法访问，则 `toRealPath` 方法将引发异常。

```java
// toUri 用法示例
Path p1 = Paths.get("/home/logfile");
// Result is file:///home/logfile
System.out.format("%s%n", p1.toUri());
```
```java
// toAbsolutePath 用法示例
// 文件名为 FileTest.java
public class FileTest {
    public static void main(String[] args) {

        if (args.length < 1) {
            System.out.println("usage: FileTest file");
            System.exit(-1);
        }

        Path inputPath = Paths.get(args[0]);  // 把输入字符串转换为 Path 对象

        Path fullPath = inputPath.toAbsolutePath();  // 把 Path 对象的路径转换为绝对路径
    }
}
```

```java
// toRealPath 用法示例
try {
    Path fp = path.toRealPath();
} catch (NoSuchFileException x) {
    System.err.format("%s: no such" + " file or directory%n", path);
    // Logic for case when file doesn't exist.
} catch (IOException x) {
    System.err.format("%s%n", x);
    // Logic for other sort of file error.
}
```

#### 连接两个 Path

> Joining Two Paths

使用 `resolve` 方法合并路径。

```java
// Solaris
Path p1 = Paths.get("/home/joe/foo");
// Result is /home/joe/foo/bar
System.out.format("%s%n", p1.resolve("bar"));

// or

// Microsoft Windows
Path p1 = Paths.get("C:\\home\\joe\\foo");
// Result is C:\home\joe\foo\bar
System.out.format("%s%n", p1.resolve("bar"));
```

将绝对路径传递给 `resolve` 方法将返回传入的路径：

```java
// Result is /home/joe
Paths.get("foo").resolve("/home/joe");
```

#### 在两条路径之间创建一条路径

此方法构造一条路径，该路径从原始路径开始，并在传入路径的指定位置处终止。 新路径是相对于原始路径的。

```java
Path p1 = Paths.get("home");
Path p3 = Paths.get("home/sally/bar");

Path p1_to_p3 = p1.relativize(p3);  // Result is sally/bar

Path p3_to_p1 = p3.relativize(p1);  // Result is ../..
```

如果只有一个路径包含 root 元素，则无法构造相对路径。 如果两个路径都包含 root 元素，则构造相对路径的能力取决于系统。

递归 [`Copy`](https://docs.oracle.com/javase/tutorial/essential/io/examples/Copy.java) 示例使用相对化和求解方法：

```java
// Copy.java
import java.nio.file.*;
import static java.nio.file.StandardCopyOption.*;
import java.nio.file.attribute.*;
import static java.nio.file.FileVisitResult.*;
import java.io.IOException;
import java.util.*;

/**
 * Sample code that copies files in a similar manner to the cp(1) program.
 */

public class Copy {

    /**
     * Returns {@code true} if okay to overwrite a  file ("cp -i")
     */
    static boolean okayToOverwrite(Path file) {
        String answer = System.console().readLine("overwrite %s (yes/no)? ", file);
        return (answer.equalsIgnoreCase("y") || answer.equalsIgnoreCase("yes"));
    }

    /**
     * Copy source file to target location. If {@code prompt} is true then
     * prompt user to overwrite target if it exists. The {@code preserve}
     * parameter determines if file attributes should be copied/preserved.
     */
    static void copyFile(Path source, Path target, boolean prompt, boolean preserve) {
        CopyOption[] options = (preserve) ?
            new CopyOption[] { COPY_ATTRIBUTES, REPLACE_EXISTING } :
            new CopyOption[] { REPLACE_EXISTING };
        if (!prompt || Files.notExists(target) || okayToOverwrite(target)) {
            try {
                Files.copy(source, target, options);
            } catch (IOException x) {
                System.err.format("Unable to copy: %s: %s%n", source, x);
            }
        }
    }

    /**
     * A {@code FileVisitor} that copies a file-tree ("cp -r")
     */
    static class TreeCopier implements FileVisitor<Path> {
        private final Path source;
        private final Path target;
        private final boolean prompt;
        private final boolean preserve;

        TreeCopier(Path source, Path target, boolean prompt, boolean preserve) {
            this.source = source;
            this.target = target;
            this.prompt = prompt;
            this.preserve = preserve;
        }

        @Override
        public FileVisitResult preVisitDirectory(Path dir, BasicFileAttributes attrs) {
            // before visiting entries in a directory we copy the directory
            // (okay if directory already exists).
            CopyOption[] options = (preserve) ?
                new CopyOption[] { COPY_ATTRIBUTES } : new CopyOption[0];

            Path newdir = target.resolve(source.relativize(dir));
            try {
                Files.copy(dir, newdir, options);
            } catch (FileAlreadyExistsException x) {
                // ignore
            } catch (IOException x) {
                System.err.format("Unable to create: %s: %s%n", newdir, x);
                return SKIP_SUBTREE;
            }
            return CONTINUE;
        }

        @Override
        public FileVisitResult visitFile(Path file, BasicFileAttributes attrs) {
            copyFile(file, target.resolve(source.relativize(file)),
                     prompt, preserve);
            return CONTINUE;
        }

        @Override
        public FileVisitResult postVisitDirectory(Path dir, IOException exc) {
            // fix up modification time of directory when done
            if (exc == null && preserve) {
                Path newdir = target.resolve(source.relativize(dir));
                try {
                    FileTime time = Files.getLastModifiedTime(dir);
                    Files.setLastModifiedTime(newdir, time);
                } catch (IOException x) {
                    System.err.format("Unable to copy all attributes to: %s: %s%n", newdir, x);
                }
            }
            return CONTINUE;
        }

        @Override
        public FileVisitResult visitFileFailed(Path file, IOException exc) {
            if (exc instanceof FileSystemLoopException) {
                System.err.println("cycle detected: " + file);
            } else {
                System.err.format("Unable to copy: %s: %s%n", file, exc);
            }
            return CONTINUE;
        }
    }

    static void usage() {
        System.err.println("java Copy [-ip] source... target");
        System.err.println("java Copy -r [-ip] source-dir... target");
        System.exit(-1);
    }

    public static void main(String[] args) throws IOException {
        boolean recursive = false;
        boolean prompt = false;
        boolean preserve = false;

        // process options
        int argi = 0;
        while (argi < args.length) {
            String arg = args[argi];
            if (!arg.startsWith("-"))
                break;
            if (arg.length() < 2)
                usage();
            for (int i=1; i<arg.length(); i++) {
                char c = arg.charAt(i);
                switch (c) {
                    case 'r' : recursive = true; break;
                    case 'i' : prompt = true; break;
                    case 'p' : preserve = true; break;
                    default : usage();
                }
            }
            argi++;
        }

        // remaining arguments are the source files(s) and the target location
        int remaining = args.length - argi;
        if (remaining < 2)
            usage();
        Path[] source = new Path[remaining-1];
        int i=0;
        while (remaining > 1) {
            source[i++] = Paths.get(args[argi++]);
            remaining--;
        }
        Path target = Paths.get(args[argi]);

        // check if target is a directory
        boolean isDir = Files.isDirectory(target);

        // copy each source file/directory to target
        for (i=0; i<source.length; i++) {
            Path dest = (isDir) ? target.resolve(source[i].getFileName()) : target;

            if (recursive) {
                // follow links when copying files
                EnumSet<FileVisitOption> opts = EnumSet.of(FileVisitOption.FOLLOW_LINKS);
                TreeCopier tc = new TreeCopier(source[i], dest, prompt, preserve);
                Files.walkFileTree(source[i], opts, Integer.MAX_VALUE, tc);
            } else {
                // not recursive so source must not be a directory
                if (Files.isDirectory(source[i])) {
                    System.err.format("%s: is a directory%n", source[i]);
                    continue;
                }
                copyFile(source[i], dest, prompt, preserve);
            }
        }
    }
}
```

#### 比较两条路径


> Comparing Two Paths

`Path` 类支持 [`equals`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Path.html#equals-java.lang.Object-)，能够测试两条路径是否相等； [`startsWith`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Path.html#startsWith-java.nio.file.Path-) 和 [`endsWith`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Path.html#endsWith-java.nio.file.Path-) 方法能够测试路径是以特定字符串开头还是结尾。

```java
Path path = ...;
Path otherPath = ...;
Path beginning = Paths.get("/home");
Path ending = Paths.get("foo");

if (path.equals(otherPath)) {
    // equality logic here
} else if (path.startsWith(beginning)) {
    // path begins with "/home"
} else if (path.endsWith(ending)) {
    // path ends with "foo"
}
```

`Path` 类实现了 [`Iterable`](https://docs.oracle.com/javase/8/docs/api/java/lang/Iterable.html) 接口。迭代器（iterator）方法返回一个对象，可以迭代路径中的名称元素。返回的第一个元素是最接近目录树根目录的元素。

```java
//  以下代码段遍历路径，打印每个 name 元素：
Path path = ...;
for (Path name: path) {
    System.out.println(name);
}
```

`Path` 类还实现 [`Comparable`](https://docs.oracle.com/javase/8/docs/api/java/lang/Comparable.html) 接口。 可以使用 `compareTo` 来比较 `Path` 对象，这对于排序很有用。

也可以将 Path 对象放入集合中。详见 [Collections](https://docs.oracle.com/javase/tutorial/collections/index.html) 。

要验证两个 `Path` 对象是否位于同一文件时，可以使用 `isSameFile` 方法。详见：[Checking Whether Two Paths Locate the Same File](https://docs.oracle.com/javase/tutorial/essential/io/check.html#same) 。

### File Operations

可能有用的文档：

- [FileSystemException](https://docs.oracle.com/javase/8/docs/api/java/nio/file/FileSystemException.html)

`File` 类是 `java.nio.file` 包的另一个主要入口点。 此类提供了一组丰富的静态方法，用于读取，写入和操作文件和目录。 `File` 方法对 `Path` 对象的实例起作用。 

#### Link Awareness

`File` 类是 “link aware” 的。每个 `File` 方法要么检测遇到符号链接时的操作，要么提供一个用于配置遇到符号链接时的行为的选项。

> The `Files` class is "link aware." Every `Files` method either detects what to do when a symbolic link is encountered, or it provides an option enabling you to configure the behavior when a symbolic link is encountered.

### 检查文件或目录

> Checking a File or Directory

#### 检查文件或目录是否存在

> Verifying the Existence of a File or Directory

使用 [`exists(Path, LinkOption...)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#exists-java.nio.file.Path-java.nio.file.LinkOption...-) 或  [`notExists(Path, LinkOption...)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#notExists-java.nio.file.Path-java.nio.file.LinkOption...-) 方法检查文件或目录是否存在。

`!Files.exists(path)` 不等同于 `Files.notExists(path)`

当验证文件是否存在时，可能会出现三个结果：

1. 文件存在。
2. 文件不存在。
3. 文件的存在性未知。（当程序无权访问该文件时，可能会出现此结果。）

如果 `exists` 和 `notExists` 都返回 `false`, 那么文件的存在性将无法验证。

#### 检查文件的可访问性

> Checking File Accessibility

使用  [`isReadable(Path)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#isReadable-java.nio.file.Path-) ， [`isWritable(Path)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#isWritable-java.nio.file.Path-) ，或 [`isExecutable(Path)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#isExecutable-java.nio.file.Path-) 方法检查文件的可访问性。

下面的代码片段验证了特定文件的存在以及该程序具有执行该文件的能力。

```java
Path file = ...;
boolean isRegularExecutableFile = Files.isRegularFile(file) &
         Files.isReadable(file) & Files.isExecutable(file);
```

#### 检查两个 Path 对象是否指向同一个文件

> Checking Whether Two Paths Locate the Same File

`isSameFile(Path，Path)`方法比较两个路径以确定它们是否在文件系统上指向相同的文件。

```java
Path p1 = ...;
Path p2 = ...;

if (Files.isSameFile(p1, p2)) {
    // Logic when the paths locate the same file
}
```

### 删除文件或目录

> Deleting a File or Directory

可以删除文件，目录或链接。

使用符号链接时，链接将被删除，而不是链接的目标。对于目录，目录必须为空，否则删除失败。

`File` 类提供了两种删除方法。

- [`delete(Path)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#delete-java.nio.file.Path-) 方法用于删除文件，如果删除失败将引发异常。例如，如果文件不存在，则会引发NoSuchFileException。

```java
// delete(Path) 用法示例
try {
    Files.delete(path);
} catch (NoSuchFileException x) {
    System.err.format("%s: no such" + " file or directory%n", path);
} catch (DirectoryNotEmptyException x) {
    System.err.format("%s not empty%n", path);
} catch (IOException x) {
    // File permission problems are caught here.
    System.err.println(x);
}
```

- [`deleteIfExists(Path)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#deleteIfExists-java.nio.file.Path-) 方法也会删除文件，但不会引发任何异常。 当有多个线程删除文件并且不想仅因为一个线程先这样做而引发异常时，静默失败很有用。

### 复制文件或目录

> Copying a File or Directory

可以使用 [`copy(Path, Path, CopyOption...)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#copy-java.nio.file.Path-java.nio.file.Path-java.nio.file.CopyOption...-) 方法复制文件或目录。如果目标文件存在，则复制失败，除非指定了 `REPLACE_EXISTING` 选项。

目录可以被复制。但是，目录内的文件不会被复制，因此即使原始目录包含文件，新目录也为空。

复制符号链接时，将会复制所链接的目标。如果要复制链接本身而不是复制链接的内容，请指定 `NOFOLLOW_LINKS` 或 `REPLACE_EXISTING` 选项。

此方法采用 varargs 参数。支持以下 `StandardCopyOption` 和 `LinkOption` 枚举：

- `REPLACE_EXISTING` 即使目标文件已经存在，也执行复制。
  -  如果目标是符号链接，则复制链接本身（而不是链接的目标）。
  -  如果目标是非空目录，则复制失败并抛出FileAlreadyExistsException 异常。
- `COPY_ATTRIBUTES` 将与文件关联的属性复制到目标文件。
  - 所支持的属性取决于文件系统和平台
  - 跨平台支持上次修改时间。

- `NOFOLLOW_LINKS` 不遵循符号链接。
  - 如果要复制的文件是符号链接，则复制该链接（而不是链接的目标）。

```java
// 复制文件示例
import static java.nio.file.StandardCopyOption.*;
...
Files.copy(source, target, REPLACE_EXISTING);
```

[`copy(InputStream, Path, CopyOptions...)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#copy-java.io.InputStream-java.nio.file.Path-java.nio.file.CopyOption...-) 方法可用于将所有字节从输入流复制到文件。

 [`copy(Path, OutputStream)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#copy-java.nio.file.Path-java.io.OutputStream-) 方法可用于将所有字节从文件复制到输出流。

```java
// Copy.java
// 使用 copy 和 Files.walkFileTree 方法来支持递归复制
import java.nio.file.*;
import static java.nio.file.StandardCopyOption.*;
import java.nio.file.attribute.*;
import static java.nio.file.FileVisitResult.*;
import java.io.IOException;
import java.util.*;

/**
 * Sample code that copies files in a similar manner to the cp(1) program.
 */

public class Copy {

    /**
     * Returns {@code true} if okay to overwrite a  file ("cp -i")
     */
    static boolean okayToOverwrite(Path file) {
        String answer = System.console().readLine("overwrite %s (yes/no)? ", file);
        return (answer.equalsIgnoreCase("y") || answer.equalsIgnoreCase("yes"));
    }

    /**
     * Copy source file to target location. If {@code prompt} is true then
     * prompt user to overwrite target if it exists. The {@code preserve}
     * parameter determines if file attributes should be copied/preserved.
     */
    static void copyFile(Path source, Path target, boolean prompt, boolean preserve) {
        CopyOption[] options = (preserve) ?
            new CopyOption[] { COPY_ATTRIBUTES, REPLACE_EXISTING } :
            new CopyOption[] { REPLACE_EXISTING };
        if (!prompt || Files.notExists(target) || okayToOverwrite(target)) {
            try {
                Files.copy(source, target, options);
            } catch (IOException x) {
                System.err.format("Unable to copy: %s: %s%n", source, x);
            }
        }
    }

    /**
     * A {@code FileVisitor} that copies a file-tree ("cp -r")
     */
    static class TreeCopier implements FileVisitor<Path> {
        private final Path source;
        private final Path target;
        private final boolean prompt;
        private final boolean preserve;

        TreeCopier(Path source, Path target, boolean prompt, boolean preserve) {
            this.source = source;
            this.target = target;
            this.prompt = prompt;
            this.preserve = preserve;
        }

        @Override
        public FileVisitResult preVisitDirectory(Path dir, BasicFileAttributes attrs) {
            // before visiting entries in a directory we copy the directory
            // (okay if directory already exists).
            CopyOption[] options = (preserve) ?
                new CopyOption[] { COPY_ATTRIBUTES } : new CopyOption[0];

            Path newdir = target.resolve(source.relativize(dir));
            try {
                Files.copy(dir, newdir, options);
            } catch (FileAlreadyExistsException x) {
                // ignore
            } catch (IOException x) {
                System.err.format("Unable to create: %s: %s%n", newdir, x);
                return SKIP_SUBTREE;
            }
            return CONTINUE;
        }

        @Override
        public FileVisitResult visitFile(Path file, BasicFileAttributes attrs) {
            copyFile(file, target.resolve(source.relativize(file)),
                     prompt, preserve);
            return CONTINUE;
        }

        @Override
        public FileVisitResult postVisitDirectory(Path dir, IOException exc) {
            // fix up modification time of directory when done
            if (exc == null && preserve) {
                Path newdir = target.resolve(source.relativize(dir));
                try {
                    FileTime time = Files.getLastModifiedTime(dir);
                    Files.setLastModifiedTime(newdir, time);
                } catch (IOException x) {
                    System.err.format("Unable to copy all attributes to: %s: %s%n", newdir, x);
                }
            }
            return CONTINUE;
        }

        @Override
        public FileVisitResult visitFileFailed(Path file, IOException exc) {
            if (exc instanceof FileSystemLoopException) {
                System.err.println("cycle detected: " + file);
            } else {
                System.err.format("Unable to copy: %s: %s%n", file, exc);
            }
            return CONTINUE;
        }
    }

    static void usage() {
        System.err.println("java Copy [-ip] source... target");
        System.err.println("java Copy -r [-ip] source-dir... target");
        System.exit(-1);
    }

    public static void main(String[] args) throws IOException {
        boolean recursive = false;
        boolean prompt = false;
        boolean preserve = false;

        // process options
        int argi = 0;
        while (argi < args.length) {
            String arg = args[argi];
            if (!arg.startsWith("-"))
                break;
            if (arg.length() < 2)
                usage();
            for (int i=1; i<arg.length(); i++) {
                char c = arg.charAt(i);
                switch (c) {
                    case 'r' : recursive = true; break;
                    case 'i' : prompt = true; break;
                    case 'p' : preserve = true; break;
                    default : usage();
                }
            }
            argi++;
        }

        // remaining arguments are the source files(s) and the target location
        int remaining = args.length - argi;
        if (remaining < 2)
            usage();
        Path[] source = new Path[remaining-1];
        int i=0;
        while (remaining > 1) {
            source[i++] = Paths.get(args[argi++]);
            remaining--;
        }
        Path target = Paths.get(args[argi]);

        // check if target is a directory
        boolean isDir = Files.isDirectory(target);

        // copy each source file/directory to target
        for (i=0; i<source.length; i++) {
            Path dest = (isDir) ? target.resolve(source[i].getFileName()) : target;

            if (recursive) {
                // follow links when copying files
                EnumSet<FileVisitOption> opts = EnumSet.of(FileVisitOption.FOLLOW_LINKS);
                TreeCopier tc = new TreeCopier(source[i], dest, prompt, preserve);
                Files.walkFileTree(source[i], opts, Integer.MAX_VALUE, tc);
            } else {
                // not recursive so source must not be a directory
                if (Files.isDirectory(source[i])) {
                    System.err.format("%s: is a directory%n", source[i]);
                    continue;
                }
                copyFile(source[i], dest, prompt, preserve);
            }
        }
    }
}
```

### 移动文件或目录

> Moving a File or Directory

可以使用 [`move(Path, Path, CopyOption...)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#move-java.nio.file.Path-java.nio.file.Path-java.nio.file.CopyOption...-)方法移动文件或目录。如果目标文件存在，则移动将失败，除非指定了 `REPLACE_EXISTING` 选项。

空目录可以移动。 如果目录不为空，则可以在不移动目录内容的情况下进行移动。 在 UNIX 系统上，在同一分区内移动目录通常包括重命名目录； 在这种情况下，即使目录包含文件，此方法也可以使用。

此方法采用 varargs 参数，支持以下 `StandardCopyOption` 枚举：

- `REPLACE_EXISTING` 即使目标文件已经存在也执行移动。 
  - 如果目标是符号链接，则将替换符号链接，但是它所指向的内容不受影响。
- `ATOMIC_MOVE` 将移动作为原子文件操作执行。
  -  如果文件系统不支持原子移动，则会引发异常。
  - 可以将文件移动到目录中，并确保监视目录的所有进程都可以访问完整的文件。

```java
// 移动文件示例
import static java.nio.file.StandardCopyOption.*;
...
Files.move(source, target, REPLACE_EXISTING);
```

### 管理元数据

> Managing Metadata 

元数据的定义是 “关于其他数据的数据”。文件系统的元数据通常称为其文件属性。 File 类包含可用于获取文件的单个属性或设置属性的方法。

| Methods                                                      | Comment                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [`size(Path)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#size-java.nio.file.Path-) | Returns the size of the specified file in bytes.             |
| [`isDirectory(Path, LinkOption)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#isDirectory-java.nio.file.Path-java.nio.file.LinkOption...-) | Returns true if the specified `Path` locates a file that is a directory. |
| [`isRegularFile(Path, LinkOption...)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#isRegularFile-java.nio.file.Path-java.nio.file.LinkOption...-) | Returns true if the specified `Path` locates a file that is a regular file. |
| [`isSymbolicLink(Path)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#isSymbolicLink-java.nio.file.Path-) | Returns true if the specified `Path` locates a file that is a symbolic link. |
| [`isHidden(Path)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#isHidden-java.nio.file.Path-) | Returns true if the specified `Path` locates a file that is considered hidden by the file system. |
| [`getLastModifiedTime(Path, LinkOption...)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#getLastModifiedTime-java.nio.file.Path-java.nio.file.LinkOption...-) [`setLastModifiedTime(Path, FileTime)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#setLastModifiedTime-java.nio.file.Path-java.nio.file.attribute.FileTime-) | Returns or sets the specified file's last modified time.     |
| [`getOwner(Path, LinkOption...)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#getOwner-java.nio.file.Path-java.nio.file.LinkOption...-) [`setOwner(Path, UserPrincipal)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#setOwner-java.nio.file.Path-java.nio.file.attribute.UserPrincipal-) | Returns or sets the owner of the file.                       |
| [`getPosixFilePermissions(Path, LinkOption...)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#getPosixFilePermissions-java.nio.file.Path-java.nio.file.LinkOption...-) [`setPosixFilePermissions(Path, Set)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#setPosixFilePermissions-java.nio.file.Path-java.util.Set-) | Returns or sets a file's POSIX file permissions.             |
| [`getAttribute(Path, String, LinkOption...)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#getAttribute-java.nio.file.Path-java.lang.String-java.nio.file.LinkOption...-) [`setAttribute(Path, String, Object, LinkOption...)`](https://docs.oracle.com/javase/8/docs/api/java/nio/file/Files.html#setAttribute-java.nio.file.Path-java.lang.String-java.lang.Object-java.nio.file.LinkOption...-) | Returns or sets the value of a file attribute.               |

