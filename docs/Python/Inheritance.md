# Inheritance

```python
class Person:
    def __init__(self, name):
        self.name = name
    def say_hello(self):
        print('hello, ', + self.name)
    def __str__(self):
        return self.name

class Student(Person):
    def __init__(self, name, school):
        super().__init__(name):
        self.school = school
    def say_hello(self):
        super().say_hello()
        print('I am rather tired')
    def __str__(self):
        return f'{self.name} attends {self.school}'
    def sing_school_song(self):
        print('Ode to ' + self.school)
```

```python
student = Student('Christopher', 'UMD')
student.say_hello()
student.sing_school_song()
# What are you?
print(isinstance(student, Student)) # True
print(isinstance(student, Person)) # True
print(isinstance(Student, Person)) # True
```

 