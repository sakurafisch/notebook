# cout

默认输出浮点数时，采取保留 6 位有效数字的策略。

## 设置有效数字

```c++
cout.precision(12);
```

## 设置精度

有时候，设置有效数字个数的意义不大，此时我们还改为设置小数点后多少位：只需加上一句 cout << fixed; 即可。

```c++
cout << fixed; // 设置精度方式调整为小数点后 n 位
cout.precision(2); // 精度调到小数点后 2 位
```

## 设置科学记数法

使用科学记数法

```c++
cout << scientific;
```

不使用科学记数法

```c++
cout.setf(ios::fixed, ios::floatfield);
```

恢复默认格式（仅对大数使用科学记数法）

```c++
cout << defaultfloat;
```

