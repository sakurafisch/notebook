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
mysql -uroot -p<密码>
mysql -hip -uroot -p<连接目标的密码>
mysql --host=ip:host --user=root --password=<连接目标的密码>
```

## 退出

```bash
exit
quit
```

## 数据库的备份和还原

### 备份

```mysql
mysqldump -u<用户名> -p<密码> <数据库名称> > <要保存的路径>
```

### 还原

- 先登录

- 创建数据库

```mysql
create database <数据库名称>;
```

- 使用数据库

```mysql
use database <数据库名称>;
```

- 执行文件

```mysql
source <文件路径>;
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

## DDL操作数据库和表

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

## DML增删改表中数据

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

## DQL查询表中数据

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

## DCL管理用户和授权

### 添加用户

```mysql
CREATE USER '<用户名>'@'<主机名>' IDENTIFIED BY '<密码>';
```

### 删除用户

```mysql
DROP USER '<用户名>'@<主机名>;
```

### 修改用户密码

```mysql
UPDATE USER SET PASSWORD = PASSWORD('<新密码>') WHERE USER = '<用户名>';
```

```mysql
SET PASSWORD FOR '<用户名>'@'<主机名>' = PASSWORD('新密码');
```

### 查询用户

- 切换到 `mysql` 数据库

```mysql
USE mysql;
```

- 查询 `user` 表

```mysql
SELECT * FROM USER;
```

通配符 `%` 表示可以在任意主机使用该用户登录数据库。

### 忘记 root 密码

- 先停止 MySQL 服务

```cmd
net stop mysql # 使用 cmd 停止 MySQL 服务
```

```bash
sudo systemctl stop mysql.service # 使用 bash 停止 MySQL 服务
```

- 使用无验证方式启动 MySQL 服务，保持该命令行窗口打开

```cmd
mysql --skip-grant-tables
```

- 新开一个命令行窗口并进入 MySQL

```cmd
mysql
```

- 更改 `root` 用户的密码

```mysql
UPDATE USER SET PASSWORD = PASSWORD('<新密码>') WHERE USER='root';
```

- 结束MySQL服务，然后重新开启

### 查询用户权限

```mysql
SHOW GRANTS FOR '<用户名>@<主机名>';
```

### 授予用户权限

```mysql
GRANT <权限列表> ON <数据库名.表名> TO '<用户名>'@'<主机名>';
```

权限有以下这些可选项：

- SELECT | DELETE | UPDATE | ALL

给某用户授予所有权限，在任意数据库任意表上：

```mysql
GRANT ALL ON *.* TO '<用户名>'@<主机名>;
```

### 撤销用户权限

```mysql
REVOKE <权限列表> ON <数据库名.表名> FROM '<用户名>'@'<主机名 >'; 
```

## 约束

约束的概念：对表中的数据进行限定，保证数据的正确性、有效性和完整性。

分类：

- 主键约束：primary key
- 非空约束：not null
- 唯一约束：unique
- 外键约束：foreign key

### 主键约束PRIMARY KEY

主键：

1. 非空且唯一
2. 一张表只能有一个字段为主键
3. 主键是表中记录的唯一标识

- 在创建表时添加主键

```mysql
CREATE TABLE stu(
	id INT PRIMARY KEY,
    NAME VARCHAR(20)
);
```

- 在创建表后添加主键

```mysql
ALTER TABLE stu MODIFY id INT PRIMARY KEY;
```

- 删除主键

```mysql
ALTER TABLE stu DROP PRIMARY KEY;
```

#### 自动增长AUTO_INCREMENT

概念：如果某一列是数值类型的，则可用 auto_increment 实现值的自动增长。

- 在创建表时设置自动增长

```mysql
CREATE TABLE stu(
	id INT PRIMARY KEY AUTO_INCREMENT,
    NAME VARCHAR(20)
);
```

- 添加自动增长

```mysql
ALTER TABLE stu MODIFY id AUTO_INCREMENT;
```

- 删除自动增长

```mysql
ALTER TABLE stu MODIFY id INT;
```

### 非空约束NOT NULL

- 在创建表时添加约束

```mysql
CREATE TABLE stu(
	id INT,
	NAME VARCHAR(20) NOT NULL
);
```

- 在创建表后，添加非空约束

```mysql
ALTER TABLE stu MODIFY NAME VARCHAR(20) NOT NULL;
```

- 删除 NAME 的非空约束

```mysql
ALTER TABLE stu MODIFY NAME VARCHAR(20);
```

### 唯一约束UNIQUE

注意：唯一约束限定的列的值可以有多个 NULL

- 在创建表时，添加唯一约束

```mysql
CREATE TABLE stu (
	id INT;
    phone_number VARCHAR(20) UNIQUE
);
```

- 在创建表后，添加唯一约束

当数据有重复时，会导致添加唯一约束失败

```mysql
ALTER TABLE stu MODIFY phone_number VARCHAR(20) UNIQUE;
```

- 删除唯一约束

```mysql
ALTER TABLE stu DROP INDEX phone_number VARCHAR(20);
```

### 外键约束FOREIGN KEY

外键约束：让表与表产生关系，从而保证数据的正确性。

- 在创建表时添加外键

```mysql
CREATE TABLE <表名>(
	... ,
    CONSTRAINT <外键名称> FOREIGN KEY (<外键列名称>) REFERENCES <主表名称>(<主表列名称>)
) ;
```

- 在创建表后添加外键

```mysql
ALTER TABLE <表名> ADD CONSTRAINT <外键名称> FOREIGN KEY (<外键字段名称>) REFERENCES <主表名称>(<主表列名称>);
```

