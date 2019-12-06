# colorama

## 概述

colorama module 用于打印不同颜色的文本。

## 用法示例

```python
import colorama

colorama.init()
print(colorama.Fore.RED + 'This is red')
```

```python
from colorama import *

init()
print(Fore.BLUE + 'This is blue')
```

```python
from colorama import init, Fore

init()
print(Fore.GREEN + 'This is green')
```

