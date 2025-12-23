# Python笔记

## 运行 Python 代码

Python 是一种解释语言，它缩短了编辑-测试-调试周期，因为不需要编译步骤。 为了运行 Python 应用，你需要运行时环境/解释器来执行代码。  

大多数运行时环境都支持两种方式执行 Python 代码：  

1. 交互模式  
2. 脚本模式

## Python 实现

Python 是在 OSI 开放源代码许可证下获得许可的，并根据需求提供多种实现方式。

- 最受欢迎的是引用实现 (CPython)，可从 [Python 网站](https://www.python.org/)中获取。 CPython 通常用于 Web 开发、应用程序开发和脚本编辑。 
- [Anaconda](https://www.anaconda.com/) 是专为科学编程任务量身定制的专业 Python 发行版。
- [IronPython](https://ironpython.net/) 是基于 .NET 运行时构建的 Python 的开放源代码实现。
- [Jupyter 笔记本](https://jupyter.org/)是基于 Web 的交互式编程环境，支持包括 Python 在内的各种编程语言。 

**Tips: WIN10用户可以在WIN10商店一键安装Python install manager。**

## 变量和基本数据类型

参考 [Python 中的变量和基本数据类型](https://docs.microsoft.com/zh-cn/learn/modules/intro-to-python/4-variables-and-data-types)

## 注释

```python
# I am a comment...
```

对于多行注释，需将井号放在每行开头。

## 读取键盘输入

`input` 读取用户在键盘上键入的内容并将其作为字符串返回。 作为参数传递到 `input` 函数的字符串是用户将看到的提示。 

```python
name = input('Enter your name:')
print(name)
```

也可以在不使用参数的情况下调用 `input` 函数：

```python
print('What is your name?')
name = input()
print(name)
```

此程序的行为几乎与第一个程序的行为相同。 不同之处在于 `print`（默认）将新行添加到输出。

### 读取数字作为输入

`input` 函数始终返回字符串（文本）。若要将值转换为真正的整数变量，可以使用 `int` 函数：

```python
x = int(input('Enter a number: '))
```

### 将数字转换为字符串

`str` 运算符将获取整数或浮点值并将其转换为字符串。 如果要将数字连接到字符串，这是必需的操作，例如：

```python
x = 5
print('The number is ' + str(x))
```

## a+b示例

```python
# Read the first number
x = int(input('Enter the first number: '))

# Get the second number
y = int(input('Enter the second number: '))

# Add them together
result = x + y

# Display the results
print(str(x) + " + " + str(y) + " = " + str(result))
```

## lambdas

```python
# lambda 写法
lambda item: item['name']
```

```python
# 常规函数的写法
def sorter(item):
    return item['name']
```

## 进一步学习

- [Python 编程风格](https://docs.python-guide.org/writing/style/)
- [Python 中可用的字符串函数](https://www.w3schools.com/python/python_ref_string.asp)
- [内置 Python 函数列表](https://docs.python.org/3/library/functions.html)
- [官方文档](https://docs.python.org/zh-cn/3/tutorial/)

