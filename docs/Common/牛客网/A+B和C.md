# A+B和C

时间限制 1000 ms 内存限制 32768 KB 代码长度限制 100 KB

## 题目描述

```
给定区间[-2的31次方, 2的31次方]内的3个整数A、B和C，请判断A+B是否大于C。
```

## 输入描述:

```
输入第1行给出正整数T(<=10)，是测试用例的个数。随后给出T组测试用例，每组占一行，顺序给出A、B和C。整数间以空格分隔。
```

## 输出描述:

```
对每组测试用例，在一行中输出“Case #X: true”如果A+B>C，否则输出“Case #X: false”，其中X是测试用例的编号（从1开始）。
```

## 输入例子:

```
41 2 32 3 42147483647 0 21474836460 -2147483648 -2147483647
```

## 输出例子:

```
Case #1: falseCase #2: trueCase #3: trueCase #4: false
```



## 源码：

```cpp
#include<iostream>
using namespace std;
int main() {
    int n;
    cin >> n;
    n++;
    long long a, b, c;
    for(int i = 1; i < n; i++) {
        cin >> a >> b >> c;
        cout << "Case #" << i << ": " << (a + b > c ? "true" : "false") << endl;
    }
    return 0;
}
```

