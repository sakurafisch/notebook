# cin 失败

当 cin 处于失败状态时，拒绝接下来的任何输入语句。因此我们需要清除这个输入失败状态，方法是调用 cin.clear() 这个函数。

```c++
cin.clear();
```

清除缓冲区中的错误输入

```c++
// 可能需要 #include <limits>
// 清除缓冲区到第一个换行符为止
cin.ignore(numeric_limits<streamsize>::max(), '\n');
```

用法示例

```c++
if (cin.fail()) {
    cin.clear();
    cin.ignore(numeric_limits<streamsize>::max(), '\n'); // 清除缓冲区一行字符
}
```

