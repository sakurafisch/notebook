# JDBC以前的样子

## 简洁

JDBC 是 Java DataBase Connectivity 的缩写。

JDBC 定义了一套操作所有关系型数据库的接口；各个数据库厂商实现这套接口，提供数据库驱动 jar 包。

我们可以使用 JDBC 接口编程，而真正执行的代码是驱动 jar 包中的类实现。

本笔记记叙了JDBC以前的样子，现在的JDBC有点不一样了。

## 使用步骤概要

1. 导入驱动 jar 包
2. 注册驱动
3. 获取数据库连接对象 Connection
4. 定义 SQL
5. 获取执行 SQL 语句的对象 Statement
6. 执行 SQL，接受返回结果
7. 处理结果
8. 释放资源

## 示例

```java
public class JdbcExample {
    public static void main(String[] args) throws Exception {
        // 注册驱动
        Class.forName("com.mysql.jdbc.Driver");  
        // 获取数据库连接对象
        Connection connection = DriverManager.getConnection("jdbc:mysql://localhost:3306/NameOfYourDB", "root", "password");
        // 定义 SQL 语句
        String sql = "UPDATE account SET balance = 500 WHERE id = 1";
        // 获取执行 SQL 语句的对象 Statement
        Statement statement = connection.createStatement();
        // 执行 SQL
        int resultCode = statement.executeUpdate(sql);
        // 处理结果
        System.out.println(resultCode);
        // 释放资源
        statement.close();
        connection.close();
    }
}
```

如果有错，可以考虑用 try catch finally 语句处理一下。

## 驱动管理对象DriverManager

功能：

- 注册驱动

```mysql
static void registerDriver(Driver driver);
```

在实际开发中，这样写就行：

```mysql
Class.forName("com.mysql.jdbc.Driver");-
```

因为在`com.mysql.jdbc.Driver类`中有如下静态代码块：

 ```java
static {
    try {
        java.sql.DriverManager.registerDriver(new Driver());
    } catch (SQLException E) {
        throw new RunTimeException("Can't register driver!");
    }
}
 ```

 MySQL5 之后的驱动 jar 包可以省略注册驱动的步骤。

- 获取数据库连接

MySQL 8.0 是不需要建立 SSL 连接的，需要显式关闭，即 url 中 `useSSL=false;`

注意：时区不设置会报错。

方法：

```java
static Connection getConnection(String url, String user, String password);
```

例如：

```java
Connection connection = DriverManager.getConnection("jdbc:mysql://localhost:3306/NameOfYourDB", "root", "password");

```

## 数据库连接对象Connection

功能

- 获取执行 SQL  的对象

```java
Statement createStatement();
```

```java
PreparedStatement preparedStatement(String sql);
```

- 管理事务

```java
void setAutoCommit(boolean autoCommit); // 传参为 false 时，开启事务
```

```java
void commit(); // 提交事务
```

```java
void rollback(); // 回滚
```

## 执行SQL的对象Statement

```mysql
boolean execute(String sql); // 可以执行任意的 SQL
```

返回值：

- 当第一个结果是一个 ResultSet 对象，返回 `true`
- 当没有结果或是更新更新计数，返回 `false`

```java
int executeUpdate(String sql); // 执行 DML 或 DDL 语句
```

返回值：影响的行数，可以通过这个影响的行数判断 DML 语句是否执行成功，如果返回值大于0则执行成功，反之失败。

```mysql
ResultSet executeQuery(String sql); // 执行 DQL 语句
```

返回值：

## 结果集对象ResultSet

```java
boolean next(); // 游标向下移动一行，超出行尾返回 false
```

```java
int getInt(int columnIndex); // 获取 int 类型
int getInt(String columnLabel);
String getString(int columnIndex); // 获取 String 类型
String getString(String columnLabel);
```

举例：

```java
ResultSet rs = null;
Connection conn = something;
Statement stmt = connection.createStatement();
sql = "string of your sql";
rs = stmt.executeQuery(sql);

while(rs.next()) {
    int id = rs.getInt(1);
    String name = rs.getString("name");
    double balance = rs.getDouble(3);
    
    System.out.println(id + "---" + name + "---" + balance);
}
```

## 执行SQL的对象PreparedStatement

预编译的 SQL，参数使用 `?` 占位符替代，可以预防 SQL 注入。

SQL注入示例

```mysql
select * from user where username = 'whatEverYouWrite' and password = `whatEverYouWrite` or 'simpleString' = 'simpleString';
```

因为后面的恒等式存在，所以该语句可执行。

使用示例

```java
public static void main(String[] args) {
    Scanner sc = new Scanner(System.in);
    System.out.println("请输入用户名：");
    String username = sc.nextLine();
    System.out.println("请输入密码");
    String password = sc.nextLine();
    boolean flag = login(username, password);
    if (flag) {
        System.out.println("登录成功!");
    } else {
        System.out.println("用户名或密码错误!");
    }
}


public boolean login(String username, String password) {
    if(username == null || password == null) {
        return false;
    }
    Connection connection = null;
    PreparedStatement pastmt = null;
    ResultSet rs = null;
    try {
        connection = JDBCUtils.getConnection();
        String sql = "SELECT * FROM user WHERE username = ? AND password = ?";
        pstmt = connection.prepareStatement(sql);
        
        // 给 ? 赋值
        pstmt.setString(1,username);
        pstmt.setString(2, password);
        
        rs.connection.prepareStatement(sql);
        return rs.next();
    } catch (SQLException e) {
        e.printStackTrace();
    } finally {
        JDBCUtils.close(rs, ptmt, conn)
    }
}
```

