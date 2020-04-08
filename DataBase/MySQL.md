# MySQL

## 文件结构

配置文件为my.ini

数据库：文件夹

表：文件

数据：数据

## Windows环境下的CLI

```cmd
services.msc # 打开服务窗口
net start mysql #启动
net stop mysql # 关闭
```

## 登录

```bash
mysql -uroot -p密码
mysql -hip -uroot -p连接目标的密码
mysql --host=ip:host --user=root --password=连接目标的密码
```

## 退出

```bash
exit
quit
```

## 客户端图形化工具SQLYog

名称：cr173

序列号：8d8120df-a5c3-4989-8f47-5afc79c56e7c

或者

名称：cr173

序列号：59adfdfe-bcb0-4762-8267-d7fccf16beda

或者

名称：cr173

序列号：ec38d297-0543-4679-b098-4baadf91f983

## 数据类型摘要

- int：整型

- double：浮点数

- date：日期，只包含年月日，yyyy-MM-dd

- datetime：日期，包含年月日时分秒，yyyy-MM-dd HH:mm:ss

- timestamp：时间戳类型，包含年月日时分秒，yy-MM-dd，若创建时不赋值或赋值为null，自动赋值为当前系统时间
- varchar：字符串类型，例如 varchar(20)

## DDL

### 创建数据库

```mysql
create database <数据库名称>;
```

```mysql
create database if not exists <数据库名称>;
```

```mysql
create database <数据库名称> character set <字符集名>;
```

### 查询数据库

```mysql
show databases;
```

### 查询某个数据库的创建语句

```mysql
show create database <数据库名称>;
```

### 修改数据库的字符集

```mysql
alter database <数据库名称> character set <字符集名称>;
```

### 删除数据库

```mysql
drop database <数据库名称>;
```

```mysql
drop database if exists <数据库名称>;
```

### 查看当前正在使用的数据库名称

```mysql
select database();
```

### 使用数据库

```mysql
use <数据库名称>;
```

### 创建表

```mysql
create table <表名>(
	<列名1> <数据类型1>,
    <列名2> <数据类型2>,
    ...
    <列名n> <数据类型n>
);
```

创建表示例

```mysql
create table student(
	id int primary key,
    name varchar(32),
    age int,
    score double(4, 1),
    birthday date,
    insert_time timestamp
);
```

### 复制表

```mysql
create table <表名> like <被复制的表名>;
```

### 查询某个数据库中所有表的名称

```mysql
show tables;
```

### 查询表结构

```mysql
desc <表名>;
```

### 修改表名

```mysql
alter table <表名> rename to <新的表名>;
```

### 修改表的字符集

```mysql
alter table <表名> character set <字符集名称>;
```

### 添加一列

```mysql
alter table <表名> add <列名> <数据类型>;
```

### 修改列名称或类型

```mysql
alter table <表名> change <列名> <新列名> <新数据类型>;
```

```mysql
alter table <表名> modify <列名> <新数据类型>
```

### 删除列

```mysql
alter table <表名> drop <列名>;
```

### 删除表

```mysql
drop table <表名>;
```

```mysql
drop table if exists <表名>
```

## DML

### 添加数据

```mysql
insert into <表名>(<列名1>, <列名2>, ... <列名n>) values (<值1>, <值2>, ... <值n>);
```

注意：需要把非数字类型用引号引起来。

添加数据，如果不写列名，则默认给所有列添加值

```mysql
insert into <表名> values (值1, 值2, ... ,值n);
```

### 修改数据

```mysql
update <表名> set <列名1> = <值1>, <列名2> = <值2>, ... [where <条件>];
```

注意：如果不加任何条件，则默认修改表中的所有记录。

### 删除数据

```mysql
delete from <表名> [where <条件>]
```

注意：如果不加条件，则默认把所有数据删除。

### 删除表，然后再创建一个一模一样的空表

```mysql
TRUNCATE TABLE <表名>;
```

## DQL

### 语法一览

```mysql
select <字段列表> from <表名列表> where <条件列表> group by <分组字段> having <分组之后的条件> order by <排序字段> <ASC | DESC> limit <分页限定>;
```

### 基础查询

- 多个字段的查询

