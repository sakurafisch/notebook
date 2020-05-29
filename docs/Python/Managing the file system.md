# Managing the file system

## Old style

All the common operations are built into Python Static or classless
`os.path`

## Python 3.6 or higher

Class based
`Path` from `pathlib` library

Better performance as it can avoid calls to the OS

## Working with paths

```python
# Python 3.6 or higher
# Grab the library
from pathlib import Path

# Where am I?
cwd = Path.cwd()
print(str(cwd))

# Combine parts to create full path and file name
new_file = Path.joinpath(cwd, 'new_file.txt')
print(str(new_file))

# Does this exist?
print(str(new_file.exists))

```

## Working with directories

```python
from pathlib import Path
cwd = Path.cwd()

# Get the parent directory
parent = cwd.parent

# Is this a directory
print(str(parent.is_dir())) # True

# Is this a file?
print(str(parent.is_file())) # False

# List child directories
for child in parent.iterdir():
    if child.is_dir():
        print(str(child)) # 打印当前目录的子文件夹
```

## Working with files

```python
from pathlib import Path
cwd = Path.cwd()
demo_file = Path(Path.joinpath(cwd, 'demo.txt'))

# Get the file name
print(str(demo_file.name))

# Get the extension
print(str(demo_file.suffix))

# Get the folder
print(str(demo_file.parent.name))

# Get the size
print(str(demo_file.stat().st_size))
```

## Opening a file

```python
stream = open(file_name, mode, buffer_size)
```

Mode:

- `r` - Read(default)

- `w` - Truncate and write
- `a` - Append if file exists
- `x` - Write, fail if file exists
- `+` - Updating (read/write)
- `t` - Text(default)
- `b` - Binary

## Reading from a file

```python
stream = open('demo.txt')

print(stream.readable()) # Can we read?
print(stream.read(1)) # Read the first character
print(stream.readline()) # Read a line

stream.close() # close the stream
```

## Write to a file

```python
stream = open('output.txt', 'wt') # write text

stream.write('H') # write a single string
stream.writelines(['ello', ',', 'world']) # write multiple strings
stream.write('\n') # Write a new line

stream.close() # close the stream (and flush data)
```

## Managing the stream

```python
stream = open('output.txt', 'wt')
stream.write('demo!')
stream.seek(0) # Put the cursor back at the start
stream.write('cool')
stream.flush() # Write the data to file
stream.close() # Flush and close the stream
```