- 删除外键

```mysql
ALTER TABLE <表名> DROP FOREIGN KEY <外键名称>;
```

#### 级联操作(谨慎使用)

- 级联更新：ON UPDATE CASCADE
- 级联删除：ON DELETE CASCADE

```mysql
ALTER TABLE employee ADD CONSTRAINT emp_dept_fk FOREIGN KEY (dep_id) REFERENCES department(id) ON UPDATE CASCADE;
```

## 多表之间的关系

### 一对一

举例：一个人只有一个身份证，一个身份证只能对应一个人。

实现：在任意一张表添加外键指向另一张表的主键，并对该外键添加唯一约束(`UNIQUE`)。

### 一对多

举例：一个部门有多个员工，一个员工只能对应一个部门。

实现：在多的一方建立外键，指向一的一方的主键。

### 多对多

举例：一个学生可以选择多门课程，一门课程可以被很多学生选择。

实现；建立一张表，比如学生选课表，联合主键是sid和cid的组合。

## 3NF范式设计思路

能创建表：第一范式（1NF）

在1NF的基础上，消除非主属性对主码的依赖：第二范式（2NF）

在2NF的基础上，消除传递依赖：第三范式（3NF）

## 多表查询

### 笛卡尔积

```mysql
select * FROM <表A>, <表B>;
```

### 内连接查询

#### 隐式内连接

使用 `where` 条件消除无用数据。

举例，查询所有员工信息和对应的部门信息

```mysql
SELECT * FROM emp, dept WHERE emp.`dept_id` = dept.`id`;
```

#### 显式内连接

```mysql
SELECT <字段列表> FROM <表A> INNER JOIN <表B> ON <条件>
```

举例

```mysql
SELECT * FROM emp INNER JOIN dept ON emp.`dept_id` = dept.`id`;
```

### 外连接查询

#### 左外连接

```mysql
SELECT <字段列表> FROM <表A> LEFT OUTER JOIN <表B> ON <条件>
```

#### 右外连接

```mysql
SELECT <字段列表> FROM <表A> RIGHT OUTER JOIN <表B> ON <条件>
```

### 子查询

若查询中包含嵌套查询，则其中的嵌套查询称为子查询。

#### 结果为单行单列的子查询

当子查询的结果为单行当列时，子查询可以作为条件，使用运算符去判断。

举例，查询工资最高的员工信息：

```mysql
SELECT * FROM emp WHERE emp.`salary` = (SELECT MAX(salary) FROM emp);
```

再举一例，查询工资低于平均工资的员工信息

```mysql
SELECT * FROM emp WHERE emp.salary < (SELECT AVG(salary) FROM emp);
```

#### 结果为多行单列的子查询

当子查询的结果为多行单列时，子查询可以作为条件，使用运算符 `IN` 去判断。

举例，查询财务部和市场部所有员工的信息

```mysql
SELECT * FROM emp WHERE dept_id IN (SELECT id FROM dept WHERE NAME = '财务部' OR NAME = '市场部');
```

#### 结果为多行多列的子查询

当子查询的结果为多行多列时，子查询可以作为一张虚拟表。

举例，查询入职日期为2011-11-11之后的员工信息和部门信息：

```mysql
SELECT * FROM dept AS table1, (SELECT * FROM emp WHERE emp.`join_date` > '2011-11-11') AS table2 WHERE table1.id = table2.id;
```

与内连接作对比：

```mysql
SELECT * FROM emp AS table1, dept AS table2 WHERE table1.`dept_id` = t2.`id` AND t1.`join_date` > '2011-11-11';
```

## 事务

### 事务的基本介绍

#### 事务的概念

如果一个包含多个步骤的业务操作，被事务管理，那么这些操作要么同时成果，要么同时失败。

#### 事务的基本操作

- 开启事务

```mysql
start transaction;
```

- 回滚

```mysql
rollback;
```

- 提交

```mysql
commit;
```

#### MySQL 事务的默认提交方式

在MySQL数据库中事务默认自动提交（在Oracle数据库中事务默认手动提交），一条 DML 语句会自动提交一次事务；当执行 `start transaction;` 后，本次事务变为手动提交。

可以查看 DML 语句的事务是否设置为自动提交：

```mysql
SELECT @@autocommit; -- 1 代表自动提交，0 代表手动提交
```

可以更改事务的默认提交方式：

```mysql
SET @@autocommit = 0; -- 设为手动提交
```

设为手动提交后，如果执行了 DML 语句而没有手动提交，当退出登录后后台会自动回滚。

### 事务的四大特征

1. 原子性
2. 持久性
3. 隔离性
4. 一致性

### 事务的隔离级别

#### 要解决的问题

事务的不同隔离级别，是为了解决这些问题：

- 脏读：一个事务读取到另一个事务中没有提交的数据
- 不可重复读（虚读）：在同一个事务中，两次读取到的数据不一样。
- 幻读：一个事务操作（DML）数据表中所有记录，另一个事务添加了一条数据，则第一个事务查询不到自己的修改。

#### 隔离级别

- read uncommited
- read committed（Oracle默认）：解决了脏读
- repeatable read（MySQL默认）：解决了脏读、不可重复读
- serializable：解决了所有问题

事务的级别由小到大安全性越来越高，但效率越来越低。

### 事务的相关操作

数据库查询隔离级别

```mysql
select @@tx_isolation;
```

数据库设置隔离级别

```mysql
set global transaction isolation level <隔离级别字符串>;
```

注意：当改变数据库的事务隔离级别后，要重新登录才能生效。