# Maven 常用命令

```bash
mvn clean # 清除产生的项目
mvn compile # 编译源代码
mvn package # 打包
mvn -dmaven.test.skip=true # 打包但跳过测试
mvn install # 安装到本地仓库
```

## 源码打包

```bash
mvn source:jar
或
mvn source:jar-no-fork
```

