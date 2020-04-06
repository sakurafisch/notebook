# Tomcat

## 部署项目的方式

- 直接将项目放到 webapps 目录下。

- 配置conf/server.xml文件

```xml
<!-- 在<Host>标签体中配置 -->
<Context docBase="path/to/app" path="/route" />
```

- 在conf/Catalina/localhost创建任意名称的xml文件。

```xml
<Context docBase="path/to/app" />
```

## WEB-INF目录

- web.xml：项目的配置文件
- classes目录：放置字节码文件的目录
- lib目录：放置依赖的jar包

浏览器不能直接访问WEB-INF目录。

## Tomcat的HTTP Request类

`org.apache.catalina.connector.RequestFacade`类实现了`HttpServletRequest`接口，后者继承了`ServletRequest`接口。

下文介绍`org.apache.catalina.connector.RequestFacade`类的常用方法。

### 获取请求行数据

```
GET /contextpath/demo?name=winnerwinter&age=12 HTTP/1.1
```

- 获取请求方式：GET

```java
String getMethod();
```

- 获取虚拟目录：/contextpath（很常用）

```java
String getContextPath();
```

- 获取Servlet路径：/demo

```java
String getServletPath();
```

- 获取get方式请求参数：name=winnerwinter&age=12

```java
String getQueryString();
```

- 获取请求URI：/contextpath/demo（很常用）

```java
String getRequestURI();  // contextpath/demo
StrinBuffer getRequestURL(); // http://123.222.222.13/contextpath/demo
```

- 获取协议及版本：HTTP/1.1

```java
String getProtocol();
```

- 获取客户机的IP地址：

```java
String getRemoteAddr();
```

### 获取请求头数据

- 通过请求头的名称获取请求头的值（很常用）

```java
String getHeader(String name);
```

举个🌰

```java
String agent = request.getHeader("user-agent");
```

再举一个🌰

```java
String referer = request.getHeader("referer");
```

- 获取所有的请求头名称

```java
Enumeration<String> getHeaderName();
```

举个🌰

```java
Enumeration<String> headerNames = request.getHeaderName();
// 遍历操作
while(headerNames.hasMoreElements()) {
    String name = headerNames.nextElement();
    // 根据key获取value
    String value = request.getHeader(name);
    System.out.println(name + "---" + value");
}
```

### 获取请求体数据

- 要先获取流对象，再从流对象中拿数据

```java
// 获取流对象
BufferedReader getReader(); // 获取字符输入流，只能操作字符数据
ServletInputStream getInputStream();  // 获取字节输入流，可以操作所有类型数据
```

举个🌰

```java
// 获取字符流
BufferedReader bufferedReader = request.getReader();
// 读取数据
String line = null;
while((line = bufferedReader.readline()) != null) {
    System.out.println(line);
}
```

### 获取请求参数的通用方式

下列方法同时适用于get方法和post方法。

```java
String getParameter(String name); //根据参数名称获取参数值
String[] getParameterValues(String name); //根据参数名称获取参数值数组
Enumeration<String> getParameterNames(); // 获取所有请求的参数名称
Map<String, String[]> getParameterMap(); // 获取所有参数的map集合
```

### 获取ServletContext

```java
ServletContext getServletContext();
```

## post方法中文乱码问题

- get方法：Tomcat 8 已经把get方法的中文乱码问题解决了。

- post方法：中文会乱码。

只需在获取参数前设置流的编码

```java
request.setCharacterEncoding("utf-8"); // 设置流的编码
```

## 请求转发

请求转发是一种在服务器内部的资源跳转方式。

### 步骤：

1. 通过request对象获取请求转发器的对象：RequestDispatcher getRequestDispatcher(String path);
2. 通过RequestDispatcher对象来进行转发：forward(ServletRequest request, ServletResponse response);

### 特点：

1. 浏览器地址栏不发生变化。
2. 只能转发到当前服务器内部资源中。
3. 转发是一次请求。

## 共享数据

域对象：一个有作用范围的对象，可以在范围内共享数据。

request域：代表一次请求的范围，一般用于请求转发的多个资源中共享数据。

方法：

```java
void setAttribute(String name, Object obj); // 存储数据
Object getAttribute(String name); // 通过key获取数据
void removeAttribute(String name); // 通过key移除key:value对 
```



