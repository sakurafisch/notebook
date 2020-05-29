# Multiple Inheritance

多重继承

## 用处

- 启用框架的特性，比如Django。

- 简化重复操作

```python
class Loggable:
    def __init__(self):
        self.title = ''
    def log(self):
        print('Log message from ' + self.title)
        
class Connection:
    def __init__(self):
        self.server = ''
    def connect(self):
        print('Connecting to database on' + self.server)

def framework(item):
    # Perform the connection
    if isinstance(item, Connection):
        item.connect()
    # Log the operation
    if isinstance(item, Loggable):
        item.log()

# Use the framework
# Inherit from Connection and Loggable
class SqlDatabase(Connection, Loggable):
    def __init__(self):
        super().__init__()
        self.title = 'Sql Connection Demo'
        self.server = 'Some_Server'

# Create an instance of our class
sql_connection = SqlDatabase()
# Use our framework
framework(sql_connection) # connects and logs

# 输出如下：
# Connecting to database on Some_Server
# Log message from sql Connection Demo
```

