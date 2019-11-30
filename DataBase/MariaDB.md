# MariaDB�ʼ�

## Ĭ���趨

Ĭ��ϵͳ���ݿ���û����ݿ�

- information_schema
- performance_schema
- mysql
- test..

Ĭ���ַ���

- �ַ���latin1����ŷ���ԣ�
- �Ƚ���latin1_swedish_ci���ɸ���Ϊutf8, utf8_general_ci��

## �������ͺ������

### �����

- `=` ��ϵ�Ƚ����
- `<>`, `!=` ������

ע���ϵ�Ƚ������һ��`=`��

### [��������](https://mariadb.com/kb/en/library/data-types/)

#### ��ֵ��������

- `[TINY | SMALL | MEDIUM]INT[(M)]` �ֱ�ռ��1��2��3��4�ֽڴ��棬����M��ʾ��ʾ��ȣ���ʵ��ռ�õĿռ��С�޹ء�
  - INT [SIGNED | UNSIGNED | ZEROFILL] Ĭ��Ϊ�з��ţ������`ZEROFILL`���ȼ���`UNSIGNED ZEROFILL`��
  - `ZEROFILL`ʹ����ֵ����ʾʱδ�ﵽ`M`��ָ��λ������ǰ������㣻������ָ������ʾ���ʱ���ᰴ��ԭ�������������ضϣ���

