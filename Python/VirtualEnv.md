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

我们首先要激活它。（Firstly we need to ```activate``` it.）

要把激活文件放在<folder_name>/Scripts文件夹中。

```cmd
# Windows systems
<folder_name>\Scripts\Activate.bat  # cmd.exe
<folder_name>\Scripts\Activate.psl  # Powershell
../<folder_name>/Scripts/activate   # bash shell
```

```bash
<folder_name>/bin/activate  # OSX/Linux (bash)
```

## 把包安装在virtual environment中

> Installing packages in a virtual environment.

```cmd
pip install colorma  # Install an individual package
```
```cmd
pip install -r requirments.txt  # Install from a list of packages
```

