# Servlet

## 简介

Servlet本质上是一个interface，约定了Tomcat执行服务程序的方式。

Servlet 是 Server Applet 的简写，意为运行在服务端的小程序。

## Servlet版本

2009年，Servlet3.0开始支持注解配置，可以不需要web.xml了，使用`@WebServlet`对类名注解即可。Java EE 6。

```
@WebServlet(urlPatterns='/route')
```
当只有`urlPatterns`这一个属性要配置时，可以写成：
```
@WebServlet('/route')
```

一个Servlet类可以设置多个路由：

```
@WebServlet('/route1', 'route2', 'route3')
```

2013年，Servlet3.1支持[WebSocket](https://zh.wikipedia.org/wiki/WebSocket)。Java EE 7。

2017年，Servlet4.0支持[HTTP/2](https://zh.wikipedia.org/wiki/HTTP/2)。Java EE 8。

## Servlet生命周期

- init ()： 初始化方法，在Servlet被实例化时调用。默认情况下，第一次被访问时被实例化。可以手动把它改成在服务程序启动时实例化。

```xml
<!-- web.xml -->
<servlet>
	<load-on-startup>5</load-on-startup>
    <!-- load-on-startup标签值为负数时，在第一次访问时实例化。否则（为0或正整数），在服务程序启动时实例化 -->
</servlet>
```

Servlet的init方法只执行一次，在内存中只存在一个对象，即Servlet是单例的。多个用户同时访问时，可能存在线程安全问题，因此尽量不要在Servlet中定义成员变量。即使定义了成员变量，也不要对其进行写操作。

- service()：处理客户端的请求，每访问一次就执行一次。
- destroy()：服务端应用正常关闭前调用，一般用于释放资源。

Servlet 是由 JVM 的垃圾回收器进行垃圾回收的。

## GenericServlet

GenericServlet是一个抽象类，它对Servlet接口的其他方法做了空实现，仅保留Service()方法作为抽象方法。

## HttpServlet

## ServletContext对象

概念：代表整个Web应用，可以和程序的容器（服务器）通信

获取方式：

```java
// 通过 request 对象获取
request.getServletContext();
```

```java
// 通过 HttpServlet 获取
this.getServletContext();
```

功能：

- 获取 MIME 类型

```java
String getMimeType(String file);
```

- 域对象：共享数据

```java
setAttribute(String name, Object value);
getAttribute(String name);
removeAttribute(String name);
```

- 获取文件的真实（服务器）路径

```java
String getRealPath(String path);
```

