# Redis

## 目录结构

redis.windows.conf：配置文件

redis-cli.exe：redis 的客户端

redis-server.exe：redis 的服务器端

## 数据结构

redis 存储的是 key: value 格式的数据，其中 key 都是字符串，value 有五种不同的数据结构。

value 的数据结构：

- 字符串类型 string
- 哈希类型 hash
- 列表类型 list
- 集合类型 set
- 有序集合类型 sortedset

## 字符串类型的操作

```bash
set key value
get key
del key
```

## 哈希类型的操作

```bash
hset key field value
hgetall key
hget key field
hdel key field
```

## 列表类型的操作

```bash
lpush key value
rpush key value
lrange key start end # 在某范围获取 若end为-1表示全部获取
lpop key
rpop key
```

## 集合类型的操作

```bash
sadd key value # 存储
smembers key # 获取所有元素
srem key value # 删除某个元素
```

## 有序集合类型的操作

```bash
zadd key score value # 存储
zrange key start end # 在某范围获取
zrem key value # 删除某个元素
```

## 通用命令

```bash
keys * # 查询所有的键
type key # 获取键对应的 value 类型
del key # 删除指定的key value
```

## 持久化

### RDB

RDB 是默认方式。

在一定的时间间隔中，检测 key 的变化情况，然后持久化数据

编辑 redis.windows.conf 文件以修改配置

```conf
save 900 1
save 300 10
save 60 10000
```

重启redis服务器，并指定配置文件名称

```cmd
redis-server.exe redis.windows.conf
```

### AOF

日志记录方式。

可以记录每一条命令的操作。每一次命令操作后，持久化数据。

编辑 redis.windows.conf 文件以修改配置

```conf
appendonly yes # 默认值为 no
```

```conf
appendfsync always # 每次操作都进行持久化
appendfsync everysec # 每一秒持久化一次
appendfsync no # 不持久化
```

重启redis服务器，并指定配置文件名称

```cmd
redis-server.exe redis.windows.conf
```