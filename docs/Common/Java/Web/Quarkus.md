# Quarkus

[Starter](https://code.quarkus.io/) 最小化组合举例： RESTEasy Reactive Jackson、Hibernate Reactive with Panache、Reactive PostgreSQL client

[官网](https://quarkus.io/)

[开始使用](https://quarkus.io/get-started/)

[Guides](https://quarkus.io/guides/)

## 同类产品比较

与 `Micronaut` 和 `Helidon` 相比，`Quarkus` 的社区更活跃。(2021年8月29日)

目前 `Spring Native` 已经进入 beta 状态，可以配合着`Spring Web Flux` 和 `Spring R2DBC`一起用。但用起来还是不太顺手。(2021年8月29日)

## 互补产品

[Kogito](https://kogito.kie.org/) is an open source, end-to-end business process automation(BPA) technology designed to develop, deploy, and execute process- and rules-based cloud-native applications on a modern container platform.

## 环境配置

请使用 `OpenJDK 11` 或 `GraalVM`，不建议使用 `OpenJDK 8` (Plz refer to [Why Dropping Java 8](https://github.com/quarkusio/quarkus/wiki/Why-Dropping-Java-8) for more info)

请使用 `Maven 3.8.1+`

## 使用工具

- [配置IDE](https://quarkus.io/guides/ide-tooling)

- Dev UI  默认链接为 http://localhost:8080/q/dev/

- CLI 参考[这里](https://quarkus.io/guides/cli-tooling)
- 使用 [Maven](https://quarkus.io/guides/maven-tooling)

## WIN 10 排错指北

[quarkus 2.1.4 - mvnw quarkus:dev not working on windows 10](https://stackoverflow.com/questions/68958241/quarkus-2-1-4-mvnw-quarkusdev-not-working-on-windows-10)

## Tips

[Hutool](https://hutool.cn/docs/#/) 是个不错的工具，可以考虑使用它帮助你的开发。视频教程可以参考 [这里](https://www.bilibili.com/video/BV1bQ4y1M7d9?zw)

## ORM

Hibernate ORM with Panache [看这里](https://quarkus.io/guides/hibernate-orm-panache)

## DI

[Context and Dependency Injection](https://quarkus.io/guides/cdi-reference#supported_features)

## 使用 Docker 运行 PostgreSQL 和Quarkus程序

### 启动 PostgreSQL

```zsh
docker run --ulimit memlock=-1:-1 -it --rm=true --memory-swappiness=0 \
    --name postgres-quarkus -e POSTGRES_USER=quarkus \
    -e POSTGRES_PASSWORD=quarkus -e POSTGRES_DB=fruits \
    -p 5432:5432 postgres:13.1
```

### 启动 Quarkus 程序

```zsh
java \
   -Dquarkus.datasource.reactive.url=postgresql://localhost/fruits \
   -Dquarkus.datasource.username=quarkus \
   -Dquarkus.datasource.password=quarkus \
   -jar target/quarkus-app/quarkus-run.jar
```

如果 Quarkus 应用已经被打包成了二进制文件，则执行

```zsh
./target/getting-started-with-reactive-runner \
  -Dquarkus.datasource.reactive.url=postgresql://localhost/fruits \
  -Dquarkus.datasource.username=quarkus \
  -Dquarkus.datasource.password=quarkus
```

也可以使用配置文件指定数据库用户名和密码，详情请参考 [configuration guide](https://quarkus.io/guides/config-reference#configuration_sources)

## Mandrel

Mandrel 是 GraalVM 的一个定制版，专为 Quarkus 设计。它删去了多语言编程的功能和一些其他功能，使得打包体积有所减小。但它目前只支持 building native executables target Linux containerized environments 。

> Mandrel is currently only recommended for building native executables that target Linux containerized environments. This means that Mandrel users should use containers to build their native executables. If you are building native executables for macOS or Windows target platforms, you should consider using Oracle GraalVM instead, because Mandrel does not currently target these platforms. Building native executables directly on bare metal Linux is possible, with details available in the [Mandrel README](https://github.com/graalvm/mandrel/blob/default/README.md) and [Mandrel releases](https://github.com/graalvm/mandrel/releases).

## Reactive

- [参考文档 1: QUARKUS REACTIVE ARCHITECTURE](https://quarkus.io/guides/quarkus-reactive-architecture) 
- [参考文档 2: MUTINY - ASYNC FOR BARE MORTAL](https://quarkus.io/guides/mutiny-primer)

在 [此仓库](https://github.com/quarkiverse) 中寻找插件

### 数据库 Reactive

请在定义实体类时使用父类 `io.quarkus.hibernate.reactive.panache.PanacheEntity`，举个例子：

```java
package org.acme.hibernate.orm.panache;

import javax.persistence.Cacheable;
import javax.persistence.Column;
import javax.persistence.Entity;

import io.quarkus.hibernate.reactive.panache.PanacheEntity;  

@Entity
@Cacheable
public class Fruit extends PanacheEntity {

	 @Column(length = 40, unique = true)
	 public String name;

}
```

### HTTP endpoint Reactive

Thanks to hints in your code (such as the `@Blocking` and `@NonBlocking` annotations), Quarkus extensions can decide when the application logic is blocking or non-blocking.

Mutiny offers two types that are both event-driven and lazy:

- A `Uni` emits a single event (an item or a failure). Unis are convenient to represent asynchronous actions that return 0 or 1 result. A good example is the result of sending a message to a message broker queue.
- A `Multi` emits multiple events (n items, 1 failure or 1 completion). Multis can represent streams of items, potentially unbounded. A good example is receiving messages from a message broker queue.

请在定义 HTTP endpoint 时以 `Uni` 对象作为返回值，比如下面代码中的`Uni<List<Fruit>>`：

```java
package org.acme.hibernate.orm.panache;

import javax.enterprise.context.ApplicationScoped;
import javax.ws.rs.Path;

@Path("/fruits")
@ApplicationScoped
public class FruitResource {
    
    @GET
    public Uni<List<Fruit>> get() {
        return Fruit.listAll(Sort.by("name"));
    }
    
    @POST
	public Uni<Response> create(Fruit fruit) {
    	return Panache.<Fruit>withTransaction(fruit::persist)
            .onItem().transform(inserted -> Response.created(URI.create("/fruits/" + inserted.id)).build());
	}
}
```

When possible, use a `Uni<List<T>>` and load the content. If you have a large set of results, implement pagination.

The code is a bit more involved. To write to a database, we need a transaction. So we use `Panache.withTransaction` to get one (asynchronously) and call the `persist` method when we receive the transaction. The `persist` method is also returning a `Uni`. This `Uni` emits the result of the insertion of the fruit in the database. Once the insertion completes (and that’s our continuation), we create a `201 CREATED` response. RESTEasy Reactive automatically reads the request body as JSON and creates the `Fruit` instance.

Tips: The `.onItem().transform(…)` can be replaced with `.map(…)`. Because `map` is a shortcut.

Test the endpoint:

The test code can be found in [FruitsEndpointTest.java](https://github.com/quarkusio/quarkus-quickstarts/blob/main/hibernate-reactive-panache-quickstart/src/test/java/org/acme/hibernate/orm/panache/FruitsEndpointTest.java)

```
> curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"peach"}' \
  http://localhost:8080/fruits
```

### Observing events

You can observe the various kind of events using:

```java
on{event}().invoke(ev → System.out.println(ev));
```

For example, for items use: `onItem().invoke(item → …);`

For failure, use: `onFailure().invoke(failure → …)`

The `invoke` method is synchronous. Sometimes you need to execute an asynchronous action. In this case use `call`, as in: `onItem().call(item → someAsyncAction(item))`. Note that `call` does not change the item, it just calls an asynchronous action, and when this one completes, it emits the original item downstream.

在 Quarkus 中，vert.x api 被转换为 multiny api，参考[这篇博文](https://quarkus.io/blog/mutiny-vertx/)。

## Build

[Build Native Image](https://quarkus.io/guides/building-native-image)

[Enable SSL](https://quarkus.io/guides/native-and-ssl)

