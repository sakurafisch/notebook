# Docker

> 备忘： ps aux 查看Linux进程

## 运行新的nginx

```shell
docker container run --publish 80:80 nginx
```

## 在后台运行新的nginx

```shell
docker container run --publish 80:80 --detach nginx
```

## 在后台运行新的nginx并对其命名

```shell
docker run --publish 80:80 --detach --name <nickname> nginx
```

## 运行旧的容器

```shell
docker start <nickname>
```

## 查看运行的容器

```shell
docker container ls
docker container ps
```

## 停止后台的容器

```shell
docker container stop <容器ID前面数位>
```

## 查看之前运行过的容器

```shell
docker container ls -a
```

## 打印某个容器的日志

```shell
docker container logs <nickname>
```

## 查看某个容器中的进程

```shell
docker container top <nickname>
```

## 移除某个容器

```shell
docker container rm <容器ID 1> <容器ID 2> ... <容器ID n>
```

## 强制移除某个容器

```shell
docker container rm -f <容器ID 1> <容器ID 2> ... <容器ID n>
```

