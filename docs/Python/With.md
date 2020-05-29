# With

## Writing to a file

```python
stream = open('output.txt', 'wt')
stream.write('Lorem ipsum dolar')
stream.close()
```

重构上面的代码

```python
try:
    stream = open('output.txt', 'wt')
    stream.write('Lorem ipsum dolar')
finally:
    stream.close()
```

Simplifying with `with`

```python
with open('output.txt', 'wt') as stream:
    stream.write('Lorem ipsum dolar')
```

