# 编译 cpython 3.7

## 在 Ubuntu 中编译 cpython 3.7

当系统中缺少一部分编译时所需的库时，编译过程依然会成功，并且能通过所有单元测试。但是在实际使用 python 时却又会出错。

所以务必在编译 cpython 前安装完整的编译工具链。

```zsh
sudo apt-get update
sudo apt-get upgrade
sudo apt-get dist-upgrade
sudo apt-get install build-essential python-dev python-setuptools python-pip python-smbus
sudo apt-get install libncursesw5-dev libgdbm-dev libc6-dev
sudo apt-get install zlib1g-dev libsqlite3-dev tk-dev
sudo apt-get install libssl-dev openssl
sudo apt-get install libffi-dev
```

```zsh
make clean
./configure --with-ssl --enable-optimizations
make -j 16
make test
make altinstall
```