- [`DECIMAL(M, D)`](https://mariadb.com/kb/en/library/decimal/) ��������MΪ��λ����DΪС��λ����
- `FLOAT[(M, D)]`, `DOUBLE[(M, D)]` �����ȸ�������˫���ȸ�������

#### ����ֵ��������

- `CHAR(M)`, `VARCHAR(M)`, `TINYTEXT`, `TEXT`
- `DATE`, `DATETIME`, `TIME`, `TIMESTAMP`, `YEAR`

�����ָ��CHAR�Ŀ�ȣ�Ĭ��ֵ�������Ե�1��

### ���ú���������

- `NOW()`, `DATETIME`
- `CURRENT_TIMESTAMP`, `TIMESTAMP`

## �û�����

- �����û� `CREATE USER 'user' IDENTIFIED BY 'password';`
- ����Ȩ�� `GRANT ALL PRIVILEGES ON dbname.* TO uesr;`, `GRANT SELCET, INSERT, UPDATE, DELETE, GRANT OPTION on dbname TO user;`

`GRANT OPTION`ʹ�û�ȡ��Ȩ�޿��Դ��ݡ�

## DDL

- ����ע�� `//`
- ����ע�� `/* ... */`

### �������ݿ�

- `SHOW DATABASES;`
- `USE ���ݿ���;`
- `CREATE DATABASE ���ݿ���;`
- `DROP DATABASE ���ݿ���;`

### �������ݱ�

#### ������

```sql
CREATE TABLE [IF NOT EXISTS] 'tbname' (
    'colName' INT(4) PRIMARY KEY COMMENT '����ǰ��û�еȺ�',
    �ֶ��� �������� [�ֶ����� | Լ��] [����] [ע��],
    ...,
) [COMMENT='����ǰ���еȺ�'] [CHARSET=...];
```

ע��CREATE TABLE��������һ�������������һ�У���Ҫ���Ӷ��š�

`SHOW CREATE TABLE ����[\G]` �鿴����������

���ڶ��ֶ������������ȶ������ֶΣ��ٶ�������`PRIMARY KEY(id, name)`

������� `FOREIGN KEY(col1, col2) REFERENCES ������ (col1, col2)`

`ALTER TABLE DROP FOREIGN KEY ���Լ������`

#### ά����

- `SHOW TABLES`
- `DESCRIBE ����` `DESC ����` �鿴��Ķ��壨���ж�����ֶΣ������ܲ鿴���ݿ�ġ����塱
- `DROPE TABLE [IF EXISTS] ����`

������������ֶε��޸�

- `ALTER TABLE �ɱ��� RENAME [TO] �±���;` [TO]��ʡ����Ӱ��
- `ALTER TABLE ���� ADD �ֶ��� �������� [����]`
- `ALTER TABLE ���� MODIFY �ֶ��� ��������` �ı��ֶ���������
- `ALTER TABLE ���� MODIFY �ֶ��� �������� FIRST | AFTER �ֶ���` �ı��ֶε�����˳��
- `ALTER TABLE ���� CHANGE ԭ�ֶ��� ���ֶ��� �������ͣ�����Ϊ�գ� [����]` �ı��ֶ������ֶ���������
- `ALTER TABLE ���� DROP �ֶ���`

�������Լ�� `ALTER TABLE ���� ADD CONSTRAINT �ֶ��������磬pk_������ PRIMARY KEY ����(�ֶ���)`

����ֶ�Լ�� `ALTER TABLE ���� ADD CONSTRAINT �ֶ��������ݣ�fk_�ӱ���_�������� FOREIGN KEY(����ֶ�) REFERENCES �������������ֶΣ�`

### �ֶ�Լ��

- �ǿ�Լ�� NOT NULL
- Ĭ��Լ�� DEFAULT �磺DEFAULT '��'
- ΨһԼ�� UNIQUE KEY
- ����Լ�� PRIMARY KEY
- ���Լ�� FOREIGN KEY �ӱ�ĳ�������������б������
- �Զ����� AUTO_INCREMENT

InnoDB֧�������MyISAM��֧���������������ı�Ҫ����InnoDB��

UNSIGNEDҲ��һ��Լ��

## DML

### ��Ĵ�������

InnoDB, MyISAM, Memory, Archive

```sql
CREATE TABLE tbname (
    # ʡ�Դ���
) ENGINE=InnoDB;
```

`ALTER TABLE ���� ENGINE=<������>`

### INSERT

- `INSERT INTO ���� [�ֶ����б�] VALUES (ֵ�б�);`
- `INSERT INTO ���� [�ֶ����б�] VALUES (ֵ�б�1), (ֵ�б�2), ...;`

����ѯ������뵽�±�

```sql
CREATE TABLE �±� (
    SELECT �ֶ�1, �ֶ�2, ... FROM ԭ��
);
```

### UPDATE

- `UPDATE ���� SET �ֶ��� = ����ֵ`

### DELETE

- `DELETE FROM ���� [WHRER ... ]`
- `TRUNCATE TABLE ����` ��־�л����¿�ʼ���

## Chapter 6 DQL

```sql
SELECT <�ֶ��� | ���ʽ | ���� | ����>
FROM ����
[WHERE <��ѯ����>]
[ORDER BY <���������> [ASC(Ĭ��) | DESC]]
```

### SELECT

- `SELECT �ֶ��� AS ����`
- `SELECT �ֶ�1+�ֶ�2 AS ����`
- `WHERE �ֶ��� IS NULL` ��ѯ��ֵ
- `SELECT ���� AS �ֶ���`

### ORDER BY

`ORDER BY �ֶ�1 [����], �ֶ�2[����]`

### LIMIT

`LIMIT [[λ��ƫ����offset], ����row_count]`

��1����¼��λ��ƫ������0��

### �Ӳ�ѯ

ʹ�õȺŵ��Ӳ�ѯ��ΪWHERE����һ���֣���Ҫ��֤�Ӳ�ѯ���ص�ֵΪ�ջ���ֻ��һ��

```sql
SELECT 'studentName' FROM 'student' WHERE 'studentNo' = (
    SELECT 'studentNo' FROM 'result'
        INNER JOIN 'subjcet' ON result.subjectNO = subject.subjectNo
        WHERE 'studentResult'=60 AND 'subjectName'='LogicJava'
);
```

### IN

����Ӳ�ѯ���ص�ֵ����һ��������ʹ��`IN`

```sql
SELECT 'studentName' FROM 'student' WHERE 'studentNo' IN (
    SELECT 'studentNo' FROM 'result'
        INNER JOIN 'subjcet' ON result.subjectNO = subject.subjectNo
        WHERE 'studentResult'=60 AND 'subjectName'='LogicJava'
);
```

����`NOT IN`��

### BETWEEN AND

`SELECT name FROM student WHERE id BETWEEN 1706300001 AND 1806300001`

����`NOT BETWEEN AND`

### LIKE

ƥ���ѯ��`%`�������ַ�������0������`_`����һ���ַ�

`SELECT * FROM books WHRER name LIKE '%��ȫ'`

### DISTINCT

`DISTINCT`������ڲ�ѯ���Ŀ�ͷ��������`COUNT()`���ʹ�ã����ز��ظ����ֶΣ��������`ALL`

`SELECT DISTINCT author FROM books WHRER name LIKE '%��ȫ'`

### SOME

`SOME`����`ANY`����ʾΪ��������һ������

`SELECT * FROM books WHRER price > SOME(5, 15, 20)` ��ʾ��ѯ�۸����5������

### EXISTS

`SELECT FROM ���� WHRER EXISTS(�Ӳ�ѯ)`

����Ӳ�ѯ�Ľ�����з����У���ִ�и���ѯ

#### ��������

Ƕ����FROM�Ӿ��е�SELECT�����ҪΪderived tableָ������ `(�Ӳ�ѯ) AS ����`

- ����`SELCET * FROM (SELECT * FORM ����)`
- ��ȷ��`SELCET * FROM (SELECT * FORM ����) AS ����`

### GROUP BY

�����ѯ�������ֶη����ѯ��ϲ����

- `SELECT subjectNo, AVG(studentResult) from result group by subjectNo`
- `SELECT COUNT(*), gradeId FROM student GROUP BY gradeId;`

ʹ��`HAVING`�Է��������ݽ���ɸѡ���޷�ʹ��`WHERE`�Է��������ݽ���ɸѡ

```sql
SELECT COUNT(*) '����', gradeId '�꼶', sex '�Ա�'
FROM student
GROUP BY gradeId
HAVING COUNT(*) > 2;
```

ִ��˳��`WHERE` -\> `GROUP BY` -\> �ۺϺ��� -\> `HAVING`

### UNION

`UNION`�ϲ������ѯ�Ľ����ȥ���ظ���Ŀ���ϲ�ʱ���Ӧ���ֶκ��������ͱ�����ͬ

`UNION ALL`��ȥ���ظ�

### ���Ӳ�ѯ

- **������** ���ݱ��й�ͬ����������ƥ��
- **������** ���ٷ���һ�����е����м�¼������ƥ��������ѡ��ش���һ�����з��ؼ�¼

#### ������

ʹ��`WHERE`��������

```sql
SELECT student.studentName, result.subjectNo, result.studentResult
FROM student, result
WHERE student.studentNo = result.studentNo;  
```

ʹ��`[INNER] JOIN ... ON ..`��������

```sql
SELECT studentName, subjectNo, studentResult
FROM student
INNER JOIN result ON student.studentNo = result.studentNo;
```

��ʹ�ñ����������������е�������������в�ѯ���ֶβ��ظ�����Ҫ����һ�в���Ҫָ��������

```sql
SELECT S.studentName, R.subjectNo, R.studentResult
FROM student as S
INNER JOIN result as R ON S.studentNo = R.studentNo;
```

���ԶԶ����ʹ��������

```sql
SELECT S.studentName, SU.subjectName, R.studentResult
FROM student AS S
INNER JOIN result AS R ON S.studentNo = R.studentNo
INNER JOIN subject AS SU ON SU.subjectNo = R.subjectNo;
```

### ������

`<LEFT | RIGHT> OUTER JOIN ... ON`

```sql
SELECT S.studentName, R.studentResult
FROM student AS S
LEFT OUTER
JOIN result AS R ON S.studentNo = R.studentNo;
```

LEFT OUTTER JOIN���ؽ����������������

## ����

�ۺϺ���

- `AVG()`
- `COUNT()`
- `MAX()`, `MIN()`
- `SUM()`

�ַ���������ע���±��1��ʼ

- `CONCAT(str1, str2, ...)`
- `INSERT(str, pos, len, newstr)`
- `LOWER(str)`, `UPPER(str)`
- `SUBSTRING(str, pos, len)`

ʱ�����ں���

- `CURDATE()` ��ǰ����
- `CURTIME()` ��ǰʱ��
- `NOW()` ��ǰ���ں�ʱ��
- `WEEK(date)`, `YEAR(date)`, `HOUR(date)`, `MINUTE(date)`
- `DATEDIFF(date1, date2)` �����������������
- `ADDDATE(date, n)` ����date

��ѧ����

- `CEIL(x)`, `FLOOR(x)`
- `RAND()`

## Chapter 8 ���� Transaction

��һϵ�����ݲ�����Ϊ���崦�����ĳһ�����ύ�ɹ����������������ݸ��ľ����ύ����Ϊ���ݿ��е�������ɲ��֣��������ִ��ʱ���������ұ���ȥȡ����ع��������ݽ�ȫ���ָ�������ǰ��״̬���������ݵĸ��ľ��������

- ԭ���� ����ĸ���Ԫ���ǲ��ɷָ��
- һ���� �������ʱ���ݴ���һ��״̬��������ʹ���ݴ��洦�ڲ��ȶ���״̬
- ������ ���в�������˴˸��룻����Ӧ�����κη�ʽ������Ӱ����������
- �־��� һ�������ύ�������ݿ��Ӱ�������õ�

InnoDB�������֧�֣�REDO��־��UNDO��־

```sql
BEGIN;  // �� START TRANSACTION;

COMMIT;
ROLLBACK;
```

δ��ʽ��������ʱ��ÿ��SQL��䶼��Ϊ����������ִ������Զ��ύ��

MySQL�����Զ��ύ `SET autocommit = 0 | 1;`

## ��ͼ

- ������ͼ `CREATE VIEW ��ͼ��������view_xxx��v_xxx��AS <SELECT ���>`
- ɾ����ͼ `DROP VIEW [IF EXISTS] ��ͼ��`
- �鿴��ͼ `SELECT * FROM ��ͼ��`

��ͼ����Ƕ����һ����ͼ������ͼ���ݵ�����޸ĺͲ�ѯֱ��Ӱ����е����ݡ�

## ����

- ��ͨ���� �����ظ��Ϳ�ֵ
- Ψһ���� �������ظ��������ֵ
- �������� ��һ����ʱ�Զ������������������������ǿղ��ظ�
- �������� ������������Ϊ�������ڲ�ѯ��ֻ��ʹ�����������������ߵ��ֶ�ʱ�������Żᱻʹ��
- ȫ������ ��Ҫ������CAHR, TEXT��������ȫ�ļ�������
- �ռ����� MyISAM only���Կռ��������ͽ���������

������ʹ������

- �������� `CREATE [UNQIUE | FULLTEXT | SPATIAL |] INDEX ������ ON ����(�ֶ���[�ַ������Ϳ���ָ����������])`
- ɾ������ `DROP INDEX ����.������`
- �鿴���� `SHOW INDEX FROM ����`

## ���ݱ���

����

- `mysqldump -u usrname -h host -p dbname[tbname1, ...] > filename.sql`

�򱣴浽�ⲿ�ļ�

```sql
SELECT * FROM ...
INTO OUTFILE 'filename' [OPTIONS]
```

�ָ�

- `mysql -u -p [dbname] < filename.sql`
- ��½MySQL��`source filename`

---

## ��װ

��Linux���а��Ͽ���ʹ��[���ù���](https://downloads.mariadb.org/mariadb/repositories/)��

### ��CentOS7�ϰ�װ

```sh
yum install mariadb-server
sudo systemctl start mariadb
sudo systemctl status mariadb
sudo systemctl enable mariadb
sudo mysql_secure_installation
mysqladmin -u root -p version
```

## ����

### �����ļ�

`/etc/my.cnf`

���ļ��п����޸�Ĭ�ϱ��롣

### �����й���

`mysql -u root -p`

- ��ʾ����Ŀ¼ `HELP contents`
- �鿴�汾�͵�½�û���`SELECT VERSION(), USER()`
- ����������ʱ���

    ```shell
    SET NAMES gbk;
    # �൱�������������
    SET character_set_client=gbk;
    SET character_set_results=gbk;
    SET character_set_connection=gbk;
    ```

- �鿴֧�ֵ����� `SHOW ENGINES`, `SHOW VARIABLES LIKE 'storage_engine'`

---

## �ĵ�����

- [MariaDB Server�ĵ�](https://mariadb.com/kb/en/library/documentation/)
