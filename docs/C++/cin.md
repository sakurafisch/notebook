# input

## cin

### 不读入换行符

`cin >>` 遇到换行符之后停止读取，但是不会读入这个换行符，在接下来的读取中还要考虑到这个换行符的存在。

1. 如果接下来还是用 `cin` 读取， `cin` 会自动跳过这个换行符继续下一行的读取，不会受影响。
2. 如果接下来用 `getline()` 读取，就需要先用 `getchar()` 吃掉这个换行符，以免 `getline()` 读入空数据。

### cin 失败

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


## getline

### 头文件

```c++
#include <string>
```

### 基本用法

```c++
istream &getline(char* buffer, streamsize num, char delim);
istream &getline(char* buffer, streamsize num);
```

- buffer： 进行读入操作的输入流
- num 存储读入的内容
- delim 终结符

```c++
string line;
getline(cin, line);
```

### 作为成员函数使用

```c++
cin.getline(char* cha, int num, char f);
```

向`cha`中输入`num`个字符，输入过程中达到`num-1`个数或者提前遇到`f`字符，输入结束。

当第三个参数省略时，系统默认为`'\0'`

```c++
char line[100];
cin.getline(line, 10, '?');
cout << line;
```
