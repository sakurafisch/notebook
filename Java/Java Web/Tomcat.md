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

## HTTP Response API

### 设置响应行

```java
setStatus(int sc)
```

举个🌰

```java
response.setStatus(302);
```

### 设置响应头

```java
setHeader(String name, String value);
```

举个🌰

```java
response.setHeader("location", "/contextpath/login")
```

### 设置响应体

- 字符输出流

```java
PrintWriter getWriter();
```

- 字节输出流

```java
ServletOutputStream getOutputStream();
```

举个🌰

```java
response.setContentType("text/html;charset=utf-8"); // 设置编码
PrintWriter printWriter = response.getWriter(); // 获取字符输出流
printerWriter.writer("<h1>This is a response message!</h1>"); //输出数据
```

再举一个🌰

```java
response.setContentType("text/heml;charset=utf-8");
ServletOutputStream servletOutputStream = response.getOutputStream();
servletOutputStream.write("你好".getBytes("utf-8"));
```

### 简单的重定向(redirect)方法

重定向的特点：

- 地址栏发生变化。
- 可以访问其他站点（服务器）的资源。
- 两次请求。因此不能用 request 对象共享数据。

```java
response.sendRedirect("/contextpath/register");
```

### 转发(forward)

转发的特点：

- 地址栏路径不变。
- 只能访问当前服务器的资源。
- 一次请求。因此可以使用 request 对象共享数据。

## 验证码

```java
int width = 100;
int height = 50;

// 创建一个对象，在内存中的图片
BufferedImage image = new BufferedImage(width, height, BufferedImage.TYPE_INT_RGB);

// 填充背景色
Graphics graphics = image.getGraphics();// 获取画笔对象
graphics.setColor(Color.PINK); // 设置画笔颜色
graphics.fillRect(0, 0, width, height);

// 画边框
graphics.setColor(Color.BLUE);
graphics.drawRect(0, 0, width -1, height -1);

String str = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
Random random = new Random();// 生成随机角标


for (int i = 1; i <= 4; i++) {
    int index = random.nextInt(str.length());
    // 获取字符
	char ch = str.charAt(index); // 随机字符
	// 写验证码
    graphics.drawString(ch + "", width /5 * i, height / 2);
}

graphics.setColor(Color.GREEN);  // 设置画笔为绿色

for (int i = 0; i < 10; i++) {
    // 随机生成坐标点
	int x1 = random.nextInt(width);
	int x2 = random.nextInt(width);
	int y1 = random.nextInt(height);
	int y2 = random.nextInt(height);
	// 画干扰线
	graphics.drawLine(1, )
}

// 将图片输出到页面显示
ImageIO.write(image, "jpg", response.getOutputStream());
```

## 文件下载

使用响应头设置资源的打开方式

```
content-disposition:attachment;filename=xxx
```

```java
// 获取请求参数，文件名称
String filename = request.getParameter("filename");
// 使用字节输入流加载文件进内存
ServletContext servletContext = this.getServletContext();
String realPath = servletContext.getRealPath("/foldername" + filename);
FileInputStream fileInputStream = new FileInputStream(realPath);
// 设置response响应头
String mimeType = servletContext.getMimeType(filename);
response.setHeader("content-type", mimeType);

// 解决中文文件名问题
String agent = request.getHeader("user-agent");
filename = DownLoadUtils.getFileName(agent, filename);

response.setHeader("content-disposition", "attachment;filename="+filename);

// 将输入流的数据写出到输出流中
ServletOutputStream = servletOutputStream = response.getOutputStream();
byte[] buff = new byte[1024 * 8];
int len = 0;
while((len = fileInputStream.read(buff)) != -1) {
    servletOutputStream.write(buff, 0, len);
}
fileInputStream.close();
```

## 会话技术

会话：一次会话中包含多次请求和响应。

一次会话：浏览器第一次给服务器资源发送请求，会话建立，直到有一方断开为止。

功能：在一次会话的多次请求间，共享数据。

方式：

1. 客户端会话技术：Cookie
2. 服务器端会话技术：Session

## Cookie

概念：客户端会话技术，将数据保存在客户端。

### 使用步骤：

- 创建 Cookie 对象，绑定数据

```java
new Cookie(String name, String value);
```

- 发送 Cookie 对象

```java
response.addCookie(Cookie cookie);
```

- 获取 Cookie，拿到数据

```java
Cookie[] getCookies(); // API 定义，返回 Cookie 数组
```

```java
request.getCookies();  // 使用示例
```

```java
// 遍历 Cookie[] 示例
Cookie[] cs = request.getCookies();
if(cs != null) {
    for (Cookie c : cs) {
        String name = c.getName();
        String value = c.getValue();
        System.out.println(name + ":" + value);
    }
}
```

```java
// 一次请求发送多个 Cookie 示例
Cookie c1 = new Cookie("msg", "hello");
Cookie c2 = new Cookie("name", "winnerwinter");
response.addCookie(c1);
response.addCookie(c2);
```

### 持久化

默认情况下，Cookie存放于内存中，随浏览器关闭而销毁。

持久化存储

```java
setMaxAge(int seconds); // 正数：持久化；负数：默认值；零：删除
```

### 同一服务器的Cookie共享

```java
setPath(String path);
```

举个🌰

```java
Cookie cookie = new Cookie("msg", "Hello");
cookie.setPath("/");
```

### 同一域名不同服务器的Cookie共享

```java
setDomain(String path);
```

举个🌰

```java
Cookie cookie = new ("msg", "Hello");
cookie.setDomain(".google.com");
```

## Session

[HttpSession](https://tomcat.apache.org/tomcat-9.0-doc/servletapi/javax/servlet/http/HttpSession.html)

概念：服务器端会话技术，在一次会话的多次请求间共享数据，将数据保存在服务器端的对象中。

API

```java
Object getAttribute(String name);
void getAttribute(String name, Object value);
void removeAttribute(String name);
```

举个🌰

```java
HttpSession httpSession = request.getSession();
httpSession.setAttribute("msg", "Hello Session");
```

```java
HttpSession httpSession = request.getSession();
Object msg = session.getAttribute("msg");
System.out.println("msg");
```

查找 Session 的 id 依赖于 Cookie。

使浏览器端存储 JSESSIONID 的 Cookie 持久化：

```java
// 使浏览器端存储 JSESSIONID 的 Cookie 持久化
Cookie cookie = new Cookie("JSESSIONID", session.getId());
cookie.setMaxAge(60 * 60);
cookie.addCookie(cookie);
```

Session 的钝化和活化：

- 钝化：在服务器正常关闭之前，将 session 对象序列化到硬盘上。
- 活化：在服务器启动后，将 session 文件转化为内存中的 session 对象。

Tomcat 会自动完成 Session 的钝化和活化过程，但在 IDEA 中活化不会成功，因为 IDEA 启动时会删除原来的 work 目录并重新创建一个。

Session 在以下情况下被销毁：

- 服务器关闭
- session 对象调用 invalidate();
- session 默认失效时间，30分钟

```xml
<!-- web.xml -->
<session-config>
	<session-timeout>30</session-timeout>
</session-config>
```

