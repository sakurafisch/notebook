# Druid

## 使用步骤概要

1. 导入 jar 包。
2. 定义 `*.properties` 配置文件，该文件命名任意，目录任意。
3. 加载配置文件 `*.properties`。
4. 获取数据库连接池对象：通过工厂类 `DruidDataSourceFactory` 来获取。
5. 获取连接：`getConnection`

## 示例

```java
package com.winnerwinter.druid;
import com.alibaba.druid.pool.DruidDataSourceFactory;
public class class DruidDemo {
    // 加载配置文件
    Properties properties = new Properties();
    InputStream inputStream = DruidDemo.class.getClassLoader().getResourceAsStream("druid.properties");
    properties.load(inputStream);
    // 获取连接池对象
    DataSource dataSource = DruidDataSourceFactory.createDataSource(properties);
    // 获取连接
    Connection connection = dataSource.getConnection();
    System.out.println(connection);
}
```

## 自定义工具类

提供一个类 JDBCUtils

提供静态代码块加载配置文件，初始化连接池对象

提供方法：

1. 获取连接方法：通过数据库连接池获取连接
2. 释放资源
3. 获取连接池的方法

```java
package com.winnerwinter.utils;

import javax.sql.DataSource;
import java.io.IOException;
import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.Properties;

public class JDBCUtils {
    private static DataSource dataSource;
    static {
        try {
            // 加载配置文件
            Properties properties = new Properties();
            // 获取 DataSource
            properties.load(JDBCUtils.class.getClassLoader().getResourceAsStream("druid.properties"));
            dataSource = DruidDataSourceFactory.createDataSource(properties);
        } catch (IOException e) {
            e.printStackTrace();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    // 获取连接
    public static Connection getConnection() throws SQLException {
        return dataSource.getConnection();
    }

    // 释放资源
    public static void close(ResultSet resultSet, Statement statement, Connection connection) {
        if (resultSet != null) {
            try {
                resultSet.close();
            } catch (SQLException e) {
                e.printStackTrace();
            }
        }

        if (statement != null) {
            try {
                statement.close();
            } catch (SQLException e) {
                e.printStackTrace();
            }
        }

        if (connection != null) {
            try {
                connection.close();
            } catch (SQLException e) {
                e.printStackTrace();
            }
        }
    }
    public static void close(Statement statement, Connection connection) {
        close(null, statement, connection);
    }

    // 获取连接池的方法
    public static DataSource getDataSource() {
        return dataSource;
    }
}
```

```java
package com.winnerwinter.utils;

import java.sql.*;

public class DruidDemo {
    public static void main(String[] args) {
        try {
            // 获取连接
            Connection connection = JDBCUtils.getConnection();
            // 定义 SQL
            String sql = "INSERT INTO account VALUE(null, ?, ?)";
            // 获取 preparedStatement 对象
            PreparedStatement preparedStatement = connection.prepareStatement(sql);
            // 给 ? 赋值
            preparedStatement.setString(1, "张三");
            preparedStatement.setDouble(2, 3000);
            // 执行 SQL
            int count = preparedStatement.executeUpdate();
            System.out.println(count);
        } catch (SQLException e) {
            e.printStackTrace();
        } finally {
            JDBCUtils.close(preparedStatement, connection);
        }
    }
}

```

