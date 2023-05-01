# Virtual environments

> Virtual environment is nothing but a folder.

## 安装

```cmd
pip install virtualenv  # Install virtual environment
```

## 创建

> Creating a virtual environment.

```cmd
python -m venv <folder_name>  # Windows systems
```
```bash
virtualenv <folder_name>  # OSX/Linux (bash)
```

## 使用

> Using virtual environments.

我们首先要 ```激活``` 它。（Firstly we need to ```activate``` it.）

激活脚本 ```Activate.*``` 位于 ```<folder_name>/Scripts``` 文件夹。

```cmd
# Windows systems
<folder_name>\Scripts\Activate.bat  # cmd.exe
<folder_name>\Scripts\Activate.psl  # Powershell
../<folder_name>/Scripts/activate   # bash shell
```

```bash
source <folder_name>/bin/activate  # OSX/Linux (bash)
```

## 把包安装在virtual environment中

> Installing packages in a virtual environment.

```cmd
pip install colorma  # Install an individual package
```
```cmd
pip install -r requirments.txt  # Install from a list of packages
```

## 使用清华大学镜像源

可以在使用 pip 的时候加参数 -i https://pypi.tuna.tsinghua.edu.cn/simple

```cmd
pip install colorma -i https://pypi.tuna.tsinghua.edu.cn/simple
```

## 导出依赖

```cmd
pip freeze > requirements.txt
```

## 内存不足

使用 `pip install` 命令时，如果内存不足（比如报 `Memory Error` 错误），请加上 `--no-cache-dir` 以禁用缓存。

可以参考文档 [Disabling caching](https://pip.pypa.io/en/stable/topics/caching/#disabling-caching)