```mysql
select <列名1>, <列名2>, ... <列名n> from <表名>;
```

```mysql
select * from <表名>;
```

- 去除重复（`DISTINCT`）

```mysql
select distinct <列名> from <表名>;
```

```mysql
select distinct <列名1> <列名2> from <表名> 
```

- 计算列，举例说明：

```mysql
SELECT NAME, math, english, IFNULL(math, 0) + IFNULL(english, 0) FROM student;
```

- 起别名，举例说明：

```mysql
SELECT NAME, math AS 数学, english AS 英语, IFNULL(math, 0) + IFNULL(english, 0) AS 总分 FROM student;
```

- 查询表中的所有记录

```mysql
select * from <表名>
```

### 条件查询

- 使用 WHERE 子句

```mysql
SELECT field1, field2,...fieldN FROM table_name1, table_name2... [WHERE condition1 [AND [OR]] condition2.....
```

以下为操作符列表，可用于 WHERE 子句中。

| 操作符            | 实例                 |
| :---------------- | :------------------- |
| =                 | (A = B) 返回false。  |
| <>, !=            | (A != B) 返回 true。 |
| >                 | (A > B) 返回false。  |
| <                 | (A < B) 返回 true。  |
| >=                | (A >= B) 返回false。 |
| <=                | (A <= B) 返回 true。 |
| BETWEEN...AND     |                      |
| IN （集合）       |                      |
| LIKE （模糊查询） |                      |
| IS NULL           |                      |
| and 或 &&         |                      |
| or 或 \|\|        |                      |
| not 或 !          |                      |

占位符

- _ ：单个任意字符

- % ：多个任意字符

### 排序查询

```mysql
select * from <表名> order by <排序字段1> <ASC | DESC>, <排序字段2> <ASC | DESC> ... ;
```

- ASC：升序，默认。
- DESC：降序。

当前面的排序字段相等时，才会判断后面的排序字段。

### 聚合函数

聚合函数把一列数据作为一个整体，进行纵向的计算。

注意：聚合函数的计算，会排除 `NULL` 值。

- count：计算个数

```mysql
select count(<列名>) from <表名>; 
```

```mysql
select count(IFNULL(<表名>, 0)) from <表名>;
```

```mysql
select count(<主键名>) from <表名>
```

```mysql
select count(*) from <表名>;
```

- max：计算最大值
- min：计算最小值
- sum：计算和
- avg：计算平均值

### 分组查询

分组查询用于统计具有相同特征的数据。

分组之后，查询的内容应为 `分组字段` 或 `聚合函数` 。

```mysql
select <列名>, <聚合函数> from <表名> group by <列名>;
```

```mysql
select <分组字段>, <聚合函数> from <表名> group by <分组字段>;
```

例如，按性别分组，查询男女生的数学平均分和人数：

```mysql
SELECT sex, AVG(math), COUNT(id) FROM student GROUP BY sex;
```

例如，按性别分组，查询男女生的数学平均分和人数，分数低于 70 分的人不参与分组：

```mysql
SELECT sex, AVG(math), COUNT(id) FROM student WHERE math > 70 GROUP BY sex;
```

例如，按性别分组，查询男女生的数学平均分和人数，分数低于 70 分的人不参与分组，分组之后人数要大于 2 人：

```mysql
SELECT sex, AVG(math), COUNT(id) FROM student WHERE math > 70 GROUP BY sex HAVING COUNT(id) > 2;
```

#### WHERE 和 HAVING  有何区别？

- WHERE 在分组之前限定，如果不满足条件，则不参与分组。
- HAVING 在分组之后限定，如果不满足条件，则不会被查询出来。
- WHERE 后不可以跟聚合函数。
- HAVING 可以利用聚合函数进行判断。

### 分页查询

```mysql
select * from <表名> limit <开始的索引>, <每页查询的条数>;
```

公式：`开始的索引` = (`当前页码` - 1) * `每页显示的条数`

例如：

第一页

```mysql
SELECT * FROM student LIMIT 0, 10;
```

第二页

```mysql
SELECT * FROM student LIMIT 10, 10;
```

第三页

```mysql
SELECT * FROM student LIMIT 20, 10;
```

