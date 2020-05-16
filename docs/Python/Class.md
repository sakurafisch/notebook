# Class

## 类定义

```python
class Presenter():
    def __init__(self, name):
        # Constructor
        self.name = name
    def say_hello(self):
        # method
        print('Hello, ' + self.name)
```

## 实例化

```python
presenter = Presenter('Chris')
```

## 更改字段的值

```python
presenter.name = 'Christopher'
```

## 调用对象的方法

```python
presenter.say_hello()
```

## 访问权限

EVERYTHING is public.

\_ (single underscore) means avoid unless you really know what you're doing.

\_\_ (double underscore) means do not use.

## 实践示例

```python
class Presenter():
    def __init__(self, name):
        # Constructor
        self.name = name
    
    @property
    def name(self):
        print('In the getter')
        return self.__name
    
    @name.setter
    def name(self, value):
        print('In setter')
        # cool validation here
        self.__name = value

presenter = Presenter('Chris')
presenter.name = 'Christopher'
print(presenter.name)

# 输出如下
# In setter
# In setter
# In the getter
# Christopher
```



