# Using LIsts as Queues

To implement a queue, use [`collections.deque`](https://docs.python.org/3/library/collections.html#collections.deque) which was designed to have fast appends and pops from both ends. For example:

```python
from collections import deque
queue = deque(["Eric", "John", "Michael"])
queue.append("Terry")
queue.append("Graham")
print(queue)
# ["Eric", "John", "Michael", "Terry", "Graham"]
x = queue.popleft();
print(x)
# "Eric"
print(queue)
# ["John", "Michael", "Terry", "Graham"]
```

