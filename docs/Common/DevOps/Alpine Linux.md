# Alpine Linux

可以看看[这个文档](http://gliderlabs.viewdocs.io/docker-alpine/)就是说

## 相关资源

- `Alpine` 官网：[https://www.alpinelinux.org/](https://www.alpinelinux.org)
- `Alpine` 官方仓库：https://github.com/alpinelinux
- `Alpine` 官方镜像：https://hub.docker.com/_/alpine/
- `Alpine` 官方镜像仓库：https://github.com/gliderlabs/docker-alpine

## apk

Alpine Linux 下的包管理工具

```shell
apk install xxx
apk search xxx # 支持正则
apk info xxx # 查看包的详细信息
apk show # list local package
apk del openssh openntp vim # 卸载并删除 包
```

### 升级

upgrade命令升级系统已安装的所以软件包（一般包括内核），当然也可指定仅升级部分软件包（通过-u或–upgrade选择指定）。

```shell
apk update # 更新最新本地镜像源
apk upgrade # 升级软件
apk add --upgrade busybox # 指定升级部分软件包
```

### 搜索

```shell
apk search # 查找所以可用软件包
apk search -v # 查找所以可用软件包及其描述内容
apk search -v 'acf*' # 通过软件包名称查找软件包
apk search -v -d 'docker' # 通过描述文件查找特定的软件包
```

### 查看包信息

info命令用于显示软件包的信息。

```shell
apk info # 列出所有已安装的软件包
apk info -a zlib # 显示完整的软件包信息
apk info --who-owns /sbin/lbu # 显示指定文件属于的包
```

## 获取并使用官方镜像

由于镜像很小，下载时间往往很短，读者可以直接使用 `docker run` 指令直接运行一个 `Alpine` 容器，并指定运行的 Linux 指令，例如：

```shell
docker run alpine echo '123'
```

## 迁移至 `Alpine` 基础镜像

目前，大部分 Docker 官方镜像都已经支持 `Alpine` 作为基础镜像，可以很容易进行迁移。

例如：

- `ubuntu/debian` -> `alpine`
- `python:3` -> `python:3-alpine`
- `ruby:2.6` -> `ruby:2.6-alpine`

另外，如果使用 `Alpine` 镜像替换 `Ubuntu` 基础镜像，安装软件包时需要用 `apk` 包管理器替换 `apt` 工具，如

```shell
apk add --no-cache <package>
```

`Alpine` 中软件安装包的名字可能会与其他发行版有所不同，可以在 `https://pkgs.alpinelinux.org/packages` 网站搜索并确定安装包名称。如果需要的安装包不在主索引内，但是在测试或社区索引中。那么可以按照以下方法使用这些安装包。

```shell
$ echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories
$ apk --update add --no-cache <package>
```

由于在国内访问 `apk` 仓库较缓慢，建议在使用 `apk` 之前先替换仓库地址为国内镜像。

```shell
RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \
      && apk add --no-cache <package>
```

## 初始化配置

```shell
apk add iproute2 # ss vs netstat
ss -ptl
apk add drill # drill vs nslookup&dig

crond # 开启 cron 服务
crontab -l -e

apk add xxx
apk search -v xxx
apk info -a xxx
apk info
echo -e "http://mirrors.aliyun.com/alpine/v3.6/main\nhttp://mirrors.aliyun.com/alpine/v3.6/community" > /etc/apk/repositories
apk update

# storage
ibu # alpine local backup

# network
echo "shortname" > /etc/hostname
hostname -F /etc/hostname
/etc/hosts
/etc/resolv.conf # conig DNS
modprobe ipv6 # enable ipv6
echo "ipv6" >> /etc/modules
iface # config interface
apk add iptables ip6tables iptables-doc
/etc/init.d/networking restart # activate change
apke add iputils # IPv6 traceroute
traceroute6 ipv6.google.com
awall # alpine wall
# setup a openvpn server

# post-install
/etc/apk/repositories
apk add cherokee --update-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ --allow-untrusted
apk search -v --description 'NTP' # show description and search from description
apk info -a zlib
apk info -vv|sort
apk info -r -R # require / depency
apk version -v -l '<' # show available updates
apk upgrade -U -a
apk add -u xxx # update xxx

/etc/runlevels # runlevel
apk add openrc # use openrc for init system
rc-update add xxx # set to start on
rc-service xxx start # equal -> /etc/init.d/xxx start
rc-status

adduser xxx
passwd xxx

apk add ansible # server
ssh-keygen
/etc/ansible/hosts
apk add python # node
ssh-copy-id

apk add man man-pages mdocml-apropos less less-doc
export PAGER=less
/etc/rc.conf # /etc/rc.conf -> funny character
apk add bash bash-doc bash-completion # bash
apk add util-linux pciutils usbutils coreutils binutils findutils grep # grep / awk
apk add build-base gcc abuild binutils binutils-doc gcc-doc # compile
apk add cmake cmake-doc extra-cmake-modules extra-cmake-modules-doc
apk add ccache ccache-doc

apk add docker # docker
rc-update add docker boot
rc-service docker start
apk add py-pip
pip install docker-compose
ln -s /usr/bin/docker-compose /usr/bin/doc

# application
apk add openssh # ssh
rc-update add sshd
/etc/init.d/sshd start
/etc/sshd_config
apk add dropbear # another openssh implementation
```

