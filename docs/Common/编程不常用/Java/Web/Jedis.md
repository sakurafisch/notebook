# Jedis

[下载jar包](https://mvnrepository.com/artifact/redis.clients/jedis)

## 开箱即用

```java
// 获取连接
Jedis jedis = new Jedis("localhost", 6379);
// 操作
jedis.set("username", "winnerwinter");
// 关闭连接
jedis.close();
```

## 指定会过期的键值对

```java
jedis.setex("key", 20, "value");
```

## 获取 hash 的所有 map 数据

```java
Map<String, String> user = jedis.hgetAll("user");
```

随即可进行遍历操作

```java
Set<String> KeySet = user.keySet();
for (String key : keySet) {
    // 获取 value
    String value = user.get(key);
    System.out.println(key + ": " + value);
}
```

## 获取 List 的所有数据

```java
List<String> mylist = jedis.lrange("mylist", 0, -1);
System.out.println(mylist);
```

## 获取 Set 的所有数据

```java
Set<String> myset = jedis.smembers("myset");
System.out.println(myset);
```

## 获取 SortedSet 的所有数据

```java
Set<String> mysortedset = jedis.zrange("mysortedset", 0, -1);
System.out.println(mysortedset);
```

## Jeids连接池 JedisPool

```java
// 创建一个配置对象
JedisPoolConfig jedisPoolConfig = new JedisPoolConfig();
jedisPoolConfig.setMaxTotal(50);
jedisPoolConfig.setMaxIdle(10);

JedisPool jedisPool = new JedisPool(config, "localhost", 6379);  //创建对象
Jedis jedis = jedisPool.getResource();  // 获取连接
jedis.set("key", "value");  // 使用
jedis.close();  // 关闭，归还到连接池中
```

