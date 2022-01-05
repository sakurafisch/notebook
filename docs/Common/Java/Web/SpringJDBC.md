# Spring JDBC

Spring 框架对 JDBC 的简单封装。

提供了一个 JDBCTemplate 对象简化 JDBC 的开发。

## 使用步骤概要

- 导入 jar 包

- 创建 JdbcTemplate 对象。依赖于数据源 DataSource。

```java
JdbcTemplate jdbcTemplate = new JdbcTemplate(dataSource);
```

- 调用 JdbcTemplate 的方法完成 CRUD 操作

update()：执行DML语句。即增删改语句。

queryForMap()：将查询结果集封装为 map 集合。

queryForList()：将查询结果集封装为 list 集合。

query()：将查询结果封装为 JavaBean 对象。

queryForObject()：将查询结果封装为对象。

## 示例

```java
public class JdbcTemplateDemo {
    public static void main(String[] args) {
        // 创建 JDBCTemplate 对象
        JdbcTemplate jdbcTemplate = new JdbcTemplate(JDBCUtils.getDataSource());
        // 调用方法
        String sql = "UPDATE account SET balance = 5000 WHERE id = ?";
        int count = jdbcTemplate.update(sql, 3);
        System.out.println(count);
    }
}
```